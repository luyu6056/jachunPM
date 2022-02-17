package codec

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/luyu6056/gnet"
	"github.com/luyu6056/tls"
)

const (
	writeDeadline = time.Second * 5
)

type WSClient struct {
	wsconn          *WSconn
	c               gnet.Conn
	readbuf, msgbuf *tls.MsgBuffer
	err             chan error
	msgRead         chan bool
	msgLock         sync.Mutex
}

func DailWebSocket(Url string) (c net.Conn, err error) {
	u, err := url.Parse(Url)
	if err != nil {
		return nil, err
	}

	var challengeKey, _ = GenerateChallengeKey()

	httphandshake := fmt.Sprintf("GET %s HTTP/1.1\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nHost: %s\r\nOrigin: %s\r\nSec-WebSocket-Key: %s\r\nSec-WebSocket-Version: 13\r\n\r\n", u.Path, u.Host, u.Scheme+"://"+u.Host, challengeKey)

	dialTcp := func(ip string) (net.Conn, error) {
		if _, err := net.ResolveTCPAddr("tcp4", ip); err == nil {
			return net.Dial("tcp4", ip)
		}
		if _, err := net.ResolveTCPAddr("tcp6", ip); err == nil {
			return net.Dial("tcp6", ip)
		}
		return nil, errors.New("无法连接 " + ip)
	}
	conn, err := dialTcp(u.Host)
	//尝试dns解析后再连接
	if err != nil {
		ns, _ := net.LookupHost(u.Host)
		for _, ip := range ns {
			conn, _ := dialTcp(ip)
			if conn != nil {
				break
			}
		}
	}

	if conn == nil {
		return nil, err
	}
	//发送握手包
	if _, err = conn.Write([]byte(httphandshake)); err != nil {
		return
	}

	http := &Httpserver{}


	http.c = &WSClientGnetInterface{conn: conn, inboundBuffer: new(tls.MsgBuffer), outboundBuffer: new(tls.MsgBuffer)}
	client := &WSClient{c: http.c, readbuf: new(tls.MsgBuffer), msgbuf: new(tls.MsgBuffer), err: make(chan error, 1), msgRead: make(chan bool, 1)}
	b := make([]byte, 1024)
	for {
		n, err := conn.Read(b)
		if err != nil {
			return nil, err
		}

		client.readbuf.Write(b[:n])
		_, req, err := http.Parsereq(client.readbuf.Bytes())
		if err != nil {
			return nil, err
		}
		if req != nil {
			http.WorkRequest=req
			break
		}
	}
	if ComputeAcceptKey(challengeKey) != http.WorkRequest.Header("Sec-WebSocket-Accept") {
		return nil, errors.New("websocket握手失败")
	}
	client.readbuf.Reset()
	client.wsconn = &WSconn{Http: http, Write: client.c.AsyncWrite, ReadFinal: true, readbuf: &tls.MsgBuffer{}}
	go func() {
		for {

			n, err := conn.Read(b)
			if err != nil {

				client.err <- err
				return
			}

			client.readbuf.Write(b[:n])
			for client.readbuf.Len() > 0 {
				msgtype, data, err := client.wsconn.ReadMessage(client.readbuf.Bytes())
				client.readbuf.Next(client.wsconn.ReadLength)

				if msgtype == noFrame {
					break
				}
				if err != nil {
					client.err <- err
					return
				}
				if msgtype == BinaryMessage {

					client.msgLock.Lock()
					client.msgbuf.Write(data)
					client.msgLock.Unlock()
					select {
					case client.msgRead <- true:
					default:
					}
				}
			}

		}
	}()

	return client, nil
}

func (client *WSClient) Close() error {
	return client.c.Close()
}
func (client *WSClient) RemoteAddr() net.Addr {
	return client.c.RemoteAddr()
}
func (client *WSClient) LocalAddr() net.Addr {
	return client.c.LocalAddr()
}
func (client *WSClient) Read(b []byte) (int, error) {
	select {
	case err := <-client.err:
		return 0, err
	case <-client.msgRead:
		client.msgLock.Lock()
		defer client.msgLock.Unlock()
		if client.msgbuf.Len() > len(b) {
			copy(b, client.msgbuf.Next(len(b)))
			select {
			case client.msgRead <- true:
			default:
			}
			return len(b), nil
		}
		copy(b, client.msgbuf.Bytes())
		l := client.msgbuf.Len()
		client.msgbuf.Reset()
		return l, nil
	}
	return 0, nil
}
func (client *WSClient) SetDeadline(time.Time) error      { return nil }
func (client *WSClient) SetReadDeadline(time.Time) error  { return nil }
func (client *WSClient) SetWriteDeadline(time.Time) error { return nil }
func (client *WSClient) Write(b []byte) (int, error) {
	client.wsconn.WriteMessage(BinaryMessage, b)
	return len(b), nil
}

type WSClientGnetInterface struct {
	conn                          net.Conn
	tlsconn                       *tls.Conn
	inboundBuffer, outboundBuffer *tls.MsgBuffer
}

func (c *WSClientGnetInterface) Close() error {
	return c.conn.Close()
}

func (c *WSClientGnetInterface) Write(data []byte) error {
	if c.conn == nil {
		return io.EOF
	}
	var err error
	if c.tlsconn != nil {
		if err = c.tlsconn.Write(data); err != nil {
			return err
		}
		c.conn.SetWriteDeadline(time.Now().Add(writeDeadline))
		_, err = c.conn.Write(c.outboundBuffer.Bytes())
		c.outboundBuffer.Reset()
	} else {
		c.conn.SetWriteDeadline(time.Now().Add(writeDeadline))
		_, err = c.conn.Write(data)

	}
	return err
}

func (c *WSClientGnetInterface) BufferLength() int {
	return c.inboundBuffer.Len()
}

func (c *WSClientGnetInterface) Context() interface{} {
	return nil
}
func (c *WSClientGnetInterface) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}
func (c *WSClientGnetInterface) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
func (c *WSClientGnetInterface) Read() []byte {
	return c.inboundBuffer.Bytes()
}
func (c *WSClientGnetInterface) ReadN(n int) (int, []byte) {
	buf := c.inboundBuffer.PreBytes(n)
	return len(buf), buf
}
func (c *WSClientGnetInterface) ResetBuffer() {
	c.inboundBuffer.Reset()
}
func (c *WSClientGnetInterface) SendTo(b []byte) error {
	_, err := c.conn.Write(b)
	return err
}
func (c *WSClientGnetInterface) SetContext(i interface{}) {}
func (c *WSClientGnetInterface) ShiftN(n int) int {
	c.inboundBuffer.Next(n)
	return n
}
func (c *WSClientGnetInterface) Wake() error { return nil }

//这里的AsyncWrite是WSconn最终调用的Write出口
func (c *WSClientGnetInterface) AsyncWrite(data []byte) error {
	return c.Write(data)
}
func (c *WSClientGnetInterface) FlushWrite(data []byte, b ...bool) {
	c.Write(data)

}
func (c *WSClientGnetInterface) UpgradeTls(config *tls.Config) error {
	fmt.Println("UpgradeTls未处理")
	return nil
}
func (c *WSClientGnetInterface) WriteNoCodec(data []byte) error { return c.Write(data) }

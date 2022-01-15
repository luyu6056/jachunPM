package protocol

import (
	"libraries"
	"net"
	"sync"
	"time"

	"github.com/luyu6056/gnet"
)

type gnetClient struct {
	*gnet.EventServer
}

var GnetClient = gnet.Client(&gnetClient{}, gnet.WithCodec(&RpcCodec{}), gnet.WithTCPNoDelay(true))

func dail(network, addr string, rpc *RpcClient) (net.Conn, error) {
	c, err := GnetClient.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	c.SetContext(rpc)
	conn := newClient(c)

	return conn, err
}

func (rs *gnetClient) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {

	return
}
func (rs *gnetClient) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	if rpc, ok := c.Context().(*RpcClient); ok {
		rpc.reconnect <- nil
	}
	return gnet.None
}
func (rs *gnetClient) React(data []byte, c gnet.Conn) (action gnet.Action) {
	if rpc, ok := c.Context().(*RpcClient); ok {
		msgnum := data[0]
		index := 1
		var n int
		for index < len(data) {
			n++
			msg, l, err := ReadOneMsgFromBytes(data[index:])
			index += l
			if err != nil {
				libraries.ReleaseLog("读消息错误%v", err)
			} else {
				msg.ReadData()
				rpc.inchan <- msg
			}

		}
		if n != int(msgnum) {
			libraries.DebugLog("读消息数量错误，请检查协议，消息总量%d,已读%d", msgnum, n)
		}
		rpc.window -= int32(msgnum)
		if rpc.Status&RpcClientStatuNormal == RpcClientStatuNormal && rpc.window < DefaultWindowSize/2 {
			data := GET_MSG_HOST_WINDOW_UPDATE()
			data.Add = DefaultWindowSize - rpc.window
			rpc.window = DefaultWindowSize
			//libraries.DebugLog("增加窗口%d，实际窗口%d", data.Add, rpc.window)
			rpc.sendStruct.SendMsgToDefault(nil, data)
			data.Put()
		}
	}

	return gnet.None
}

type client struct {
	msgbuf  *libraries.MsgBuffer
	conn    gnet.Conn
	err     chan error
	msgRead chan bool
	msgLock sync.Mutex
}

func newClient(conn gnet.Conn) *client {
	return &client{
		conn:    conn,
		msgbuf:  new(libraries.MsgBuffer),
		msgRead: make(chan bool, 1),
		err:     make(chan error, 1),
	}
}
func (c *client) Close() error {
	c.conn.Close()
	return nil
}
func (c *client) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
func (c *client) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}
func (c *client) Read(b []byte) (int, error) {

	return 0, nil
}
func (c *client) SetDeadline(time.Time) error      { return nil }
func (c *client) SetReadDeadline(time.Time) error  { return nil }
func (c *client) SetWriteDeadline(time.Time) error { return nil }
func (c *client) Write(b []byte) (int, error) {

	c.conn.AsyncWrite(b)
	return len(b), nil
}

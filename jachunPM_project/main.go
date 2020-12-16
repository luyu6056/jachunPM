package main

//2019-10-16 修改为gnet作为http服务器
import (
	"bbs/config"
	_ "bbs/controllers/admin"
	"bbs/controllers/public"
	_ "bbs/controllers/web"
	"bbs/db"
	"bbs/db/mysql"
	"bbs/libraries"
	"bbs/models"
	"bbs/server"
	"bbs/vaptcha"
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"

	_ "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/luyu6056/gnet"
	"github.com/luyu6056/tls"
	"github.com/panjf2000/ants/v2"
)

type mainServer struct {
	*gnet.EventServer
	addr      string
	pool0     *ants.Pool
	pool1     *ants.Pool
	RpcServer bool
}

var Context_pool = sync.Pool{New: func() interface{} {
	return &server.Context{Buf: new(libraries.MsgBuffer), In: new(libraries.MsgBuffer), In2: new(libraries.MsgBuffer)}
}}

func (hs *mainServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {

	mysql.Init_register_db = append(mysql.Init_register_db, db.DB_init)
	mysql.Init_register_db = append(mysql.Init_register_db, db.DBlog_init)
	mysql.Mysql_init()
	models.Model_init()
	//web.Spider_GetThread()
	libraries.Log("server started on %s (loops: %d)", hs.addr, srv.NumEventLoop)

	return
}

func (hs *mainServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {

	return
}

func (hs *mainServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	switch svr := c.Context().(type) {
	case *server.RpcServer:

		if svr.ServerId == 0 {
			if v, ok := rpcServerConnMap.Load(svr.GroupId); ok {
				v.(*sync.Map).Range(func(_, v1 interface{}) bool {
					ctx := Context_pool.Get().(*server.Context)
					ctx.Conn = v1.(*server.ClientConn)
					ctx.Conn_m = v.(*sync.Map)
					svr.Conn = ctx.Conn
					public.ConnDown(ctx)
					Context_pool.Put(ctx)
					return true
				})
				rpcServerConnMap.Delete(svr.GroupId)
			}
		}

	case *server.WSconn:
		if svr.Conn != nil {
			ctx := Context_pool.Get().(*server.Context)
			ctx.Conn = svr.Conn
			public.ConnDown(ctx)
			Context_pool.Put(ctx)
		}
	case *server.Httpserver:
		if svr.Conn != nil {
			ctx := Context_pool.Get().(*server.Context)
			ctx.Conn = svr.Conn
			public.ConnDown(ctx)
			Context_pool.Put(ctx)
		}
		svr.Conn = nil
		server.Httppool.Put(svr)

		//hs := svr.Http
		//svr.Http = nil
		//server.Httppool.Put(hs)
	case *server.Http2server:
		if err == gnet.ErrServerShutdown {
			svr.Close()
		} else {
			svr.ReadPool.Submit(func() {
				svr.Close()
			})
		}

	}
	c.SetContext(nil)
	return
}

func (hs *mainServer) React(data []byte, c gnet.Conn) (action gnet.Action) {

	ctx := Context_pool.Get().(*server.Context)
	defer Context_pool.Put(ctx)
	switch svr := c.Context().(type) {
	case *server.Httpserver:
		if svr.Request.Connection == "close" {
			action = gnet.Close
		}
		switch svr.Request.Path {
		case "/hello":
			c.AsyncWrite([]byte("HTTP/1.1 200 OK\r\nContent-Length: 11\r\n\r\nhello word1"))
			return
		case "/post":
			ctx.In.Reset()
			n, _ := base64.StdEncoding.Decode(ctx.In.Make(len(data)-2), data[2:])
			ctx.In.Truncate(n)
			svr.BeginRequest("test", ctx)

		case "/upload":

			ctx.In.ResetBuf(data)
			svr.BeginRequest("upload", ctx)
		case "/ws", "/admin":
			//tmp_ctx.Request.Body.Reset()
			//tmp_ctx.Request.Body.Write(data[len(data)-8 : len(data)])
			err := svr.Upgradews(c)
			if err != nil {
				action = gnet.Close
				return
			}
			//注册到不同的接口
			switch svr.Request.Path {
			case "/ws":
				svr.Ws.Handler = server.WsWebhand
			case "/admin":
				svr.Ws.Handler = server.WsAdminhand
			}
			server.Httppool.Put(svr)
			return gnet.None
		case "/vaptcha":
			vaptcha.Hander(svr, c)
		default:
			svr.Static()
		}
		return
	case *server.WSconn:

		ctx.In.ResetBuf(data)
		ctx.Conn = svr.Conn
		for ctx.In.Len() > 0 {
			svr.Handler(ctx)
		}
	case *server.Http2server:
		svr.SendPool.Invoke(svr.WorkStream)
		/*case *server.RpcServer:
			ctx := Context_pool.Get().(*server.Context)
			v, ok := rpcServerConnMap.Load(svr.GroupId)
			if !ok {
				Context_pool.Put(ctx)
				return
			}
			ctx.Conn_m = v.(*sync.Map)
			ctx.In2.Reset()
			ctx.In2.Write(data)
			if svr.ServerId == 0 { //主逻辑
				hs.pool0.Submit(func() {
					for ctx.In2.Len() > 10 { //2字节消息头,4字节fd,4字节cmd
						msglen := int(binary.LittleEndian.Uint16(ctx.In2.Next(2)))
						if ctx.In2.Len() < msglen {
							break
						}

						ctx.In.Reset()
						ctx.In.Write(ctx.In2.Next(msglen))
						var fd [4]byte
						copy(fd[:], ctx.In.Next(4))
						if v, ok := ctx.Conn_m.Load(fd); ok {
							ctx.Conn = v.(*server.ClientConn)
							svr.Conn = ctx.Conn
							server.WebHandle(ctx)
						}

					}
					Context_pool.Put(ctx)
				})
			} else {
				hs.pool1.Submit(func() {
					for ctx.In2.Len() > 10 { //2字节消息头,4字节fd,4字节cmd
						msglen := int(binary.LittleEndian.Uint16(ctx.In2.Next(2)))
						if ctx.In2.Len() < msglen {
							break
						}
						ctx.In.Reset()
						ctx.In.Write(ctx.In2.Next(msglen))
						var fd [4]byte
						copy(fd[:], ctx.In.Next(4))
						cmd := protocol.READ_int32(ctx.In)
						switch cmd {
						case protocol.CMD_MSG_C2S_Conn_Client:
							data := protocol.READ_MSG_C2S_Conn_Client(ctx.In)
							public.ConnClient(data, ctx, svr)
							data.Put()
							binary.LittleEndian.PutUint32(fd[:], uint32(data.Fd))
							svr.Conn = ctx.Conn
							public.SendServerOK(ctx)
						case protocol.CMD_MSG_Conn_Down:
							data := protocol.READ_MSG_Conn_Down(ctx.In)
							binary.LittleEndian.PutUint32(fd[:], uint32(data.Fd))
							if v, ok = rpcServerConnMap.Load(data.GroupId); ok {
								ctx.Conn_m = v.(*sync.Map)
								if v, ok := ctx.Conn_m.Load(fd); ok {
									ctx.Conn = v.(*server.ClientConn)
									public.ConnDown(ctx)
								}
							}
							data.Put()
						}

					}
					Context_pool.Put(ctx)
				})

			}

		case nil:

			ctx := Context_pool.Get().(*server.Context)
			ctx.In.Reset()
			if len(data) == 52 {
				ctx.In.Write(data[4:])
				cmd := protocol.READ_int32(ctx.In)
				if cmd == protocol.CMD_MSG_C2S_Regedit {
					data := protocol.READ_MSG_C2S_Regedit(ctx.In)
					now := time.Now().UnixNano()
					if (now/1e9-data.Time/1e9) < 2 && data.Key == libraries.MD5_S(config.Server.RpcKey+strconv.Itoa(int(data.Time))) {
						s := server.NewRpcServer(c)

						s.ServerId = data.Serverid
						s.GroupId = data.GroupId
						s.ZstdDecoder, _ = zstd.NewReader(nil)
						if s.ServerId == 0 {
							conn_m := new(sync.Map)
							conn_m.Store([4]byte{0, 0, 0, 0}, &server.ClientConn{})
							rpcServerConnMap.Store(data.GroupId, conn_m)
						}
						c.SetContext(s)
					}
				}
			}
			Context_pool.Put(ctx)*/
	}
	return
}

var rpcServerConnMap sync.Map

func main() {

	//p := pool.NewWorkerPool()
	pool0, _ := ants.NewPool(int(config.Server.DBMaxConn), ants.WithPreAlloc(true))  //限制协程数量避免mysql的chan竞争
	pool1, _ := ants.NewPool(int(config.Server.DBMaxConn), ants.WithPreAlloc(false)) //限制协程数量避免mysql的chan竞争
	svr := &mainServer{addr: config.Server.Listen, pool0: pool0, pool1: pool1, RpcServer: false}
	defer svr.pool0.Release()
	defer svr.pool1.Release()
	// Start serving!
	var tlsconfig *tls.Config
	if config.Server.IsHttps {
		tlsconfig = &tls.Config{
			NextProtos:               []string{"h2", "http/1.1"},
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
				tls.TLS_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
			},
		}
		server_cert, err := ioutil.ReadFile(config.Server.TLScert)
		if err != nil {
			log.Fatalf("读取服务器证书cert错误 err: %v", err)
		}
		server_key, err := ioutil.ReadFile(config.Server.TLSkey)
		if err != nil {
			log.Fatalf("读取服务器证书key错误 err: %v", err)
		}
		if ca, err := ioutil.ReadFile(config.Server.TLSca); err == nil {
			ca = bytes.TrimLeft(ca, "\n")
			server_cert = bytes.Replace(server_cert, ca, nil, 1)
			server_cert = append(server_cert, ca...)
			certPool := x509.NewCertPool()
			if ok := certPool.AppendCertsFromPEM(ca); ok {
				tlsconfig.ClientCAs = certPool
			} else {
				libraries.DEBUG("ca_cert加载失败", err)
			}
		} else {
			libraries.DEBUG("ca_cert读取错误")
		}

		cert, err := tls.X509KeyPair(server_cert, server_key)
		if err != nil {
			log.Fatalf("tls.LoadX509KeyPair err: %v", err)
		}
		tlsconfig.Certificates = []tls.Certificate{cert}

		go func() {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				r.Header.Add("strict-transport-security", "max-age=31536000; includeSubDomains; preload")
				http.Redirect(w, r, config.Server.Origin, http.StatusFound)
			})
			http.ListenAndServe("0.0.0.0:80", nil)
		}()
	}

	option := []gnet.Option{gnet.WithLoopNum(runtime.NumCPU() * 2), gnet.WithTCPKeepAlive(time.Second * 600), gnet.WithCodec(&server.Tlscodec{}), gnet.WithReusePort(true), gnet.WithOutbuf(1024), gnet.WithTls(tlsconfig), gnet.WithMultiOut(false)}
	log.Fatal(gnet.Serve(svr, config.Server.Listen, option...))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

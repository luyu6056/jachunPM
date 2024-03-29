package main

import (
	"bytes"
	"crypto/x509"
	"io/ioutil"
	"jachunPM_http/db"
	"jachunPM_http/handler"
	"jachunPM_http/setting"
	"libraries"
	"log"
	"net/http"
	_ "net/http/pprof"
	"protocol"
	"runtime"
	"strconv"
	"time"

	"github.com/panjf2000/ants/v2"

	"codec"

	"github.com/luyu6056/cache"
	"github.com/luyu6056/gnet"
	"github.com/luyu6056/tls"
)

type httpServer struct {
	*gnet.EventServer
	addr string
}

func main() {
	cache.StartWebServer("0.0.0.0:809")
	var err error
	handler.HostConn, err = protocol.NewClient(protocol.HttpServerNo, setting.Setting.HostIP, setting.Setting.TokenKey)
	if err != nil {
		libraries.ReleaseLog("连接host %s 服务启动失败%v", setting.Setting.HostIP, err)
		return
	} else {
		db.Init()
		handler.HostConn.SetTickHand(handler.HandleTick)
		handler.HostConn.HandleMsg = handler.Handler
		handler.HostConn.DB = db.DB
		go handler.HostConn.Start()
		time.Sleep(time.Second * 2)
		handler.Init()
		handler.Company_updateCache()
	}

	go http.ListenAndServe("0.0.0.0:"+strconv.Itoa(8100+protocol.HttpServerNo), nil)
	svr := &httpServer{addr: setting.Setting.ListenHttp}
	// Start serving!
	var tlsconfig *tls.Config
	if setting.Setting.HttpsTLScert != "" && setting.Setting.HttpsTLSca != "" && setting.Setting.HttpsTLSkey != "" {
		tlsconfig = &tls.Config{
			NextProtos:               []string{"http/1.1"},
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
		server_cert, err := ioutil.ReadFile(setting.Setting.HttpsTLScert)
		if err != nil {
			log.Fatalf("读取服务器证书cert错误 err: %v", err)
		}
		server_key, err := ioutil.ReadFile(setting.Setting.HttpsTLSkey)
		if err != nil {
			log.Fatalf("读取服务器证书key错误 err: %v", err)
		}
		if ca, err := ioutil.ReadFile(setting.Setting.HttpsTLSca); err == nil {
			ca = bytes.TrimLeft(ca, "\n")
			server_cert = bytes.Replace(server_cert, ca, nil, 1)
			server_cert = append(server_cert, ca...)
			certPool := x509.NewCertPool()
			if ok := certPool.AppendCertsFromPEM(ca); ok {
				tlsconfig.ClientCAs = certPool
			} else {
				log.Fatalf("ca_cert加载失败 %v", err)
			}
		} else {
			log.Fatalf("ca_cert读取错误 %v")
		}

		cert, err := tls.X509KeyPair(server_cert, server_key)
		if err != nil {
			log.Fatalf("tls.LoadX509KeyPair err: %v", err)
		}
		tlsconfig.Certificates = []tls.Certificate{cert}

		/*go func() {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				r.Header.Add("strict-transport-security", "max-age=31536000; includeSubDomains; preload")
				http.Redirect(w, r, setting.Setting.Origin, http.StatusFound)
			})
			http.ListenAndServe("0.0.0.0:80", nil)
		}()*/
	}

	log.Fatal(gnet.Serve(svr, svr.addr, gnet.WithLoopNum(runtime.NumCPU()), gnet.WithTCPKeepAlive(time.Second*600), gnet.WithCodec(&codec.Tlscodec{}), gnet.WithReusePort(true), gnet.WithOutbuf(128), gnet.WithTls(tlsconfig), gnet.WithMultiOut(false), gnet.WithTCPNoDelay(true)))
}
func (hs *httpServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	libraries.DebugLog("httpserver started on %s (loops: %d)", hs.addr, srv.NumEventLoop)
	return
}

func (hs *httpServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	time.AfterFunc(time.Second*10, func() {
		time.AfterFunc(time.Second*10, func() {
			if ctx, ok := c.Context().(*codec.WSconn); !ok || ctx.Conn == nil || ctx.Conn.Session == nil {
				c.Close()
			}
		})
	})
	return
}

func (hs *httpServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	switch svr := c.Context().(type) {
	case *codec.WSconn:

	case *codec.Httpserver:
		svr.Close()

	case *codec.Http2server:
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

var reactPool, _ = ants.NewPool(runtime.NumCPU() * 2)

func (hs *httpServer) React(data []byte, c gnet.Conn) (action gnet.Action) {
	switch svr := c.Context().(type) {
	case *codec.Httpserver:
		if data == nil {
			svr.Wake()
		} else {
			req := svr.WorkRequest
			reactPool.Submit(func() {
				if handler.HttpHandler(req) == gnet.Shutdown {
					svr.Close()
				}
				req.Wake()
			})
			svr.WorkRequest = codec.NewRequest(svr)
		}

		//req.Httpfinish()
		//req.Recovery()
		return
	case *codec.WSconn:

	case *codec.Http2server:
		//获取当前帧
		stream := svr.WorkStream
		//异步执行
		reactPool.Submit(func() {
			if handler.HttpHandler(stream) == gnet.Shutdown {
				svr.Close()
			}
		})

	}
	return
}

package main

import (
	"bytes"
	"crypto/x509"
	"io/ioutil"
	"jachunPM_http/config"
	"jachunPM_http/handler"
	"libraries"
	"log"
	"net/http"
	_ "net/http/pprof"
	"protocol"
	"runtime"
	"server"
	"strconv"
	"time"

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
	handler.HostConn, err = protocol.NewClient(protocol.HttpServerNo, config.Server.HostIP, config.Server.TokenKey)
	if err != nil {
		libraries.ReleaseLog("服务启动失败%v", err)
	} else {
		//db.Init()
		handler.HostConn.HandleMsg = handler.Handler
		go handler.HostConn.Start()
	}
	for key, config := range config.Config[protocol.DefaultLang] {
		handler.HostConn.SetConfig(key, config)
	}

	go http.ListenAndServe("0.0.0.0:"+strconv.Itoa(8100+protocol.HttpServerNo), nil)
	svr := &httpServer{addr: config.Server.ListenHttp}
	// Start serving!
	var tlsconfig *tls.Config
	if config.Server.HttpsTLScert != "" && config.Server.HttpsTLSca != "" && config.Server.HttpsTLSkey != "" {
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
		server_cert, err := ioutil.ReadFile(config.Server.HttpsTLScert)
		if err != nil {
			log.Fatalf("读取服务器证书cert错误 err: %v", err)
		}
		server_key, err := ioutil.ReadFile(config.Server.HttpsTLSkey)
		if err != nil {
			log.Fatalf("读取服务器证书key错误 err: %v", err)
		}
		if ca, err := ioutil.ReadFile(config.Server.HttpsTLSca); err == nil {
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

		go func() {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				r.Header.Add("strict-transport-security", "max-age=31536000; includeSubDomains; preload")
				http.Redirect(w, r, config.Server.Origin, http.StatusFound)
			})
			http.ListenAndServe("0.0.0.0:80", nil)
		}()
	}

	log.Fatal(gnet.Serve(svr, svr.addr, gnet.WithLoopNum(runtime.NumCPU()), gnet.WithTCPKeepAlive(time.Second*600), gnet.WithCodec(&server.Tlscodec{}), gnet.WithReusePort(true), gnet.WithOutbuf(1024), gnet.WithTls(tlsconfig), gnet.WithMultiOut(false)))
}
func (hs *httpServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	libraries.DebugLog("httpserver started on %s (loops: %d)", hs.addr, srv.NumEventLoop)
	return
}

func (hs *httpServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	time.AfterFunc(time.Second*10, func() {
		if c.Context() == nil {
			c.Close()
		}
	})
	return
}

func (hs *httpServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	switch svr := c.Context().(type) {
	case *server.WSconn:

	case *server.Httpserver:
		svr.Close()
		server.Httppool.Put(svr)
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

func (hs *httpServer) React(data []byte, c gnet.Conn) (action gnet.Action) {

	switch svr := c.Context().(type) {
	case *server.Httpserver:
		action = handler.HttpHandler(svr)
		if svr.Request.Connection == "close" {
			action = gnet.Close
		}
		c.AsyncWrite(svr.Out.Bytes())
		svr.Recovery()
		return
	case *server.WSconn:

	case *server.Http2server:
		//svr.SendPool.Invoke(svr.WorkStream)

	}
	return
}

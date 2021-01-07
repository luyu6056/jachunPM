package main

import (
	"jachunPM_commom/config"
	"jachunPM_commom/db"
	"jachunPM_commom/rpcHost"
	"libraries"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"server"
	"sync"
	"time"

	"github.com/luyu6056/cache"
	"github.com/luyu6056/gnet"
)

func init() {
	libraries.SetLogLever(libraries.LogLeverDebug)
	rand.Seed(time.Now().UnixNano())

	db.Init()
	rpcHost.MsgnoInit()
	go rpcHost.HostServerHandlerMsgIn()
	go rpcHost.HostServerHandlerOutChan()
}

type rpcServer struct {
	*gnet.EventServer
	addr string
}

var Context_pool = sync.Pool{New: func() interface{} {
	return &server.Context{Buf: new(libraries.MsgBuffer), In: new(libraries.MsgBuffer), In2: new(libraries.MsgBuffer)}
}}

func main() {
	cache.StartWebServer("0.0.0.0:807")
	go http.ListenAndServe("0.0.0.0:8100", nil)
	rpcSvr := &rpcServer{addr: config.Config.ListenRpc}

	gnet.Serve(rpcSvr, rpcSvr.addr, gnet.WithLoopNum(runtime.NumCPU()*2), gnet.WithTCPKeepAlive(time.Second*600), gnet.WithCodec(&rpcHost.RpcCodec{}), gnet.WithReusePort(true), gnet.WithOutbuf(1024), gnet.WithMultiOut(false))
}

func (rs *rpcServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	libraries.DebugLog("rpchost started on %s (loops: %d)", rs.addr, srv.NumEventLoop)
	return
}
func (rs *rpcServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	time.AfterFunc(time.Second*10, func() {
		if c.Context() == nil {
			c.Close()
		}
	})
	return
}
func (rs *rpcServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	switch svr := c.Context().(type) {
	case *rpcHost.RpcServer:
		svr.Close()
	}
	c.SetContext(nil)
	return gnet.None
}
func (rs *rpcServer) React(data []byte, c gnet.Conn) (action gnet.Action) {
	ctx := Context_pool.Get().(*server.Context)
	defer Context_pool.Put(ctx)
	ctx.In.ResetBuf(data)
	err := rpcHost.SendMsgToRemote(ctx, c)
	if err != nil {
		libraries.DebugLog("host处理消息失败%v", err)
		return gnet.Close
	}

	return gnet.None
}

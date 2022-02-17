package main

import (
	"jachunPM_commom/db"
	"jachunPM_commom/rpcHost"
	"jachunPM_commom/setting"
	"libraries"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"protocol"
	"runtime"
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

func main() {
	cache.StartWebServer("0.0.0.0:807")
	go http.ListenAndServe("0.0.0.0:8100", nil)
	rpcSvr := &rpcServer{addr: setting.Setting.ListenRpc}

	gnet.Serve(rpcSvr, rpcSvr.addr, gnet.WithLoopNum(runtime.NumCPU()), gnet.WithCodec(&protocol.RpcCodec{}), gnet.WithReusePort(false), gnet.WithOutbuf(256), gnet.WithMultiOut(true), gnet.WithTCPNoDelay(true))
}

func (rs *rpcServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	libraries.DebugLog("rpchost started on %s (loops: %d)", rs.addr, srv.NumEventLoop)
	return
}
func (rs *rpcServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(rpcHost.NewRpcServer(c)) //装载未注册消息
	time.AfterFunc(time.Second*10, func() {
		if v, ok := c.Context().(*rpcHost.RpcServer); ok && v.Id == -1 {
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
	err := rpcHost.HandlerMsg(data, c)
	if err != nil {
		libraries.DebugLog("host处理消息失败%v", err)
		return gnet.Close
	}

	return gnet.None
}

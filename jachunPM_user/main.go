package main

import (
	"jachunPM_user/config"
	"jachunPM_user/db"
	"jachunPM_user/handler"
	"libraries"
	"net/http"
	_ "net/http/pprof"
	"protocol"
	"strconv"
)

func main() {
	var err error
	handler.HostConn, err = protocol.NewClient(protocol.UserServerNo, config.Config.HostIP, config.Config.TokenKey)
	go func() {
		http.ListenAndServe("0.0.0.0:"+strconv.Itoa(8100+protocol.UserServerNo), nil)
	}()
	if err != nil {
		libraries.ReleaseLog("服务启动失败%v", err)
	} else {
		db.Init()
		handler.HostConn.HandleMsg = handler.Handler
		handler.HostConn.HandleTick = handler.HandleTick
		handler.HostConn.Start()
	}

}

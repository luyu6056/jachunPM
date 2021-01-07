package main

import (
	"jachunPM_新服务模板/config"
	"jachunPM_新服务模板/db"
	"jachunPM_新服务模板/handler"
	"libraries"
	"net/http"
	_ "net/http/pprof"
	"protocol"
	"strconv"
)

func main() {
	var err error
	handler.HostConn, err = protocol.NewClient(protocol.ProjectServerNo, config.Config.HostIP, config.Config.TokenKey)
	go func() {
		http.ListenAndServe("0.0.0.0:"+strconv.Itoa(8100+protocol.ProjectServerNo), nil)
	}()
	if err != nil {
		libraries.ReleaseLog("服务启动失败%v", err)
	} else {
		handler.HostConn.DB = db.Init()
		handler.HostConn.HandleMsg = handler.Handler
		handler.HostConn.HandleTick = handler.HandleTick
		handler.HostConn.Start()
	}

}

package main

import (
	"jachunPM_project/db"
	"jachunPM_project/handler"
	"jachunPM_project/setting"
	"libraries"
	"net/http"
	_ "net/http/pprof"
	"protocol"
	"strconv"
)

func main() {
	var err error
	handler.HostConn, err = protocol.NewClient(protocol.ProjectServerNo, setting.Setting.HostIP, setting.Setting.TokenKey)
	go func() {
		http.ListenAndServe("0.0.0.0:"+strconv.Itoa(8100+protocol.ProjectServerNo), nil)
	}()
	if err != nil {
		libraries.ReleaseLog("服务启动失败%v", err)
	} else {
		handler.HostConn.DB = db.Init()
		handler.HostConn.HandleMsg = handler.Handler
		handler.HostConn.SetTickHand(handler.HandleTick)
		handler.HostConn.Start()
	}

}

package main

import (
	"jachunPM_log/db"
	"jachunPM_log/handler"
	"jachunPM_log/setting"
	"libraries"
	"net/http"
	_ "net/http/pprof"
	"protocol"
	"strconv"
)

func main() {

	var err error
	handler.HostConn, err = protocol.NewClient(protocol.LogServerNo, setting.Setting.HostIP, setting.Setting.TokenKey)
	go func() {
		http.ListenAndServe("0.0.0.0:"+strconv.Itoa(8100+protocol.LogServerNo), nil)
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

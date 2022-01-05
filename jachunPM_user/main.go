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
	go func() {
		http.ListenAndServe("0.0.0.0:"+strconv.Itoa(8100+protocol.UserServerNo), nil)
	}()
	var err error
	DB := db.Init()
	handler.HostConn, err = protocol.NewClient(protocol.UserServerNo, config.Config.HostIP, config.Config.TokenKey)

	if err != nil {
		libraries.ReleaseLog("服务启动失败%v", err)
	} else {
		handler.HostConn.DB = DB
		handler.HostConn.HandleMsg = handler.Handler
		handler.HostConn.SetTickHand(handler.HandleTick)
		handler.HostConn.Start()
	}

}

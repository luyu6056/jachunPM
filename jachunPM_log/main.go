package main

import (
	"jachunPM_log/config"
	"jachunPM_log/handler"
	"libraries"
	"protocol"
)

func main() {
	var err error
	handler.HostConn, err = protocol.NewClient(protocol.LogServerNo, config.Config.HostIP, config.Config.TokenKey)
	if err != nil {
		libraries.ReleaseLog("服务启动失败%v", err)
	} else {
		handler.HostConn.HandleMsg = handler.Handler
		handler.HostConn.Start()
	}

}

package handler

import (
	"jachunPM_http/config"
	"protocol"
	"time"
)

func HandleTick(t time.Time) {
	firstFlag := protocol.RpcTickStatusFirst
	if HostConn.Status&firstFlag == firstFlag {
		HostConn.Status -= protocol.RpcTickStatusFirst
		for key, config := range config.Config[protocol.DefaultLang] {
			HostConn.SetConfig(key, config)
		}
	} else {

	}

}

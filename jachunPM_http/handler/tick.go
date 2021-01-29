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
		for lang := range protocol.CountryToStr {
			for key, config := range config.Config[lang] {
				HostConn.SetConfig(lang, key, config)
			}
		}

	} else {

	}

}

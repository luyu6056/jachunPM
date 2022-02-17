package handler

import (
	"protocol"
	"time"
)

func HandleTick(t time.Time) {
	firstFlag := protocol.RpcTickStatusFirst
	if HostConn.Status&firstFlag == firstFlag {
		HostConn.Status -= protocol.RpcTickStatusFirst

	} else {

	}

}

package db

import (
	"fmt"
	"libraries"
	"protocol"
	"runtime/debug"
	"sync"
	"time"
)

const (
	maxUpdateNum = 100
)

var (
	updatePool_log_msg = sync.Pool{New: func() interface{} {
		return &Log_msg{}
	}}
	updatePool_ZstdMsg = sync.Pool{New: func() interface{} {
		return &ZstdMsg{}
	}}
	updateChan_log_msg = make(chan *Log_msg, maxUpdateNum)
	updateChan_ZstdMsg= make(chan *ZstdMsg, maxUpdateNum)
)

func MsgtoLog(msg *protocol.Msg, logs []*Log_msg) []*Log_msg {
	var lastTTl int
	if len(logs) > 0 {
		lastTTl = logs[len(logs)-1].Ttl
	}
	log := updatePool_log_msg.Get().(*Log_msg)
	log.Cmd = protocol.CmdToName[msg.Cmd]
	log.LocalNo = uint8(msg.Local)
	log.LocalId = uint8(msg.Local >> 8)
	log.Msgno = msg.Msgno
	log.TimeOut=msg.TtlTimeout
	log.RemoteNo = uint8(msg.GetRemoteID())
	log.RemoteId = uint8(msg.GetRemoteID() >> 8)
	log.Timestamp = time.Now()
	log.Err = ""
	log.Ttl = lastTTl + 1
	if msg.Cmd == protocol.CMD_MSG_HOST_QueryErr {
		msg.ReadDataWithCopy()
		if data, ok := msg.Data.(*protocol.MSG_HOST_QueryErr); ok {
			log.Err = data.Err
			log.Stack = string(data.Stack)
		}
	}
	//大于一定的ttl阈值才保存，避免每条消息都存
	if log.Ttl > 50 {
		for _, l := range logs[1:] {
			updateChan_log_msg <- l
		}
		updateChan_log_msg <- log
		return []*Log_msg{log}
	} else {
		logs = append(logs, log)
	}
	return logs
}
func ToZstd(b []byte,cmd int32) {
	zstd:=updatePool_ZstdMsg.Get().(*ZstdMsg)
	zstd.Cmd=cmd
	zstd.Name=protocol.CmdToName[cmd]
	zstd.Msg=make([]byte,len(b))
	copy(zstd.Msg,b)
	updateChan_ZstdMsg<-zstd
}
func UpdatedbInit() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()

		}
		go UpdatedbInit()
	}()
	var (
		update_log_msg = make([]*Log_msg, 0)
		update_ZstdMsg =make([]*ZstdMsg,0)
	)
	for {
		select {
		case log := <-updateChan_log_msg:
			update_log_msg = append(update_log_msg, log)
			for i := 0; i < len(updateChan_log_msg); i++ {
				log := <-updateChan_log_msg
				update_log_msg = append(update_log_msg, log)
			}
			_, err := DB.Table(TABLE_LOG_MSG).ReplaceAll(update_log_msg)
			if err != nil {
				libraries.ReleaseLog("插入LOG_MSG失败%v", err)
			}
			for _, v := range update_log_msg {
				updatePool_log_msg.Put(v)
			}
			update_log_msg = update_log_msg[:0]
			case zstd:=<-updateChan_ZstdMsg:

				update_ZstdMsg = append(update_ZstdMsg, zstd)
				for i := 0; i < len(updateChan_ZstdMsg); i++ {
					zstd := <-updateChan_ZstdMsg
					update_ZstdMsg = append(update_ZstdMsg, zstd)
				}

				for _,v:= range update_ZstdMsg{
					v.Sha=libraries.SHA256_S(string(v.Msg))
					 err := DB.Table("zstd").Where("count(Cmd) < 1000").Replace(v)
					if err != nil {
						libraries.ReleaseLog("插入zstd失败%v", err)
					}
				}


				for _, v := range update_ZstdMsg {
					v.Msg=v.Msg[:0]
					updatePool_ZstdMsg.Put(v)
				}
				update_ZstdMsg = update_ZstdMsg[:0]
		}
	}
}

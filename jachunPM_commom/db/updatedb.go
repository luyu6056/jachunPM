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
	updateChan_log_msg = make(chan *Log_msg, maxUpdateNum)
)

func WriteMsgLog(msg *protocol.Msg) {
	log := updatePool_log_msg.Get().(*Log_msg)
	log.Cmd = protocol.CmdToName[msg.Cmd]
	log.LocalNo = uint8(msg.Local)
	log.LocalId = uint8(msg.Local >> 8)
	log.Msgno = msg.Msgno
	log.RemoteNo = uint8(msg.Remote)
	log.RemoteId = uint8(msg.Remote >> 8)
	log.Timestamp = time.Now()
	log.Err = ""
	if msg.Ttl >= protocol.MaxMsgTtl {
		msg.Ttl = 255
	}
	log.Ttl = msg.Ttl
	if msg.Cmd == protocol.CMD_MSG_COMMON_QueryErr {
		msg.ReadDataWithCopy()
		if data, ok := msg.Data.(*protocol.MSG_COMMON_QueryErr); ok {
			log.Err = data.Err
			log.Stack = string(data.Stack)
		}
	}
	updateChan_log_msg <- log
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
		}
	}
}

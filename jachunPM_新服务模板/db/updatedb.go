package db

import (
	"fmt"
	"libraries"
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

func WriteMsgLog(msgno uint32, ttl uint8, local, remote uint16, cmd int32) {
	log := updatePool_log_msg.Get().(*Log_msg)
	log.Cmd = cmd
	log.LocalNo = uint8(local)
	log.LocalId = uint8(local >> 8)
	log.Msgno = msgno
	log.RemoteNo = uint8(remote)
	log.RemoteId = uint8(remote >> 8)
	log.Timestamp = time.Now()
	log.Ttl = ttl
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
			_, err := DB.Table(TABLE_LOG_MSG).InsertAll(update_log_msg)
			if err != nil {
				libraries.ReleaseLog("插入LOG_MSG失败", err)
			}
			for _, v := range update_log_msg {
				updatePool_log_msg.Put(v)
			}
			update_log_msg = update_log_msg[:0]
		}
	}
}

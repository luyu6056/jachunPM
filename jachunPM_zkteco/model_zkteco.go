package handler

import (
	"errors"
	"fmt"
	"libraries"
	"net"
	"strings"
	"sync"
	"time"
	"unsafe"
)

//中控考勤机相关

var (

	CMD_CONNECT         = []byte{0xe8, 0x03} //开始连接。
	CMD_EXIT            = []byte{0xe9, 0x03} //断开。
	CMD_ENABLEDEVICE    = []byte{0xea, 0x03} //将机器状态更改为“正常工作”。
	CMD_DISABLEDEVICE   = []byte{0xeb, 0x03} //禁用指纹、rfid 阅读器和键盘。
	CMD_RESTART         = []byte{0xec, 0x03} //重启机器。
	CMD_POWEROFF        = []byte{0xed, 0x03} //关闭机器。
	CMD_SLEEP           = []byte{0xee, 0x03} //将机器状态更改为“空闲”。
	CMD_RESUME          = []byte{0xef, 0x03} //将机器状态更改为“唤醒”。
	CMD_CAPTUREFINGER   = []byte{0xf1, 0x03} //采集指纹图片。
	CMD_TEST_TEMP       = []byte{0xf3, 0x03} //测试指纹是否存在。
	CMD_CAPTUREIMAGE    = []byte{0xf4, 0x03} //捕获整个图像。
	CMD_REFRESHDATA     = []byte{0xf5, 0x03} //刷新机器存储的数据。
	CMD_REFRESHOPTION   = []byte{0xf6, 0x03} //刷新配置参数。
	CMD_TESTVOICE       = []byte{0xf9, 0x03} //测试语音。
	CMD_GET_VERSION     = []byte{0x4c, 0x04} //请求固件版本。
	CMD_CHANGE_SPEED    = []byte{0x4d, 0x04} //改变传输速度。
	CMD_AUTH            = []byte{0x4e, 0x04} //使用 commkey 请求开始会话。
	CMD_PREPARE_DATA    = []byte{0xdc, 0x05} //准备数据传输。
	CMD_DATA            = []byte{0xdd, 0x05} //数据包。
	CMD_FREE_DATA       = []byte{0xde, 0x05} //释放缓冲区用于数据传输。
	CMD_DATA_WRRQ       = []byte{0xdf, 0x05} //读取/写入大型数据集。
	CMD_DATA_RDY        = []byte{0xe0, 0x05} //表示已准备好接收数据。
	CMD_DB_RRQ          = []byte{0x07, 0x00} //读取保存的数据。
	CMD_USER_WRQ        = []byte{0x08, 0x00} //上传用户数据。
	CMD_USERTEMP_RRQ    = []byte{0x09, 0x00} //读取用户指纹模板。
	CMD_USERTEMP_WRQ    = []byte{0x0a, 0x00} //上传用户指纹模板。
	CMD_OPTIONS_RRQ     = []byte{0x0b, 0x00} //读取机器的配置值。
	CMD_OPTIONS_WRQ     = []byte{0x0c, 0x00} //更改机器的配置值。
	CMD_ATTLOG_RRQ      = []byte{0x0d, 0x00} //请求考勤日志。
	CMD_CLEAR_DATA      = []byte{0x0e, 0x00} //删除数据。
	CMD_CLEAR_ATTLOG    = []byte{0x0f, 0x00} //删除考勤记录。
	CMD_DELETE_USER     = []byte{0x12, 0x00} //删除用户。
	CMD_DELETE_USERTEMP = []byte{0x13, 0x00} //删除用户指纹模板。
	CMD_CLEAR_ADMIN     = []byte{0x14, 0x00} //清除管理员权限。
	CMD_USERGRP_RRQ     = []byte{0x15, 0x00} //读取用户组。
	CMD_USERGRP_WRQ     = []byte{0x16, 0x00} //设置用户组。
	CMD_USERTZ_RRQ      = []byte{0x17, 0x00} //获取用户时区。
	CMD_USERTZ_WRQ      = []byte{0x18, 0x00} //设置用户时区。
	CMD_GRPTZ_RRQ       = []byte{0x19, 0x00} //获取组时区。
	CMD_GRPTZ_WRQ       = []byte{0x1a, 0x00} //设置组时区。
	CMD_TZ_RRQ          = []byte{0x1b, 0x00} //获取设备时区。
	CMD_TZ_WRQ          = []byte{0x1c, 0x00} //设置设备时区。
	CMD_ULG_RRQ         = []byte{0x1d, 0x00} //获取组组合解锁。
	CMD_ULG_WRQ         = []byte{0x1e, 0x00} //设置组组合解锁。
	CMD_UNLOCK          = []byte{0x1f, 0x00} //在指定的时间内解锁门。
	CMD_CLEAR_ACC       = []byte{0x20, 0x00} //将访问控制恢复为默认值。
	CMD_CLEAR_OPLOG     = []byte{0x21, 0x00} //删除操作日志。
	CMD_OPLOG_RRQ       = []byte{0x22, 0x00} //读取操作日志。
	CMD_GET_FREE_SIZES  = []byte{0x32, 0x00} //请求机器状态（剩余空间）。
	CMD_ENABLE_CLOCK    = []byte{0x39, 0x00} //启用屏幕时钟中的“：”。
	CMD_STARTVERIFY     = []byte{0x3c, 0x00} //将机器设置为认证状态。
	CMD_STARTENROLL     = []byte{0x3d, 0x00} //开始注册程序。
	CMD_CANCELCAPTURE   = []byte{0x3e, 0x00} //关闭用户的正常认证。
	CMD_STATE_RRQ       = []byte{0x40, 0x00} //查询状态。
	CMD_WRITE_LCD       = []byte{0x42, 0x00} //将字符打印到设备屏幕。
	CMD_CLEAR_LCD       = []byte{0x43, 0x00} //清除屏幕字幕。
	CMD_GET_PINWIDTH    = []byte{0x45, 0x00} //请求用户 ID 的最大大小。
	CMD_SMS_WRQ         = []byte{0x46, 0x00} //上传短消息。
	CMD_SMS_RRQ         = []byte{0x47, 0x00} //下载短消息。
	CMD_DELETE_SMS      = []byte{0x48, 0x00} //删除短消息。
	CMD_UDATA_WRQ       = []byte{0x49, 0x00} //设置用户短信。
	CMD_DELETE_UDATA    = []byte{0x4a, 0x00} //删除用户短信。
	CMD_DOORSTATE_RRQ   = []byte{0x4b, 0x00} //获取门状态。
	CMD_WRITE_MIFARE    = []byte{0x4c, 0x00} //将数据写入 Mifare 卡。
	CMD_EMPTY_MIFARE    = []byte{0x4e, 0x00} //清除 Mifare 卡。
	CMD_VERIFY_WRQ      = []byte{0x4f, 0x00} //更改给定用户的验证方式。
	CMD_VERIFY_RRQ      = []byte{0x50, 0x00} //读取给定用户的验证方式。
	CMD_TMP_WRITE       = []byte{0x57, 0x00} //从缓冲区传输 fp 模板。
	CMD_CHECKSUM_BUFFER = []byte{0x77, 0x00} //获取机器缓冲区的校验和。
	CMD_DEL_FPTMP       = []byte{0x86, 0x00} //删除指纹模板。
	CMD_GET_TIME        = []byte{0xc9, 0x00} //请求机器时间。
	CMD_SET_TIME        = []byte{0xca, 0x00} //设置机器时间。
	CMD_REG_EVENT       = []byte{0xf4, 0x01} //实时事件。
)

const (
	packetStart         = uint32(0x7d825050)
	CMD_ACK_OK         = 0x07d0 //请求已成功处理。
	CMD_ACK_ERROR      = 0x07d1 //处理请求时出错。
	CMD_ACK_DATA       = 0x07d2
	CMD_ACK_RETRY      = 0x07d3
	CMD_ACK_REPEAT     = 0x07d4
	CMD_ACK_UNAUTH     = 0x07d5 //未授权连接。
	CMD_ACK_UNKNOWN    = 0xffff //收到未知命令。
	CMD_ACK_ERROR_CMD  = 0xfffd
	CMD_ACK_ERROR_INIT = 0xfffc
	CMD_ACK_ERROR_DATA = 0xfffb
	CMD_ACK_REG_EVENT  = 0x01f4
)
const (
	EF_ATTLOG       = 1   //0x1 考勤录入。
	EF_FINGER       = 2   //0x2 手指按下。
	EF_ENROLLUSER   = 4   //0x4 注册用户。
	EF_ENROLLFINGER = 8   //0x8 登记指纹。
	EF_BUTTON       = 16  //0x10 按下键盘键。
	EF_UNLOCK       = 32  //0x20 无
	EF_VERIFY       = 128 //0x80 注册用户放置手指。
	EF_FPFTR        = 256 //0x100 登记过程中的指纹分数。
	EF_ALARM        = 512 //0x200 触发警报。
)

var ackToStr = map[uint16]string{
	0x07d0: "CMD_ACK_OK",    //请求已成功处理。
	0x07d1: "CMD_ACK_ERROR", //处理请求时出错。
	0x07d2: "CMD_ACK_DATA",
	0x07d3: "CMD_ACK_RETRY",
	0x07d4: "CMD_ACK_REPEAT",
	0x07d5: "CMD_ACK_UNAUTH",  //未授权连接。
	0xffff: "CMD_ACK_UNKNOWN", //收到未知命令。
	0xfffd: "CMD_ACK_ERROR_CMD",
	0xfffc: "CMD_ACK_ERROR_INIT",
	0xfffb: "CMD_ACK_ERROR_DATA",
	0x01f4: "CMD_REG_EVENT",
}

type zkecoConn struct {
	conn         net.Conn
	Msg          []byte
	Body         []byte
	checksumData []byte
	lock         sync.Mutex
	inbuffer     *libraries.MsgBuffer
	status       zkstatus
	Reply_number uint16
	Session_id   uint16
	replay       chan *msg
	err          error
}
type zkstatus int

const (
	none zkstatus = iota
	authok
	close
)

type msg struct {
	Command_id   uint16
	Checksum     uint16
	Session_id   uint16
	Reply_number uint16
	Data         []byte
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
var zktecoInfo *zkecoConn
func zktecoStart(ipaddr string)  {
	conn, err := net.Dial("tcp4", ipaddr)
	z := &zkecoConn{conn: conn, Msg: make([]byte, 8, 65535), Body: make([]byte, 8, 65535), inbuffer: &libraries.MsgBuffer{}, replay: make(chan *msg)}
	if err != nil {
		z.err = errors.New("无法连接考勤机,错误 " + err.Error())
		return z
	}

	z.Msg[0], z.Msg[1], z.Msg[2], z.Msg[3] = byte(packetStart), byte(packetStart>>8), byte(packetStart>>16), byte(packetStart>>24)
	if err = z.CONNECT(); err != nil {
		z.err = errors.New("无法连接考勤机,错误 " + err.Error())
		return z
	}
	z.status = authok
	go z.handler()
	return z
}
func (z *zkecoConn) Close() {
	z.conn.Close()
}
func (z *zkecoConn) CONNECT() (err error) {
	m, err := z.Write(CMD_CONNECT, nil)
	if err != nil {
		return err
	}
	z.Session_id = m.Session_id
	if m.Command_id == CMD_ACK_UNAUTH {
		a := auth(int32(m.Session_id))
		if m, err = z.Write(CMD_AUTH, a[:]); err != nil {
			return err
		} else if m.Command_id != CMD_ACK_OK {
			return errors.New("连接失败，auth失败")
		}
	}
	//z.Write(CMD_REG_EVENT, []byte{0xff, 0xff, 00, 00})
	z.Write(CMD_ENABLEDEVICE, nil)
	return nil
}
func (z *zkecoConn) Write(cmd []byte, data []byte) (m *msg, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()
	z.Msg = z.Msg[:8]
	z.Body = z.Body[:8]
	copy(z.Body, cmd)
	z.Body[4], z.Body[5] = byte(z.Session_id), byte(z.Session_id>>8)
	z.Body[6], z.Body[7] = byte(z.Reply_number), byte((z.Reply_number)>>8)
	z.Body = append(z.Body, data...)
	z.checksum()
	l := len(z.Body)
	z.Msg[4], z.Msg[5], z.Msg[6], z.Msg[7] = byte(l), byte(l>>8), byte(l>>16), byte(l>>24)
	z.Msg = append(z.Msg, z.Body...)
	fmt.Print("发送")
	printMsg(z.Msg)
	z.conn.Write(z.Msg)
	z.Reply_number++
	select {
	case m = <-z.replay:
		return m, nil
	case <-time.After(time.Second * 10):
		return nil, errors.New("连接超时")
	}
}
func (z *zkecoConn) checksum() {
	z.checksumData = z.checksumData[:0]
	z.checksumData = append(z.checksumData, z.Body[:2]...)
	z.checksumData = append(z.checksumData, z.Body[4:]...)
	if len(z.checksumData)%2 == 1 {
		z.checksumData = append(z.checksumData, 0)
	}
	chk_32b := 0
	j := 1
	for j < len(z.checksumData) {
		num_16b := int(z.checksumData[j-1]) + int(z.checksumData[j])<<8
		chk_32b = chk_32b + num_16b
		j += 2
	}
	chk_32b = (chk_32b & 0xffff) + ((chk_32b & 0xffff0000) >> 16)
	chk_16b := chk_32b ^ 0xFFFF
	z.Body[2], z.Body[3] = byte(chk_16b), byte(chk_16b>>8)
}
func printMsg(b []byte) {
	var p []string
	for _, v := range b {
		s := fmt.Sprintf("%x", v)
		if len(s) == 1 {
			p = append(p, "0"+s)
		} else {
			p = append(p, s)
		}
	}
	fmt.Println(strings.Join(p, " "))
}
func (z *zkecoConn) handler() {
	defer func() {
		z.status = close
	}()
	for {
		b := make([]byte, 65535)
		n, err := z.conn.Read(b)
		if err != nil {
			//z.err = errors.New("读取错误" + err.Error())
			return
		}
		z.inbuffer.Write(b[:n])
		if z.inbuffer.Len() > 8 {
			if m := z.ReadMsg(); m != nil {
				fmt.Printf("收到%+v 消息:", ackToStr[m.Command_id])
				if m.Data != nil {
					printMsg(m.Data)
				}
				fmt.Println("")
				switch m.Command_id {
				case CMD_ACK_REG_EVENT:
					switch m.Session_id {
					case EF_ATTLOG:
						if len(m.Data) == 36 {
							name := string(m.Data[:9])
							date := fmt.Sprintf("20%d/%d/%d %d:%d:%d", m.Data[26], m.Data[27], m.Data[28], m.Data[29], m.Data[30], m.Data[31])
							fmt.Println(name, date)
						}
					}
				default:
					select {
					case z.replay <- m:
					default:

					}
				}
			}
		}

	}
}
func (z *zkecoConn) ReadMsg() *msg {

	ptr := uintptr(*(*unsafe.Pointer)(unsafe.Pointer(z.inbuffer)))
	if *(*uint32)(unsafe.Pointer(ptr)) != packetStart {
		z.err = errors.New("收到错误的消息")
		z.Close()
		return nil
	}
	msglen := int(*(*uint32)(unsafe.Pointer(ptr + 4))) + 8
	if z.inbuffer.Len() < msglen { //收到的消息不够一条长度
		return nil
	}
	m := &msg{ *(*uint16)(unsafe.Pointer(ptr + 8)), *(*uint16)(unsafe.Pointer(ptr + 10)), *(*uint16)(unsafe.Pointer(ptr + 12)), *(*uint16)(unsafe.Pointer(ptr + 14)), nil}
	if msglen > 16 {
		s := (*sliceHeader)(unsafe.Pointer(&m.Data))
		s.Data = unsafe.Pointer(ptr + 16)
		s.Len = msglen - 16
		s.Cap = msglen - 16
	}
	z.inbuffer.Shift(msglen)
	return m
}
func auth(a1 int32) (result [4]byte) {
	v2 := int32(1)
	v3 := int32(0)
	for i := 0; i < 32; i++ {
		v3 = 2 * v3
		if (v2 & a1) != 0 {
			v3 |= 1
		}
		v2 *= 16
	}

	v8 := a1 + v2
	TickCount := byte(3)
	result = *(*[4]byte)(unsafe.Pointer(&v8))
	v5 := result[1] ^ 0x4B
	result[0], result[1] = result[2]^0x53, result[3]^0x4f
	result[3] = v5 ^ TickCount
	result[0] = TickCount ^ result[0]
	result[1] ^= TickCount
	result[2] = TickCount
	return
}
func init() {

}

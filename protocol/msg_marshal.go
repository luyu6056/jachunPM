package protocol

import (
	"libraries"
	"unsafe"
)

func WRITE_int8(data int8, buf *libraries.MsgBuffer) {
	buf.WriteByte(byte(data))
}
func WRITE_int16(data int16, buf *libraries.MsgBuffer) {
	b := buf.Make(2)
	b[0] = byte(data)
	b[1] = byte(data >> 8)
}
func WRITE_int32(data int32, buf *libraries.MsgBuffer) {
	b := buf.Make(4)
	b[0] = byte(data)
	b[1] = byte(data >> 8)
	b[2] = byte(data >> 16)
	b[3] = byte(data >> 24)

}
func WRITE_int64(data int64, buf *libraries.MsgBuffer) {
	b := buf.Make(8)
	b[0] = byte(data)
	b[1] = byte(data >> 8)
	b[2] = byte(data >> 16)
	b[3] = byte(data >> 24)
	b[4] = byte(data >> 32)
	b[5] = byte(data >> 40)
	b[6] = byte(data >> 48)
	b[7] = byte(data >> 56)
}
func WRITE_uint(data uint, buf *libraries.MsgBuffer) {
	WRITE_int64(int64(data), buf)
}
func WRITE_int(data int, buf *libraries.MsgBuffer) {
	WRITE_int64(int64(data), buf)
}
func WRITE_uint8(data uint8, buf *libraries.MsgBuffer) {
	WRITE_int8(int8(data), buf)
}
func WRITE_uint16(data uint16, buf *libraries.MsgBuffer) {
	WRITE_int16(int16(data), buf)
}

func WRITE_uint32(data uint32, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(data), buf)
}
func WRITE_uint64(data uint64, buf *libraries.MsgBuffer) {
	WRITE_int64(int64(data), buf)
}
func WRITE_string(data string, buf *libraries.MsgBuffer) {
	length := len(data)
	b := buf.Make(4)
	b[0] = byte(length)
	b[1] = byte(length >> 8)
	b[2] = byte(length >> 16)
	b[3] = byte(length >> 24)
	x := (*[2]uintptr)(unsafe.Pointer(&data))
	h := [3]uintptr{x[0], x[1], x[1]}
	buf.Write(*(*[]byte)(unsafe.Pointer(&h)))
}
func WRITE_ErrCode(data ErrCode, buf *libraries.MsgBuffer) {
	WRITE_int16(int16(data), buf)
}
func WRITE_bool(b bool, buf *libraries.MsgBuffer) {
	if b {
		buf.WriteByte(1)
	} else {
		buf.WriteByte(0)
	}
}
func READ_int8(buf *libraries.MsgBuffer) int8 {
	b := buf.Next(1)
	if len(b) == 1 {
		return int8(b[0])
	}
	return 0
}
func READ_int16(buf *libraries.MsgBuffer) int16 {
	b := buf.Next(2)
	return int16(b[0]) | int16(b[1])<<8

}
func READ_int32(buf *libraries.MsgBuffer) int32 {
	b := buf.Next(4)
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}
func READ_int64(buf *libraries.MsgBuffer) int64 {
	b := buf.Next(8)
	return int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
}
func READ_int(buf *libraries.MsgBuffer) int {
	return int(READ_int64(buf))
}
func READ_uint(buf *libraries.MsgBuffer) uint {
	return uint(READ_int64(buf))
}
func READ_uint8(buf *libraries.MsgBuffer) uint8 {
	return uint8(READ_int8(buf))
}
func READ_uint16(buf *libraries.MsgBuffer) uint16 {
	return uint16(READ_int16(buf))
}
func READ_uint32(buf *libraries.MsgBuffer) uint32 {
	return uint32(READ_int32(buf))
}
func READ_uint64(buf *libraries.MsgBuffer) uint64 {
	return uint64(READ_int64(buf))
}

func READ_string(buf *libraries.MsgBuffer) string {
	b := buf.Next(4)
	length := int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | int(b[3])<<24
	if length > 0 {
		b := make([]byte, length)
		copy(b, buf.Next(length))
		return *(*string)(unsafe.Pointer(&b))
	}
	return ""
}
func READ_ErrCode(buf *libraries.MsgBuffer) ErrCode {
	return ErrCode(READ_int16(buf))
}
func READ_bool(buf *libraries.MsgBuffer) bool {
	b := buf.Next(1)
	return b[0] == 1
}
func READ_MSG_DATA(buf *libraries.MsgBuffer) MSG_DATA {
	cmd := READ_int32(buf)
	if f, ok := cmdMapFunc[cmd]; ok {
		return f(buf)
	}
	return nil
}

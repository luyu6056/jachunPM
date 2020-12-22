package protocol

import (
	"libraries"
	"reflect"
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
func WRITE_map(i interface{}, buf *libraries.MsgBuffer) {
	r := reflect.ValueOf(i)
	if r.Kind() != reflect.Map && r.Kind() != reflect.Invalid {
		panic("WRITE_map传入" + r.Kind().String())
	}
	if r.IsNil() || r.Len() == 0 {
		WRITE_uint32(0, buf)
		return
	} else if r.Len() < 65535 {
		WRITE_uint32(uint32(r.Len()), buf)
	} else {
		//这么大的map就不要传来传去或者分批传
		panic("WRITE_map写入len大于65535")
	}
	for _, key := range r.MapKeys() {
		v := r.MapIndex(key)
		write_reflect(key, buf)
		write_reflect(v, buf)
	}
}
func write_reflect(v reflect.Value, buf *libraries.MsgBuffer) {
	switch v.Kind() {
	case reflect.Int, reflect.Uint, reflect.Uint64, reflect.Int64:
		WRITE_int64(v.Int(), buf)
	case reflect.Int8:
		WRITE_int8(v.Interface().(int8), buf)
	case reflect.Int16:
		WRITE_int16(v.Interface().(int16), buf)
	case reflect.Int32:
		WRITE_int32(v.Interface().(int32), buf)
	case reflect.Uint8:
		WRITE_int8(int8(v.Interface().(uint8)), buf)
	case reflect.Uint16:
		WRITE_int16(int16(v.Interface().(uint16)), buf)
	case reflect.Uint32:
		WRITE_int32(int32(v.Interface().(uint32)), buf)
	case reflect.String:
		WRITE_string(v.String(), buf)
	case reflect.Bool:
		WRITE_bool(v.Bool(), buf)
	case reflect.Map:
		WRITE_map(v.Interface(), buf)
	case reflect.Slice:
		if vv, ok := v.Interface().([]byte); ok {
			WRITE_int32(int32(len(vv)), buf)
			buf.Write(vv)
		} else {
			WRITE_int32(int32(v.Len()), buf)
			for i := 0; i < v.Len(); i++ {
				write_reflect(v.Index(i), buf)
			}
		}
	default:

		panic("无法处理的map写入类型" + v.Kind().String())

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
func READ_map(i interface{}, buf *libraries.MsgBuffer) {
	l := READ_uint32(buf)
	if l == 0 {
		return
	}
	r := reflect.ValueOf(i)
	if r.Kind() != reflect.Ptr {
		panic("WRITE_map传入" + r.Kind().String())
	}
	r = r.Elem()
	r.Set(reflect.MakeMap(r.Type()))
	for i := 0; i < int(l); i++ {
		r.SetMapIndex(read_reflect(r.Type().Key(), buf), read_reflect(r.Type().Elem(), buf))
	}
}
func read_reflect(v reflect.Type, buf *libraries.MsgBuffer) reflect.Value {
	switch v.Kind() {
	case reflect.Int:
		return reflect.ValueOf(READ_int(buf))
	case reflect.Int64:
		return reflect.ValueOf(READ_int64(buf))
	case reflect.Int8:
		return reflect.ValueOf(READ_int8(buf))
	case reflect.Int16:
		return reflect.ValueOf(READ_int16(buf))
	case reflect.Int32:
		return reflect.ValueOf(READ_int32(buf))
	case reflect.Uint:
		return reflect.ValueOf(READ_uint(buf))
	case reflect.Uint8:
		return reflect.ValueOf(READ_uint8(buf))
	case reflect.Uint16:
		return reflect.ValueOf(READ_uint16(buf))
	case reflect.Uint32:
		return reflect.ValueOf(READ_uint32(buf))
	case reflect.Uint64:
		return reflect.ValueOf(READ_uint64(buf))
	case reflect.String:
		return reflect.ValueOf(READ_string(buf))
	case reflect.Bool:
		return reflect.ValueOf(READ_bool(buf))
	case reflect.Map:
		r := reflect.New(v).Elem()
		READ_map(r.Interface(), buf)
		return r
	case reflect.Slice:
		l := READ_int32(buf)
		if v.Elem().Kind() == reflect.Uint8 {
			b := make([]byte, l)
			copy(b, buf.Next(int(l)))
			return reflect.ValueOf(b)
		} else {
			r := reflect.MakeSlice(v, int(l), int(l))
			for i := 0; i < int(l); i++ {
				ii := r.Index(i)
				ii.Set(read_reflect(v.Elem(), buf))
			}
			return r
		}
	default:
		panic("无法处理的map读取类型" + v.Kind().String())
	}
}

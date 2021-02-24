package protocol

import (
	"libraries"
	"math"
	"reflect"
	"unsafe"

	"github.com/luyu6056/reflect2"
)

const (
	interfaceTypeInt = iota
	interfaceTypeInt8
	interfaceTypeInt16
	interfaceTypeInt32
	interfaceTypeInt64
	interfaceTypeUint
	interfaceTypeUint8
	interfaceTypeUint16
	interfaceTypeUint32
	interfaceTypeUint64
	interfaceTypeString
	interfaceTypeSliceInt
	interfaceTypeSliceInt8
	interfaceTypeSliceInt16
	interfaceTypeSliceInt32
	interfaceTypeSliceInt64
	interfaceTypeSliceUint
	interfaceTypeSliceUint8
	interfaceTypeSliceUint16
	interfaceTypeSliceUint32
	interfaceTypeSliceUint64
	interfaceTypeSliceString
	interfaceTypeBool
	interfaceTypeFloat32
	interfaceTypeFloat64
	interfaceTypeMss //map[string]string
	interfaceTypeMsi //map[string]interface{}
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
func WRITE_float32(f float32, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(math.Float32bits(f)), buf)
}
func WRITE_float64(f float64, buf *libraries.MsgBuffer) {
	WRITE_int64(int64(math.Float64bits(f)), buf)
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
func WRITE_any(i interface{}, buf *libraries.MsgBuffer) {
	switch v := i.(type) {
	case int:
		WRITE_int8(interfaceTypeInt, buf)
		WRITE_int(v, buf)
	case int8:
		WRITE_int8(interfaceTypeInt8, buf)
		WRITE_int8(v, buf)
	case int16:
		WRITE_int8(interfaceTypeInt16, buf)
		WRITE_int16(v, buf)
	case int32:
		WRITE_int8(interfaceTypeInt32, buf)
		WRITE_int32(v, buf)
	case int64:
		WRITE_int8(interfaceTypeInt64, buf)
		WRITE_int64(v, buf)
	case uint:
		WRITE_int8(interfaceTypeUint, buf)
		WRITE_uint(v, buf)
	case uint8:
		WRITE_int8(interfaceTypeUint8, buf)
		WRITE_uint8(v, buf)
	case uint16:
		WRITE_int8(interfaceTypeUint16, buf)
		WRITE_uint16(v, buf)
	case uint32:
		WRITE_int8(interfaceTypeUint32, buf)
		WRITE_uint32(v, buf)
	case uint64:
		WRITE_int8(interfaceTypeUint64, buf)
		WRITE_uint64(v, buf)
	case string:
		WRITE_int8(interfaceTypeString, buf)
		b := libraries.Str2bytes(v)
		WRITE_int32(int32(len(b)), buf)
		buf.Write(b)
	case []int:
		WRITE_int8(interfaceTypeSliceInt, buf)
		WRITE_int32(int32(len(v)), buf)
		for _, _v := range v {
			WRITE_int(_v, buf)
		}
	case []int8:
		WRITE_int8(interfaceTypeSliceInt8, buf)
		b := *(*[]byte)(unsafe.Pointer(&v))
		WRITE_int32(int32(len(b)), buf)
		buf.Write(b)
	case []int16:
		WRITE_int8(interfaceTypeSliceInt16, buf)
		WRITE_int32(int32(len(v)), buf)
		for _, _v := range v {
			WRITE_int16(_v, buf)
		}
	case []int32:
		WRITE_int8(interfaceTypeSliceInt32, buf)
		WRITE_int32(int32(len(v)), buf)
		for _, _v := range v {
			WRITE_int32(_v, buf)
		}
	case []int64:
		WRITE_int8(interfaceTypeSliceInt64, buf)
		WRITE_int32(int32(len(v)), buf)
		for _, _v := range v {
			WRITE_int64(_v, buf)
		}
	case []uint:
		WRITE_int8(interfaceTypeSliceUint, buf)
		WRITE_int32(int32(len(v)), buf)
		for _, _v := range v {
			WRITE_uint(_v, buf)
		}
	case []uint8:
		WRITE_int8(interfaceTypeSliceUint8, buf)
		b := []byte(v)
		WRITE_int32(int32(len(b)), buf)
		buf.Write(b)
	case []uint16:
		WRITE_int8(interfaceTypeSliceUint16, buf)
		WRITE_int32(int32(len(v)), buf)
		for _, _v := range v {
			WRITE_uint16(_v, buf)
		}
	case []uint32:
		WRITE_int8(interfaceTypeSliceUint32, buf)
		WRITE_int32(int32(len(v)), buf)
		for _, _v := range v {
			WRITE_uint32(_v, buf)
		}
	case []uint64:
		WRITE_int8(interfaceTypeSliceUint64, buf)
		WRITE_int32(int32(len(v)), buf)
		for _, _v := range v {
			WRITE_uint64(_v, buf)
		}
	case []string:
		WRITE_int8(interfaceTypeSliceString, buf)
		WRITE_int32(int32(len(v)), buf)
		for _, s := range v {
			b := libraries.Str2bytes(s)
			WRITE_int32(int32(len(b)), buf)
			buf.Write(b)
		}
	case bool:
		WRITE_int8(interfaceTypeBool, buf)
		WRITE_bool(v, buf)
	case float32:
		WRITE_int8(interfaceTypeFloat32, buf)
		WRITE_float32(v, buf)
	case float64:
		WRITE_int8(interfaceTypeFloat64, buf)
		WRITE_float64(v, buf)
	case map[string]string:
		WRITE_int8(interfaceTypeMss, buf)
		WRITE_map(v, buf)
	case map[string]interface{}:
		WRITE_int8(interfaceTypeMsi, buf)
		WRITE_map(v, buf)
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
	case reflect.Float32:
		WRITE_float32(v.Interface().(float32), buf)
	case reflect.Float64:
		WRITE_float64(v.Float(), buf)
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
	case reflect.Interface:
		WRITE_any(v.Interface(), buf)
	case reflect.Struct:
		switch i := v.Interface().(type) {
		case HtmlKeyValueStr:
			WRITE_HtmlKeyValueStr(i, buf)
		default:
			panic("无法处理的map写入Struct类型" + v.Type().Name())
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
func READ_float32(buf *libraries.MsgBuffer) float32 {
	return math.Float32frombits(READ_uint32(buf))
}
func READ_float64(buf *libraries.MsgBuffer) float64 {
	return math.Float64frombits(READ_uint64(buf))
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
func READ_any(i interface{}, buf *libraries.MsgBuffer) {
	uint_ptr := reflect2.PtrOf(i)
	t := READ_int8(buf)
	switch t {
	case interfaceTypeInt:
		*((*int)(unsafe.Pointer(uint_ptr))) = READ_int(buf)

	case interfaceTypeInt8:
		*((*int8)(unsafe.Pointer(uint_ptr))) = READ_int8(buf)
	case interfaceTypeInt16:
		*((*int16)(unsafe.Pointer(uint_ptr))) = READ_int16(buf)
	case interfaceTypeInt32:
		*((*int32)(unsafe.Pointer(uint_ptr))) = READ_int32(buf)
	case interfaceTypeInt64:
		*((*int64)(unsafe.Pointer(uint_ptr))) = READ_int64(buf)
	case interfaceTypeUint:
		*((*uint)(unsafe.Pointer(uint_ptr))) = READ_uint(buf)
	case interfaceTypeUint8:
		*((*uint8)(unsafe.Pointer(uint_ptr))) = READ_uint8(buf)
	case interfaceTypeUint16:
		*((*uint16)(unsafe.Pointer(uint_ptr))) = READ_uint16(buf)
	case interfaceTypeUint32:
		*((*uint32)(unsafe.Pointer(uint_ptr))) = READ_uint32(buf)
	case interfaceTypeUint64:
		*((*uint64)(unsafe.Pointer(uint_ptr))) = READ_uint64(buf)
	case interfaceTypeString:
		l := READ_int32(buf)
		b := buf.Next(int(l))
		*((*string)(unsafe.Pointer(uint_ptr))) = string(b)
	case interfaceTypeSliceInt:
		l := READ_int32(buf)
		s := make([]int, l)
		for k := range s {
			s[k] = READ_int(buf)
		}
		*((*[]int)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceInt8:
		l := READ_int32(buf)
		s := make([]byte, l)
		copy(s, buf.Next(int(l)))
		*((*[]byte)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceInt16:
		l := READ_int32(buf)
		s := make([]int16, l)
		for k := range s {
			s[k] = READ_int16(buf)
		}
		*((*[]int16)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceInt32:
		l := READ_int32(buf)
		s := make([]int32, l)
		for k := range s {
			s[k] = READ_int32(buf)
		}
		*((*[]int32)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceInt64:
		l := READ_int32(buf)
		s := make([]int64, l)
		for k := range s {
			s[k] = READ_int64(buf)
		}
		*((*[]int64)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint:
		l := READ_int32(buf)
		s := make([]uint, l)
		for k := range s {
			s[k] = READ_uint(buf)
		}
		*((*[]uint)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint8:
		l := READ_int32(buf)
		s := make([]uint8, l)
		for k := range s {
			s[k] = READ_uint8(buf)
		}
		*((*[]uint8)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint16:
		l := READ_int32(buf)
		s := make([]uint16, l)
		for k := range s {
			s[k] = READ_uint16(buf)
		}
		*((*[]uint16)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint32:
		l := READ_int32(buf)
		s := make([]uint32, l)
		for k := range s {
			s[k] = READ_uint32(buf)
		}
		*((*[]uint32)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint64:
		l := READ_int32(buf)
		s := make([]uint64, l)
		for k := range s {
			s[k] = READ_uint64(buf)
		}
		*((*[]uint64)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceString:
		l := READ_int32(buf)
		s := make([]string, l)
		for k := range s {
			_l := READ_int32(buf)
			b := buf.Next(int(_l))
			s[k] = string(b)
		}
		*((*[]string)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeBool:
		*((*bool)(unsafe.Pointer(uint_ptr))) = READ_bool(buf)
	case interfaceTypeFloat32:
		*((*float32)(unsafe.Pointer(uint_ptr))) = READ_float32(buf)
	case interfaceTypeFloat64:
		*((*float64)(unsafe.Pointer(uint_ptr))) = READ_float64(buf)
	case interfaceTypeMss:
		var m map[string]string
		READ_map(&m, buf)
		*((*map[string]string)(unsafe.Pointer(uint_ptr))) = m
	case interfaceTypeMsi:
		var m map[string]interface{}
		READ_map(&m, buf)
		*((*map[string]interface{})(unsafe.Pointer(uint_ptr))) = m
	}

}
func read_any_result(buf *libraries.MsgBuffer) interface{} {
	t := READ_int8(buf)
	switch t {
	case interfaceTypeInt:
		return READ_int(buf)
	case interfaceTypeInt8:
		return READ_int8(buf)
	case interfaceTypeInt16:
		return READ_int16(buf)
	case interfaceTypeInt32:
		return READ_int32(buf)
	case interfaceTypeInt64:
		return READ_int64(buf)
	case interfaceTypeUint:
		return READ_uint(buf)
	case interfaceTypeUint8:
		return READ_uint8(buf)
	case interfaceTypeUint16:
		return READ_uint16(buf)
	case interfaceTypeUint32:
		return READ_uint32(buf)
	case interfaceTypeUint64:
		return READ_uint64(buf)
	case interfaceTypeString:
		l := READ_int32(buf)
		b := buf.Next(int(l))
		return string(b)
	case interfaceTypeSliceInt:
		l := READ_int32(buf)
		s := make([]int, l)
		for k := range s {
			s[k] = READ_int(buf)
		}
		return s
	case interfaceTypeSliceInt8:
		l := READ_int32(buf)
		s := make([]byte, l)
		copy(s, buf.Next(int(l)))
		return s
	case interfaceTypeSliceInt16:
		l := READ_int32(buf)
		s := make([]int16, l)
		for k := range s {
			s[k] = READ_int16(buf)
		}
		return s
	case interfaceTypeSliceInt32:
		l := READ_int32(buf)
		s := make([]int32, l)
		for k := range s {
			s[k] = READ_int32(buf)
		}
		return s
	case interfaceTypeSliceInt64:
		l := READ_int32(buf)
		s := make([]int64, l)
		for k := range s {
			s[k] = READ_int64(buf)
		}
		return s
	case interfaceTypeSliceUint:
		l := READ_int32(buf)
		s := make([]uint, l)
		for k := range s {
			s[k] = READ_uint(buf)
		}
		return s
	case interfaceTypeSliceUint8:
		l := READ_int32(buf)
		s := make([]uint8, l)
		for k := range s {
			s[k] = READ_uint8(buf)
		}
		return s
	case interfaceTypeSliceUint16:
		l := READ_int32(buf)
		s := make([]uint16, l)
		for k := range s {
			s[k] = READ_uint16(buf)
		}
		return s
	case interfaceTypeSliceUint32:
		l := READ_int32(buf)
		s := make([]uint32, l)
		for k := range s {
			s[k] = READ_uint32(buf)
		}
		return s
	case interfaceTypeSliceUint64:
		l := READ_int32(buf)
		s := make([]uint64, l)
		for k := range s {
			s[k] = READ_uint64(buf)
		}
		return s
	case interfaceTypeSliceString:
		l := READ_int32(buf)
		s := make([]string, l)
		for k := range s {
			_l := READ_int32(buf)
			b := buf.Next(int(_l))
			s[k] = string(b)
		}
		return s
	case interfaceTypeBool:
		return READ_bool(buf)
	case interfaceTypeFloat32:
		return READ_float32(buf)
	case interfaceTypeFloat64:
		return READ_float64(buf)
	case interfaceTypeMss:
		var m map[string]string
		READ_map(&m, buf)
		return m
	case interfaceTypeMsi:
		var m map[string]interface{}
		READ_map(&m, buf)
		return m
	}
	return nil
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
	case reflect.Float32:
		return reflect.ValueOf(READ_float32(buf))
	case reflect.Float64:
		return reflect.ValueOf(READ_float64(buf))
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
	case reflect.Interface:
		return reflect.ValueOf(read_any_result(buf))
	case reflect.Struct:
		switch v.Name() {
		case "HtmlKeyValueStr":
			return reflect.ValueOf(READ_HtmlKeyValueStr(buf))
		default:
			panic("无法处理的map读取Struct类型" + v.Name())
		}
	default:
		panic("无法处理的map读取类型" + v.Kind().String())
	}
}

//特殊增加的
func WRITE_HtmlKeyValueStr(kv HtmlKeyValueStr, buf *libraries.MsgBuffer) {
	WRITE_string(kv.Key, buf)
	WRITE_string(kv.Value, buf)
}
func READ_HtmlKeyValueStr(buf *libraries.MsgBuffer) HtmlKeyValueStr {
	return HtmlKeyValueStr{
		Key:   READ_string(buf),
		Value: READ_string(buf),
	}
}

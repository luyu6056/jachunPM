package protocol

import (
	"fmt"
	"libraries"
	"math"
	"reflect"
	"runtime/debug"
	"unsafe"

	"github.com/luyu6056/reflect2"
)

const (
	MaxVarintLen16 = 3
	MaxVarintLen32 = 5
	MaxVarintLen64 = 10
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
	interfaceTypeSliceInterface
	interfaceTypeSliceKVs
	interfaceTypeBool
	interfaceTypeFloat32
	interfaceTypeFloat64
	interfaceTypeMss //map[string]string
	interfaceTypeMsi //map[string]interface{}
	interfaceTypeKVs //HtmlKeyValueStr
	interfaceTypeKVi //HtmlKeyValueInterface
	//深度遍历反射类型
	interfaceTypeM //map
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
func WRITE_int64(x int64, buf *libraries.MsgBuffer) {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	WRITE_uint64(ux, buf)
}
func WRITE_uint(data uint, buf *libraries.MsgBuffer) {
	WRITE_uint64(uint64(data), buf)
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
func WRITE_uint64(x uint64, buf *libraries.MsgBuffer) {
	var b []byte
	if x>>32 > 0 {
		b = buf.Make(MaxVarintLen64)
	} else if x>>16 > 0 {
		b = buf.Make(MaxVarintLen32)
	} else {
		b = buf.Make(MaxVarintLen16)
	}
	i := 0
	for x >= 0x80 {
		b[i] = byte(x) | 0x80
		x >>= 7
		i++
	}
	b[i] = byte(x)
	buf.Truncate(buf.Len() - (len(b) - i - 1))
}
func WRITE_string(data string, buf *libraries.MsgBuffer) {
	WRITE_int(len(data), buf)
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
	b := buf.Make(4)
	u := math.Float32bits(f)
	b[0] = byte(u)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
	b[3] = byte(u >> 24)
}
func WRITE_float64(f float64, buf *libraries.MsgBuffer) {
	b := buf.Make(8)
	u := math.Float64bits(f)
	b[0] = byte(u)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
	b[3] = byte(u >> 24)
	b[4] = byte(u >> 32)
	b[5] = byte(u >> 40)
	b[6] = byte(u >> 48)
	b[7] = byte(u >> 56)
}

func WRITE_map(i interface{}, buf *libraries.MsgBuffer) {
	r := reflect.ValueOf(i)
	if r.Kind() != reflect.Map && r.Kind() != reflect.Invalid {
		panic("WRITE_map传入" + r.Kind().String())
	}
	if r.IsNil() || r.Len() == 0 {
		WRITE_int(0, buf)
		return
	} else if r.Len() < 65535 {
		WRITE_int(r.Len(), buf)
	} else {
		//这么大的map就不要传来传去或者分批传
		panic("WRITE_map写入len大于65535")
	}
	_r := r.MapRange()
	for _r.Next() {
		WRITE_any(_r.Key().Interface(), buf)
		WRITE_any(_r.Value().Interface(), buf)
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
		WRITE_int(len(b), buf)
		buf.Write(b)
	case []int:
		WRITE_int8(interfaceTypeSliceInt, buf)
		WRITE_int(len(v), buf)
		for _, _v := range v {
			WRITE_int(_v, buf)
		}
	case []int8:
		WRITE_int8(interfaceTypeSliceInt8, buf)
		b := *(*[]byte)(unsafe.Pointer(&v))
		WRITE_int(len(b), buf)
		buf.Write(b)
	case []int16:
		WRITE_int8(interfaceTypeSliceInt16, buf)
		WRITE_int(len(v), buf)
		for _, _v := range v {
			WRITE_int16(_v, buf)
		}
	case []int32:
		WRITE_int8(interfaceTypeSliceInt32, buf)
		WRITE_int(len(v), buf)
		for _, _v := range v {
			WRITE_int32(_v, buf)
		}
	case []int64:
		WRITE_int8(interfaceTypeSliceInt64, buf)
		WRITE_int(len(v), buf)
		for _, _v := range v {
			WRITE_int64(_v, buf)
		}
	case []uint:
		WRITE_int8(interfaceTypeSliceUint, buf)
		WRITE_int(len(v), buf)
		for _, _v := range v {
			WRITE_uint(_v, buf)
		}
	case []uint8:
		WRITE_int8(interfaceTypeSliceUint8, buf)
		b := []byte(v)
		WRITE_int(len(b), buf)
		buf.Write(b)
	case []uint16:
		WRITE_int8(interfaceTypeSliceUint16, buf)
		WRITE_int(len(v), buf)
		for _, _v := range v {
			WRITE_uint16(_v, buf)
		}
	case []uint32:
		WRITE_int8(interfaceTypeSliceUint32, buf)
		WRITE_int(len(v), buf)
		for _, _v := range v {
			WRITE_uint32(_v, buf)
		}
	case []uint64:
		WRITE_int8(interfaceTypeSliceUint64, buf)
		WRITE_int(len(v), buf)
		for _, _v := range v {
			WRITE_uint64(_v, buf)
		}
	case []string:
		WRITE_int8(interfaceTypeSliceString, buf)
		WRITE_int(len(v), buf)
		for _, s := range v {
			b := libraries.Str2bytes(s)
			WRITE_int(len(b), buf)
			buf.Write(b)
		}
	case []interface{}:
		WRITE_int8(interfaceTypeSliceInterface, buf)
		WRITE_int(len(v), buf)
		for _, s := range v {
			WRITE_any(s, buf)
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
	case HtmlKeyValueStr:
		WRITE_int8(interfaceTypeKVs, buf)
		WRITE_HtmlKeyValueStr(v, buf)
	case HtmlKeyValueInterface:
		WRITE_int8(interfaceTypeKVi, buf)
		WRITE_HtmlKeyValueInterface(v, buf)
	case []HtmlKeyValueStr:
		WRITE_int8(interfaceTypeSliceKVs, buf)
		WRITE_int(len(v), buf)
		for _, s := range v {
			WRITE_HtmlKeyValueStr(s, buf)
		}
	default:
		r := reflect.ValueOf(i)
		if r.Kind()==reflect.Map{
			WRITE_int8(interfaceTypeM,buf)
			WRITE_map(i,buf)
			return
		}
		panic("WRITE_any未设置类型" + fmt.Sprintf("%T", v))
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
	ux := READ_uint64(buf) // ok to continue in presence of error
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x
}
func READ_int(buf *libraries.MsgBuffer) int {
	return int(READ_int64(buf))
}
func READ_uint(buf *libraries.MsgBuffer) uint {
	return uint(READ_uint64(buf))
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
	var x uint64
	var s uint
	for i := 0; i < MaxVarintLen64; i++ {
		b := buf.Next(1)[0]
		if b < 0x80 {
			if i == MaxVarintLen64-1 && b > 1 {
				panic("overflow")
			}
			return x | uint64(b)<<s
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
	panic("overflow")
	return x
}

func READ_string(buf *libraries.MsgBuffer) string {
	length := READ_int(buf)
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
	b := buf.Next(4)
	u := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
	return math.Float32frombits(u)
}
func READ_float64(buf *libraries.MsgBuffer) float64 {
	b := buf.Next(8)
	u := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
	return math.Float64frombits(u)
}

func READ_MSG_DATA(buf *libraries.MsgBuffer) (out MSG_DATA) {
	defer func() {
		if err := recover(); err != nil {
			out = &MSG_HOST_QueryErr{Err: fmt.Sprintf("%v\r\n", err) + string(debug.Stack())}
			return
		}
	}()
	cmd := READ_int32(buf)
	if f, ok := cmdMapFunc[cmd]; ok {
		return f(buf)
	}
	return nil
}

func READ_map(i interface{}, buf *libraries.MsgBuffer) {
	l := READ_int(buf)
	if l == 0 {
		return
	}
	r := reflect.ValueOf(i)
	if r.Kind() != reflect.Ptr {
		panic("WRITE_map传入" + r.Kind().String())
	}
	r = r.Elem()
	r.Set(reflect.MakeMap(r.Type()))
	for i := 0; i < l; i++ {
		r.SetMapIndex(reflect.ValueOf(read_any_result(buf)), reflect.ValueOf(read_any_result(buf,r.Type().Elem())))
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
		l := READ_int(buf)
		b := buf.Next(l)
		*((*string)(unsafe.Pointer(uint_ptr))) = string(b)
	case interfaceTypeSliceInt:
		l := READ_int(buf)
		s := make([]int, l)
		for k := range s {
			s[k] = READ_int(buf)
		}
		*((*[]int)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceInt8:
		l := READ_int(buf)
		s := make([]byte, l)
		copy(s, buf.Next(int(l)))
		*((*[]byte)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceInt16:
		l := READ_int(buf)
		s := make([]int16, l)
		for k := range s {
			s[k] = READ_int16(buf)
		}
		*((*[]int16)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceInt32:
		l := READ_int(buf)
		s := make([]int32, l)
		for k := range s {
			s[k] = READ_int32(buf)
		}
		*((*[]int32)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceInt64:
		l := READ_int(buf)
		s := make([]int64, l)
		for k := range s {
			s[k] = READ_int64(buf)
		}
		*((*[]int64)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint:
		l := READ_int(buf)
		s := make([]uint, l)
		for k := range s {
			s[k] = READ_uint(buf)
		}
		*((*[]uint)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint8:
		l := READ_int(buf)
		s := make([]uint8, l)
		for k := range s {
			s[k] = READ_uint8(buf)
		}
		*((*[]uint8)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint16:
		l := READ_int(buf)
		s := make([]uint16, l)
		for k := range s {
			s[k] = READ_uint16(buf)
		}
		*((*[]uint16)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint32:
		l := READ_int(buf)
		s := make([]uint32, l)
		for k := range s {
			s[k] = READ_uint32(buf)
		}
		*((*[]uint32)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceUint64:
		l := READ_int(buf)
		s := make([]uint64, l)
		for k := range s {
			s[k] = READ_uint64(buf)
		}
		*((*[]uint64)(unsafe.Pointer(uint_ptr))) = s
	case interfaceTypeSliceString:
		l := READ_int(buf)
		s := make([]string, l)
		for k := range s {
			_l := READ_int(buf)
			b := buf.Next(_l)
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
	case interfaceTypeKVi:
		*((*HtmlKeyValueInterface)(unsafe.Pointer(uint_ptr))) = READ_HtmlKeyValueInterface(buf)
	case interfaceTypeKVs:
		*((*HtmlKeyValueStr)(unsafe.Pointer(uint_ptr))) = READ_HtmlKeyValueStr(buf)
	default:
		panic(fmt.Sprintf("READ_any未处理类型%v", t))
	}

}
func read_any_result(buf *libraries.MsgBuffer,r_t...reflect.Type) interface{} {
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
		l := READ_int(buf)
		b := buf.Next(l)
		return string(b)
	case interfaceTypeSliceInt:
		l := READ_int(buf)
		s := make([]int, l)
		for k := range s {
			s[k] = READ_int(buf)
		}
		return s
	case interfaceTypeSliceInt8:
		l := READ_int(buf)
		s := make([]byte, l)
		copy(s, buf.Next(int(l)))
		return s
	case interfaceTypeSliceInt16:
		l := READ_int(buf)
		s := make([]int16, l)
		for k := range s {
			s[k] = READ_int16(buf)
		}
		return s
	case interfaceTypeSliceInt32:
		l := READ_int(buf)
		s := make([]int32, l)
		for k := range s {
			s[k] = READ_int32(buf)
		}
		return s
	case interfaceTypeSliceInt64:
		l := READ_int(buf)
		s := make([]int64, l)
		for k := range s {
			s[k] = READ_int64(buf)
		}
		return s
	case interfaceTypeSliceUint:
		l := READ_int(buf)
		s := make([]uint, l)
		for k := range s {
			s[k] = READ_uint(buf)
		}
		return s
	case interfaceTypeSliceUint8:
		l := READ_int(buf)
		s := make([]uint8, l)
		for k := range s {
			s[k] = READ_uint8(buf)
		}
		return s
	case interfaceTypeSliceUint16:
		l := READ_int(buf)
		s := make([]uint16, l)
		for k := range s {
			s[k] = READ_uint16(buf)
		}
		return s
	case interfaceTypeSliceUint32:
		l := READ_int(buf)
		s := make([]uint32, l)
		for k := range s {
			s[k] = READ_uint32(buf)
		}
		return s
	case interfaceTypeSliceUint64:
		l := READ_int(buf)
		s := make([]uint64, l)
		for k := range s {
			s[k] = READ_uint64(buf)
		}
		return s
	case interfaceTypeSliceString:
		l := READ_int(buf)
		s := make([]string, l)
		for k := range s {
			_l := READ_int(buf)
			b := buf.Next(_l)
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
	case interfaceTypeKVs:
		return READ_HtmlKeyValueStr(buf)
	case interfaceTypeKVi:
		return READ_HtmlKeyValueInterface(buf)
	case interfaceTypeSliceInterface:
		l := READ_int(buf)
		s := make([]interface{}, l)
		for k := range s {
			s[k] = read_any_result(buf,r_t...)
		}
		return s
	case interfaceTypeSliceKVs:
		l := READ_int(buf)
		s := make([]HtmlKeyValueStr, l)
		for k := range s {
			s[k] = READ_HtmlKeyValueStr(buf)
		}
		return s
	case interfaceTypeM:
		if len(r_t)==0{
			panic(fmt.Sprintf("read_any_result interfaceTypeM 未传入Type"))
		}
		r:=reflect.MakeMap(r_t[0])
		l := READ_int(buf)
		if l == 0 {
			return nil
		}
		for i := 0; i < l; i++ {
			r.SetMapIndex(reflect.ValueOf(read_any_result(buf)), reflect.ValueOf(read_any_result(buf,r.Type().Elem())))
		}
		return r.Interface()
	default:
		panic(fmt.Sprintf("read_any_result未处理类型%v", t))
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
		l := READ_int(buf)
		if v.Elem().Kind() == reflect.Uint8 {
			b := make([]byte, l)
			copy(b, buf.Next(l))
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
		case "HtmlKeyValueInterface":
			return reflect.ValueOf(READ_HtmlKeyValueInterface)
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
func WRITE_HtmlKeyValueInterface(kv HtmlKeyValueInterface, buf *libraries.MsgBuffer) {
	WRITE_string(kv.Key, buf)
	WRITE_any(kv.Value, buf)
}
func READ_HtmlKeyValueInterface(buf *libraries.MsgBuffer) HtmlKeyValueInterface {
	return HtmlKeyValueInterface{
		Key:   READ_string(buf),
		Value: read_any_result(buf),
	}
}

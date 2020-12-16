package mysql

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"libraries"
	"log"
	"math"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

const (
	text_pk_type_str = "varchar(255)"
	Uintptr_offset   = 32 << (^uint(0) >> 63) / 8
)

type storeEngine struct {
	name string
	aria *ariasetting
}
type ariasetting struct {
	TRANSACTIONAL  bool   //事务 默认关
	PAGE_CHECKSUM  bool   //校验 默认关
	TABLE_CHECKSUM bool   //默认 关
	ROW_FORMAT     string //页格式 默认DYNAMIC
}

//mysql结构
type Mysql_columns struct {
	Name        string
	Sql_type    string
	Null        string
	Sql_default interface{}
	Primary     bool
	Autoinc     bool
}

type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
type Transaction struct {
	conn *Mysql_Conn
}

func (t *Transaction) GetTransaction() *Mysql_Conn {
	return t.conn
}

func (t *Transaction) EndTransaction() {
	if conn := t.conn; conn != nil {
		t.conn = nil
		//rollback
		conn.Exec([]byte{114, 111, 108, 108, 98, 97, 99, 107})
		conn.db.EndTransaction(conn)
	}
}

func (t *Transaction) Commit() error {
	_, _, err := t.conn.Exec([]byte{99, 111, 109, 109, 105, 116})
	return err
}

func (t *Transaction) Rollback() error {
	_, _, err := t.conn.Exec([]byte{114, 111, 108, 108, 98, 97, 99, 107})
	return err
}

/*执行select专用
 *返回数据结构模式[]map[string]string
 */
func (db *MysqlDB) QueryString(format string, args ...interface{}) (maps []map[string]string, err error) {
	return QueryMap(Str2bytes(format), args, db, &Transaction{})
}

func QueryMap(sql []byte, prepare_arg []interface{}, db *MysqlDB, t *Transaction) (maps []map[string]string, err error) {
	row := rows_pool.Get().(*MysqlRows)
	defer rows_pool.Put(row)

	var columns []MysqlColumn
	var ts *Mysql_Conn
	if t != nil {
		ts = t.GetTransaction()
	}

Retry:
	if ts != nil {
		if prepare_arg != nil {
			stmt, err := ts.Prepare(sql)
			if err != nil {
				return nil, err
			}
			columns, err = stmt.Query(prepare_arg, row)
		} else {
			columns, err = ts.Query(sql, row)
		}

		if err != nil {
			return
		}
	} else {
		columns, err = db.Query(sql, row, prepare_arg)
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				goto Retry
			} else if strings.Contains(err.Error(), "broken pipe") { //unix断连
				goto Retry
			} else {
				return nil, err
			}
		}
	}
	if row.result_len == 0 {
		return
	}
	maps = make([]map[string]string, row.result_len)

	if !row.IsBinary {
		for index, msglen := range row.msg_len {
			row.Buffer2.Reset()
			row.Buffer2.Write(row.Buffer.Next(msglen))

			//将行数据保存到record字典
			record := make(map[string]string, len(columns))
			for _, column := range columns {
				row.buffer, err = ReadLength_Coded_Byte(row.Buffer2)
				if err != nil {
					return
				}
				record[string(column.name)] = string(row.buffer)
			}
			maps[index] = record
		}
	} else {
		nulllen := (len(columns) + 7 + 2) / 8
		for index, msglen := range row.msg_len {
			data := row.Buffer.Next(msglen)
			if data[0] != 0 {
				return nil, errors.New("返回协议错误，返回的内容不是Binary Protocol")
			}
			pos := 1 + nulllen
			nullMask := data[1 : 1+pos]
			record := make(map[string]string, len(columns))

			for i, column := range columns {
				key := string(column.name)
				if nullMask[i/8]>>(uint(i)&7) == 1 {
					record[key] = "NULL"
					continue
				}

				record[key], err = binaryToStr(columns[i], data, &pos, row)
				if err != nil {
					return
				}
			}
			maps[index] = record
		}
	}

	return maps, nil
}
func Query(sql []byte, prepare_arg []interface{}, db *MysqlDB, t *Transaction, r interface{}) (err error) {
	var is_struct, is_ptr bool
	var obj_t, type_struct reflect.Type
	var field_m map[string]*Field_struct
	var header *SliceHeader
	var ref_ptr unsafe.Pointer

	obj_t = reflect.TypeOf(r)
	if obj_t.Kind() != reflect.Ptr {
		err = errors.New("传入的不是指针无法赋值")
		return
	}
	obj_t = obj_t.Elem()

	switch obj_t.Kind() {
	case reflect.Slice:
		//取出slice里面的类型
		ref_ptr = reflect2.PtrOf(r)
		header = (*SliceHeader)(ref_ptr)
		header.Len = 0
		type_struct = obj_t.Elem()
		switch type_struct.Kind() {
		case reflect.Struct:
		case reflect.Ptr:
			type_struct = type_struct.Elem()
			if type_struct.Kind() == reflect.Struct {
				is_ptr = true
			} else {
				err = errors.New("不支持的反射类型,只能对“[]结构体”进行反射")
				return
			}
		default:
			err = errors.New("不支持的反射类型,只能对“[]结构体”进行反射")
			return
		}

	case reflect.Struct:
		type_struct = obj_t
		is_struct = true
		ref_ptr = reflect2.PtrOf(r)
	case reflect.Ptr:
		is_ptr = true
		type_struct = obj_t.Elem()
		switch type_struct.Kind() {

		case reflect.Struct:
			ref_ptr = reflect2.PtrOf(r)
			is_struct = true
		default:
			err = errors.New("不支持的反射类型")
			return
		}
	default:
		err = errors.New("只能对slice和结构体进行反射赋值")
		return
	}

	row := rows_pool.Get().(*MysqlRows)
	defer rows_pool.Put(row)

	var columns []MysqlColumn
	var ts *Mysql_Conn
	if t != nil {
		ts = t.GetTransaction()
	}

Retry:
	if ts != nil {
		if prepare_arg != nil {
			stmt, err := ts.Prepare(sql)
			if err != nil {
				return err
			}
			columns, err = stmt.Query(prepare_arg, row)
		} else {
			columns, err = ts.Query(sql, row)
		}

		if err != nil {
			return
		}
	} else {
		columns, err = db.Query(sql, row, prepare_arg)
		if err != nil {

			if strings.Contains(err.Error(), "EOF") {
				goto Retry
			} else if strings.Contains(err.Error(), "broken pipe") { //unix断连
				goto Retry
			} else {
				return err
			}
		}
	}

	if row.result_len == 0 {
		return nil
	}
	if row.field_m[type_struct.Name()] == nil {
		row.field_m[type_struct.Name()] = make(map[string]*Field_struct)
	}
	field_m = row.field_m[type_struct.Name()]

	var field_struct *Field_struct
	var uint_ptr, offset uintptr

	if is_struct {
		offset = 0
		if is_ptr {
			if reflect2.IsNil(*(*interface{})(unsafe.Pointer(ref_ptr))) {
				*(*uintptr)(ref_ptr) = reflect.New(type_struct).Pointer()
			}
		}
		row.msg_len = row.msg_len[:1]
	} else {
		if header.Len < row.result_len {
			if header.Cap < row.result_len {
				valType := reflect2.TypeOf(r)
				var elemType = valType.(*reflect2.UnsafePtrType).Elem()

				elemType.(*reflect2.UnsafeSliceType).UnsafeGrow(ref_ptr, row.result_len)
			} else {
				header.Len = row.result_len
			}
		}
		ref_ptr = header.Data
		if is_ptr {
			offset = Uintptr_offset
		} else {
			offset = type_struct.Size()
		}

	}
	if !row.IsBinary {
		for index, mglen := range row.msg_len {
			uint_ptr = uintptr(ref_ptr) + offset*uintptr(index)
			if is_ptr {
				if reflect2.IsNil(*(*interface{})(unsafe.Pointer(uint_ptr))) {
					//obj_v.Index(index).Set(reflect.New(obj_v.Type().Elem()))
					*((*uintptr)(unsafe.Pointer(uint_ptr))) = reflect.New(type_struct).Pointer()
				}
				uint_ptr = *(*uintptr)(unsafe.Pointer(uint_ptr)) //获取指针真正的地址
			}

			row.Buffer2.Reset()
			row.Buffer2.Write(row.Buffer.Next(mglen))

			for _, column := range columns {

				row.buffer, err = ReadLength_Coded_Byte(row.Buffer2)
				if err != nil {
					return err
				}

				if v, ok := field_m[*(*string)(unsafe.Pointer(&column.name))]; ok {
					if v.Kind == reflect.Invalid {
						continue
					}
					field_struct = v
				} else {
					real_key := string(column.name)

					field, ok := type_struct.FieldByName(real_key)
					if !ok {
						field_m[real_key] = &Field_struct{Kind: reflect.Invalid}
						continue
					}
					field_struct = &Field_struct{Offset: field.Offset, Kind: field.Type.Kind(), Field_t: field.Type}
					if err := checkKind(field.Type.Kind()); err != nil {
						return errors.New("不支持的类型，字段名称" + string(column.name) + "预计类型" + field_struct.Kind.String())
					}
					field_m[real_key] = field_struct

				}

				switch field_struct.Kind {
				case reflect.Int:
					ii, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&row.buffer)))
					*((*int)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = ii
				case reflect.Int8:
					ii, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&row.buffer)))
					*((*int8)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = int8(ii)
				case reflect.Int16:
					ii, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&row.buffer)))
					*((*int16)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = int16(ii)
				case reflect.Int32:
					ii, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&row.buffer)))
					*((*int32)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = int32(ii)
				case reflect.Int64:
					ii, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&row.buffer)))
					*((*int64)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = int64(ii)
				case reflect.Uint:
					ii, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&row.buffer)))
					*((*uint)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = uint(ii)
				case reflect.Uint8:
					ii, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&row.buffer)))
					*((*uint8)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = uint8(ii)
				case reflect.Uint16:
					ii, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&row.buffer)))
					*((*uint16)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = uint16(ii)
				case reflect.Uint32:
					ii, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&row.buffer)))
					*((*uint32)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = uint32(ii)
				case reflect.Uint64:
					ii, _ := strconv.ParseUint(*(*string)(unsafe.Pointer(&row.buffer)), 10, 64)
					*((*uint64)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = uint64(ii)
				case reflect.Float32:
					f, _ := strconv.ParseFloat(*(*string)(unsafe.Pointer(&row.buffer)), 32)
					*((*float32)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = float32(f)
				case reflect.Float64:
					f, _ := strconv.ParseFloat(*(*string)(unsafe.Pointer(&row.buffer)), 64)
					*((*float64)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = f
				case reflect.String:
					if str := string(row.buffer); str != "NULL" {
						*((*string)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = string(row.buffer)
					}

				case reflect.Bool:
					*((*bool)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = row.buffer[0] == 49
				case reflect.Struct:
					switch field_struct.Field_t.String() {
					case "time.Time":
						*((*time.Time)(unsafe.Pointer(uint_ptr + field_struct.Offset))), _ = time.ParseInLocation("2006-01-02 15:04:05", string(row.buffer), time.Local)
					default:
						field := reflect.NewAt(field_struct.Field_t, unsafe.Pointer(uint_ptr+field_struct.Offset))
						jsoniter.Unmarshal(row.buffer, field.Interface())

					}

				case reflect.Slice, reflect.Map:
					field := reflect.NewAt(field_struct.Field_t, unsafe.Pointer(uint_ptr+field_struct.Offset))
					jsoniter.Unmarshal(row.buffer, field.Interface())
				case reflect.Ptr:
					if *(*string)(unsafe.Pointer(&row.buffer)) != "NULL" {
						if len(row.buffer) == 0 || (len(row.buffer) == 1 && row.buffer[0] == 0xC0) {
							continue
						}
						field := reflect.New(field_struct.Field_t.Elem())
						err := jsoniter.Unmarshal(row.buffer, field.Interface())
						if err == nil {
							*((*uintptr)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = field.Pointer()
						}
					}
				}

			}

		}
	} else {
		nulllen := (len(columns) + 7 + 2) / 8
		for index, msglen := range row.msg_len {
			uint_ptr = uintptr(ref_ptr) + offset*uintptr(index)
			if is_ptr {
				if reflect2.IsNil(*(*interface{})(unsafe.Pointer(uint_ptr))) {
					//obj_v.Index(index).Set(reflect.New(obj_v.Type().Elem()))
					*((*uintptr)(unsafe.Pointer(uint_ptr))) = reflect.New(type_struct).Pointer()
				}
				uint_ptr = *(*uintptr)(unsafe.Pointer(uint_ptr)) //获取指针真正的地址
			}

			data := row.Buffer.Next(msglen)
			if data[0] != 0 {
				return errors.New("返回协议错误，返回的内容不是Binary Protocol")
			}
			pos := 1 + nulllen
			nullMask := data[1:pos]
			for i, column := range columns {

				if ((nullMask[(i+2)>>3] >> uint((i+2)&7)) & 1) == 1 {
					continue
				}
				if v, ok := field_m[*(*string)(unsafe.Pointer(&column.name))]; ok {
					if v.Kind == reflect.Invalid {
						binaryToStr(column, data, &pos, row) //跳过这段pos
						continue
					}
					field_struct = v
				} else {
					real_key := string(column.name)

					field, ok := type_struct.FieldByName(real_key)
					if !ok {
						field_m[real_key] = &Field_struct{Kind: reflect.Invalid}
						binaryToStr(column, data, &pos, row) //跳过这段pos
						continue
					}
					field_struct = &Field_struct{Offset: field.Offset, Kind: field.Type.Kind(), Field_t: field.Type}
					if err := checkKind(field.Type.Kind()); err != nil {
						return err
					}
					field_m[real_key] = field_struct
				}
				switch field_struct.Kind {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Bool:
					var u64 uint64
					switch column.fieldtype {
					case fieldTypeTiny:
						if column.fleldflag&flagUnsigned != 0 {
							u64 = uint64(data[pos])
						} else {
							u64 = uint64(int8(data[pos]))
						}
						pos++
					case fieldTypeShort, fieldTypeYear:
						if column.fleldflag&flagUnsigned != 0 {
							u64 = uint64(binary.LittleEndian.Uint16(data[pos : pos+2]))
						} else {
							u64 = uint64(int16(binary.LittleEndian.Uint16(data[pos : pos+2])))
						}
						pos += 2
					case fieldTypeInt24, fieldTypeLong:
						if column.fleldflag&flagUnsigned != 0 {
							u64 = uint64(binary.LittleEndian.Uint32(data[pos : pos+4]))
						} else {
							u64 = uint64(int32(binary.LittleEndian.Uint32(data[pos : pos+4])))
						}
						pos += 4

					case fieldTypeLongLong:
						if column.fleldflag&flagUnsigned != 0 {
							val := binary.LittleEndian.Uint64(data[pos : pos+8])
							if val > math.MaxUint64 {
								return errors.New("字段" + string(column.name) + "为整数，但是结果大于MaxUint64无法赋值")
							} else {
								u64 = uint64(val)
							}
						} else {
							u64 = uint64(binary.LittleEndian.Uint64(data[pos : pos+8]))
						}
						pos += 8

					case fieldTypeFloat, fieldTypeDouble:

						return errors.New("字段" + string(column.name) + "为整数，但是返回结果为小数浮点，无法赋值")

					// Length coded Binary Strings
					case fieldTypeDecimal, fieldTypeNewDecimal, fieldTypeVarChar,
						fieldTypeBit, fieldTypeEnum, fieldTypeSet, fieldTypeTinyBLOB,
						fieldTypeMediumBLOB, fieldTypeLongBLOB, fieldTypeBLOB,
						fieldTypeVarString, fieldTypeString, fieldTypeGeometry, fieldTypeJSON:

						msglen, err := ReadLength_Coded_Slice(data[pos:], &pos)
						if err != nil {
							return errors.New("字段" + string(column.name) + "赋值错误:" + err.Error())
						}
						u64, err = strconv.ParseUint(string(data[pos:pos+msglen]), 10, 64)
						if err != nil {
							return errors.New("字段" + string(column.name) + ",原始值:" + string(data[pos:pos+msglen]) + ",赋值错误:" + err.Error())
						}
						pos += msglen
					case
						fieldTypeDate, fieldTypeNewDate, // Date YYYY-MM-DD
						fieldTypeTime,                         // Time [-][H]HH:MM:SS[.fractal]
						fieldTypeTimestamp, fieldTypeDateTime: // Timestamp YYYY-MM-DD HH:MM:SS[.fractal]
						return errors.New("字段" + string(column.name) + "为整数，但是返回结果为日期，无法赋值")
					// Please report if this happens!
					default:
						return fmt.Errorf("unknown field type %d", columns[i].fieldtype)
					}
					switch field_struct.Kind {
					case reflect.Int:
						*((*int)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = int(u64)
					case reflect.Int8:
						*((*int8)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = int8(u64)
					case reflect.Int16:
						*((*int16)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = int16(u64)
					case reflect.Int32:
						*((*int32)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = int32(u64)
					case reflect.Int64:
						*((*int64)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = int64(u64)
					case reflect.Uint:
						*((*uint)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = uint(u64)
					case reflect.Uint8:
						*((*uint8)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = uint8(u64)
					case reflect.Uint16:
						*((*uint16)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = uint16(u64)
					case reflect.Uint32:
						*((*uint32)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = uint32(u64)
					case reflect.Uint64:
						*((*uint64)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = u64
					case reflect.Bool:
						*((*bool)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = u64 == 1
					}
				case reflect.Float32, reflect.Float64:
					var f float64
					switch column.fieldtype {
					case fieldTypeTiny:
						if column.fleldflag&flagUnsigned != 0 {
							f = float64(data[pos])
						} else {
							f = float64(int8(data[pos]))
						}
						pos++
					case fieldTypeShort, fieldTypeYear:
						if column.fleldflag&flagUnsigned != 0 {
							f = float64(binary.LittleEndian.Uint16(data[pos : pos+2]))
						} else {
							f = float64(int16(binary.LittleEndian.Uint16(data[pos : pos+2])))
						}
						pos += 2
					case fieldTypeInt24, fieldTypeLong:
						if column.fleldflag&flagUnsigned != 0 {
							f = float64(binary.LittleEndian.Uint32(data[pos : pos+4]))
						} else {
							f = float64(int32(binary.LittleEndian.Uint32(data[pos : pos+4])))
						}
						pos += 4

					case fieldTypeLongLong:
						if column.fleldflag&flagUnsigned != 0 {
							val := binary.LittleEndian.Uint64(data[pos : pos+8])
							if val > math.MaxInt64 {
								f, err = strconv.ParseFloat(string(uint64ToString(val)), 64)
								if err != nil {
									return errors.New("字段" + string(column.name) + ",原始值:" + string(uint64ToString(val)) + ",赋值错误:" + err.Error())
								}
							} else {
								f = float64(val)
							}
						} else {
							f = float64(binary.LittleEndian.Uint64(data[pos : pos+8]))
						}
						pos += 8

					case fieldTypeFloat:
						f = float64(math.Float32frombits(binary.LittleEndian.Uint32(data[pos : pos+4])))
						pos += 4
					case fieldTypeDouble:
						f = math.Float64frombits(binary.LittleEndian.Uint64(data[pos : pos+8]))
						pos += 8
					// Length coded Binary Strings
					case fieldTypeDecimal, fieldTypeNewDecimal, fieldTypeVarChar,
						fieldTypeBit, fieldTypeEnum, fieldTypeSet, fieldTypeTinyBLOB,
						fieldTypeMediumBLOB, fieldTypeLongBLOB, fieldTypeBLOB,
						fieldTypeVarString, fieldTypeString, fieldTypeGeometry, fieldTypeJSON:

						msglen, err := ReadLength_Coded_Slice(data[pos:], &pos)
						if err != nil {
							return errors.New("字段" + string(column.name) + "赋值错误:" + err.Error())
						}
						f, err = strconv.ParseFloat(string(data[pos:pos+msglen]), 64)
						if err != nil {
							return errors.New("字段" + string(column.name) + ",原始值:" + string(data[pos:pos+msglen]) + ",赋值错误:" + err.Error())
						}
						pos += msglen
					case
						fieldTypeDate, fieldTypeNewDate, // Date YYYY-MM-DD
						fieldTypeTime,                         // Time [-][H]HH:MM:SS[.fractal]
						fieldTypeTimestamp, fieldTypeDateTime: // Timestamp YYYY-MM-DD HH:MM:SS[.fractal]
						return errors.New("字段" + string(column.name) + "为浮点，但是返回结果为日期，无法赋值")
					// Please report if this happens!
					default:
						return fmt.Errorf("unknown field type %d", columns[i].fieldtype)
					}
					if field_struct.Kind == reflect.Float32 {
						*((*float32)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = float32(f)
					} else {
						*((*float64)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = f
					}
				case reflect.String, reflect.Struct, reflect.Slice, reflect.Map:
					str, err := binaryToStr(column, data, &pos, row)
					if err != nil {
						return errors.New("字段" + string(column.name) + "读取错误1" + err.Error())
					}
					if field_struct.Kind == reflect.String {
						*((*string)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = str
						continue
					}
					if str == "" || str == "NULL" {
						continue
					}
					switch {
					case field_struct.Kind == reflect.Struct && field_struct.Field_t.String() == "time.Time":
						*((*time.Time)(unsafe.Pointer(uint_ptr + field_struct.Offset))), _ = time.ParseInLocation("2006-01-02 15:04:05", str, row.conn.loc)
					case field_struct.Kind == reflect.Ptr:
						if *(*string)(unsafe.Pointer(&row.buffer)) != "NULL" {
							if len(row.buffer) == 0 || (len(row.buffer) == 1 && row.buffer[0] == 0xC0) {
								continue
							}
							field := reflect.New(field_struct.Field_t.Elem())
							err = jsoniter.Unmarshal(row.buffer, field.Interface())
							if err == nil {
								*((*uintptr)(unsafe.Pointer(uint_ptr + field_struct.Offset))) = field.Pointer()
							}
						}
					default:
						field := reflect.NewAt(field_struct.Field_t, unsafe.Pointer(uint_ptr+field_struct.Offset))
						err = jsoniter.Unmarshal([]byte(str), field.Interface())
					}
					if err != nil {
						return errors.New("字段" + string(column.name) + ",原始值:" + str + "   json解析错误:" + err.Error())
					}
				}

			}
		}
	}
	return
}

/*执行sql语句
 *返回新增ID和error
 *
 */
func Insert(insert_sql []byte, prepare_arg []interface{}, db *MysqlDB, t *Transaction) (LastInsertId, rowsAffected int64, err error) {

	var ts *Mysql_Conn
	if t != nil {
		ts = t.GetTransaction()
	}
	if ts != nil {
		if prepare_arg != nil {

			stmt, err := ts.Prepare(insert_sql)
			if err != nil {
				return 0, 0, err
			}
			err = stmt.Exec(prepare_arg)
			return stmt.lastInsertId, stmt.rowsAffected, err
		} else {
			LastInsertId, rowsAffected, err = ts.Exec(insert_sql)

		}

	} else {
	Retry:
		LastInsertId, rowsAffected, err = db.Exec(insert_sql, prepare_arg)

		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				goto Retry
			} else if strings.Contains(err.Error(), "broken pipe") { //unix断连
				goto Retry
			} else {
				return 0, 0, err
			}
		}
	}
	return
}

/*执行sql语句
 *返回error
 *
 */
func Exec(query_sql []byte, prepare_arg []interface{}, db *MysqlDB, t *Transaction) (err error) {
	var ts *Mysql_Conn
	if t != nil {
		ts = t.GetTransaction()
	}
	if ts != nil {
		if prepare_arg != nil {
			stmt, err := ts.Prepare(query_sql)
			if err != nil {
				return err
			}
			err = stmt.Exec(prepare_arg)
			return err
		} else {
			_, _, err = ts.Exec(query_sql)
		}
	} else {
	Retry:
		_, _, err = db.Exec(query_sql, prepare_arg)
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				goto Retry
			} else if strings.Contains(err.Error(), "broken pipe") { //unix断连
				goto Retry
			} else {
				return err
			}
		}

	}
	return
}

//执行语句并取受影响行数
func Query_getaffected(query_sql []byte, prepare_arg []interface{}, db *MysqlDB, t *Transaction) (rowsAffected int64, err error) {
	var ts *Mysql_Conn
	if t != nil {
		ts = t.GetTransaction()
	}
	if ts != nil {
		if prepare_arg != nil {
			stmt, err := ts.Prepare(query_sql)
			if err != nil {
				return 0, err
			}
			err = stmt.Exec(prepare_arg)
			return stmt.rowsAffected, err
		} else {
			_, rowsAffected, err = ts.Exec(query_sql)

		}
	} else {
	Retry:
		_, rowsAffected, err = db.Exec(query_sql, prepare_arg)

		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				goto Retry
			} else if strings.Contains(err.Error(), "broken pipe") { //unix断连
				goto Retry
			} else {
				return 0, err
			}
		}
	}
	return
}

/*列出所有表
func (this *Mysql) ShowTables(master string) (list orm.ParamsList) {
	if master != "slave" && master != "default" {
		master = "default"
	}
	s := o
	s.Using(master)
	sql := "SHOW TABLES"
	s.Raw(sql).ValuesFlat(&list)
	return
}*/

/*列出表结构
func (this *Mysql) ShowColumns(table string, master string) map[string]Mysql_columns {
	sql := "SHOW COLUMNS FROM `" + table + "`"
	result, err := this.Select(sql, master, new(Transaction))
	Errorlog(err, "初始化错误，无法列出表结构")
	re := make(map[string]Mysql_columns)
	for _, tmp := range result {
		re[tmp["Field"].(string)] = Mysql_columns{Name: tmp["Field"].(string), Sql_type: tmp["Type"].(string), Null: tmp["Null"].(string), Sql_default: tmp["Default"], Primary: (tmp["Key"].(string) == "PRI"), Autoinc: (tmp["Extra"].(string) == "auto_increment")}
	}
	return re
}*/
var mysqlLoc = time.UTC

func Open(dsn string) (*MysqlDB, error) {
	db := &MysqlDB{}
	var str [][]string
	if str, _ = Preg_match_result(`([^:]+):([^@]*)@(tcp)?(unix)?\(([^)]*)\)\/([^?]+)(\?[^?]+)`, dsn, 1); len(str) == 0 {
		log.Fatal("mysql初始化失败，解析连接字串错误" + dsn)

	}

	libraries.DebugLog(time.Now().Format("2006-01-02 15:04:05") + "连接据库" + dsn)
	var charset = "utf8"

	if str[0][7] != "" {
		for _, s := range strings.Split(str[0][7], "&") {
			if value := strings.Split(url.PathEscape(s), "="); len(value) == 2 {
				switch value[0] {
				case "charset":
					charset = value[1]
				case "loc":
					if newloc, err := time.LoadLocation(value[1]); err == nil {
						mysqlLoc = newloc
					}
				}
			}
		}
	}
	db = mysql_open(str[0][1], str[0][2], str[0][5], str[0][6], charset, mysqlLoc, nil)

	err := db.Ping()
	if err != nil {
		return nil, errors.New("数据库连接失败" + err.Error())

	}
	return db, nil

}
func (db *MysqlDB) StoreEngine(storeEngine string) *MysqlDB {
	db.storeEngine.name = storeEngine
	return db
}
func (db *MysqlDB) Sync2(i ...interface{}) (errs []error) {

	var default_engine string
	var support_tokudb bool
	var support_Archive bool
	var support_Aria bool
	res, err := db.QueryString("show engines")
	if err != nil {
		return []error{err}
	}
	for _, v := range res {
		if v["Support"] == "DEFAULT" {
			default_engine = v["Engine"]
		}
		switch v["Engine"] {
		case "TokuDB":
			support_tokudb = (v["Support"] == "DEFAULT" || v["Support"] == "YES")
		case "Archive":
			support_Archive = (v["Support"] == "DEFAULT" || v["Support"] == "YES")
		case "Aria":
			support_Aria = (v["Support"] == "DEFAULT" || v["Support"] == "YES")
		}

	}
	if db.storeEngine.name == "" {
		db.storeEngine.name = default_engine
	}
	switch db.storeEngine.name {
	case "Archive":
		switch {
		case support_Archive:
			db.storeEngine.name = "Archive"
		case support_tokudb:
			db.storeEngine.name = "TokuDB"
		case support_Aria:
			db.storeEngine.name = "Aria"
		default:
			db.storeEngine.name = "MyISAM"
		}
	case "TokuDB":
		if !support_tokudb {
			switch {
			case support_Aria:
				db.storeEngine.name = "Aria"
			default:
				db.storeEngine.name = "MyISAM"
			}
		}
	case "Aria":
		if !support_Aria {
			db.storeEngine.name = "MyISAM"
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(i))
	for _, v := range i {
		go func(v interface{}) {
			defer wg.Done()
			buf := bytes.NewBuffer(nil)
			buf2 := bytes.NewBuffer(nil)
			obj := reflect.ValueOf(v)
			if obj.Kind() != reflect.Ptr {
				errs = append(errs, errors.New("sync2需要传入指针型struct"))
				return
			}
			r := obj.Elem()
			t := r.Type()
			table_name := t.Name()
			if f := obj.MethodByName("TableName"); f.Kind() == reflect.Func {
				rs := f.Call(nil)
				if len(rs) == 1 && rs[0].Kind() == reflect.String {
					table_name = rs[0].String()
				}
			}
			res, err := db.QueryString(`show tables like '`+table_name+`'`, nil)
			if err != nil {
				errs = append(errs, errors.New(table_name+":"+err.Error()))
				return
			}

			index := map[string]bool{} //普通索引
			if res == nil {
				buf.Reset()
				buf.WriteString("CREATE TABLE `")
				buf.WriteString(table_name)
				buf.WriteString("` (")
				buf2.Reset()
				buf2.WriteString("PRIMARY KEY (")
				var have_pk bool

				for i := 0; i < r.NumField(); i++ {
					var is_pk, notnull bool
					var default_str, extra_str string
					field := r.Field(i)
					field_t := t.Field(i)
					field_str := field_t.Name
					tag := field_t.Tag.Get(`db`)
					if tag == `-` {
						continue
					}
					if strings.Contains(tag, "pk") {
						is_pk = true
						have_pk = true
						buf2.WriteString("`" + field_str + "`")
						buf2.WriteByte(44)
						notnull = true
					}
					if strings.Contains(tag, "notnull") || strings.Contains(tag, "not null") {
						notnull = true
					}
					if strings.Contains(tag, `index`) {
						index[field_str] = true
						notnull = true
					}

					if sc, _ := Preg_match_result(`default\((\d+)\)`, tag, 1); len(sc) > 0 {
						default_str = " DEFAULT '" + sc[0][1] + "'"
					} else if sc, _ := Preg_match_result(`default\('([^']*)'\)`, tag, 1); len(sc) > 0 {
						default_str = " DEFAULT '" + sc[0][1] + "'"
					} else if sc, _ := Preg_match_result(`default\(current_timestamp\(\)\)`, tag, 1); len(sc) > 0 {
						default_str = " DEFAULT current_timestamp()"
					}
					if sc, _ := Preg_match_result(`extra\('([^']*)'\)`, tag, 1); len(sc) > 0 {
						extra_str = sc[0][1]
					}
					buf.WriteString("`" + field_str + "` ")
					var is_text bool
					switch field.Kind() {
					case reflect.Int64, reflect.Int:
						buf.WriteString("bigint(20)")
						if default_str == "" {
							default_str = " DEFAULT '0'"
						}
					case reflect.Uint64, reflect.Uint:
						buf.WriteString("bigint(20) unsigned")
						if default_str == "" {
							default_str = " DEFAULT '0'"
						}
					case reflect.String:
						if sc, _ := Preg_match_result(`type:(varchar\(\d+\))`, tag, 1); len(sc) > 0 {
							buf.WriteString(sc[0][1])
							if default_str == "" {
								default_str = " DEFAULT ''"
							}
							break
						}
						if sc, _ := Preg_match_result(`type:(char\(\d+\))`, tag, 1); len(sc) > 0 {
							buf.WriteString(sc[0][1])
							if default_str == "" {
								default_str = " DEFAULT ''"
							}
							break
						}
						if is_pk {
							buf.WriteString(text_pk_type_str)
							if default_str == "" {
								default_str = " DEFAULT ''"
							}
							break
						}
						is_text = true
						switch {
						case strings.Contains(tag, `type:mediumtext`):
							buf.WriteString("mediumtext")
						case strings.Contains(tag, `type:longtext`):
							buf.WriteString("longtext")
						case strings.Contains(tag, `type:tinytext`):
							buf.WriteString("tinytext")
						default:
							buf.WriteString("text")
						}
					case reflect.Int32:
						buf.WriteString("int(11)")
						if default_str == "" {
							default_str = " DEFAULT '0'"
						}
					case reflect.Uint32:
						buf.WriteString("int(11) unsigned")
						if default_str == "" {
							default_str = " DEFAULT '0'"
						}
					case reflect.Int8:
						buf.WriteString("tinyint(3)")
						if default_str == "" {
							default_str = " DEFAULT '0'"
						}
					case reflect.Uint8:
						buf.WriteString("tinyint(3) unsigned")
						if default_str == "" {
							default_str = " DEFAULT '0'"
						}
					case reflect.Int16:
						buf.WriteString("smallint(6)")
						if default_str == "" {
							default_str = " DEFAULT '0'"
						}
					case reflect.Uint16:
						buf.WriteString("smallint(6) unsigned")
						if default_str == "" {
							default_str = " DEFAULT '0'"
						}
					case reflect.Float32:
						buf.WriteString("float")
						if default_str == "" {
							default_str = " DEFAULT 0"
						}
					case reflect.Bool:
						buf.WriteString("tinyint(1)")
						if default_str == "" {
							default_str = " DEFAULT '0'"
						}
					case reflect.Struct:
						switch field.Interface().(type) {
						case time.Time:
							switch {
							case strings.Contains(tag, "type:datetime"):
								buf.WriteString("datetime")
							case strings.Contains(tag, "type:timestamp"):
								buf.WriteString("timestamp")
							case strings.Contains(tag, "type:time"):
								buf.WriteString("time")
							case strings.Contains(tag, "type:date"):
								buf.WriteString("date")
							default:
								buf.WriteString("datetime")
							}

							if default_str == "" {
								default_str = " DEFAULT current_timestamp()"
							}
						default:

							is_text = true
							switch {
							case strings.Contains(tag, "type:longblob"):
								buf.WriteString("longblob")
							case strings.Contains(tag, "type:mediumblob"):
								buf.WriteString("mediumblob")
							case strings.Contains(tag, "type:tinyblob"):
								buf.WriteString("tinyblob")
							case strings.Contains(tag, "type:blob"):
								buf.WriteString("blob")
							case strings.Contains(tag, "type:longtext"):
								buf.WriteString("longtext")
							case strings.Contains(tag, "type:mediumtext"):
								buf.WriteString("mediumblob")
							case strings.Contains(tag, "type:tinytext"):
								buf.WriteString("tinytext")
							default:
								buf.WriteString("text")
							}

						}
					default:
						is_text = true
						switch {
						case strings.Contains(tag, "type:longblob"):
							buf.WriteString("longblob")
						case strings.Contains(tag, "type:mediumblob"):
							buf.WriteString("mediumblob")
						case strings.Contains(tag, "type:tinyblob"):
							buf.WriteString("tinyblob")
						case strings.Contains(tag, "type:blob"):
							buf.WriteString("blob")
						case strings.Contains(tag, "type:longtext"):
							buf.WriteString("longtext")
						case strings.Contains(tag, "type:mediumtext"):
							buf.WriteString("mediumblob")
						case strings.Contains(tag, "type:tinytext"):
							buf.WriteString("tinytext")
						default:
							buf.WriteString("text")
						}

					}
					if is_pk {
						buf.WriteString(" NOT NULL")
						if strings.Contains(tag, "auto_increment") {
							buf.WriteString(" AUTO_INCREMENT")
						} else {
							buf.WriteString(default_str)
						}

						buf.WriteByte(44)
						continue
					}

					if notnull {
						buf.WriteString(" NOT NULL")
					} else {
						buf.WriteString(" NULL")
					}
					if strings.Contains(tag, "auto_increment") {
						buf.WriteString(" AUTO_INCREMENT")
					} else if !is_text {
						buf.WriteString(default_str)
					}
					buf.WriteString(" ")
					buf.WriteString(extra_str)
					buf.WriteByte(44)
				}
				if have_pk {
					buf.Write(buf2.Next(buf2.Len() - 1))
					buf.WriteString(")")
				} else {
					l := buf.Len()
					buf.Write(buf.Next(l)[:l-1])
				}
				buf.WriteString(") ENGINE=")
				buf.WriteString(db.storeEngine.name)
				buf.WriteString(" DEFAULT CHARSET=utf8")
				if db.storeEngine.name == "Aria" {
					if db.storeEngine.aria != nil {
						if db.storeEngine.aria.TRANSACTIONAL {
							buf.WriteString(" TRANSACTIONAL = 1")
						} else {
							buf.WriteString(" TRANSACTIONAL = 0")
						}
						if db.storeEngine.aria.PAGE_CHECKSUM {
							buf.WriteString(" PAGE_CHECKSUM = 1")
						} else {
							buf.WriteString(" PAGE_CHECKSUM = 0")
						}
						if db.storeEngine.aria.TABLE_CHECKSUM {
							buf.WriteString(" TABLE_CHECKSUM = 1")
						} else {
							buf.WriteString(" TABLE_CHECKSUM = 0")
						}
						if db.storeEngine.aria.ROW_FORMAT != "" {
							buf.WriteString(" ROW_FORMAT = ")
							buf.WriteString(db.storeEngine.aria.ROW_FORMAT)
						}

					} else {
						buf.WriteString(" TRANSACTIONAL = 0 PAGE_CHECKSUM = 0 TABLE_CHECKSUM = 0 ROW_FORMAT = DYNAMIC")
					}
				}
				err := Exec(buf.Bytes(), nil, db, &Transaction{})
				if err != nil {
					errs = append(errs, errors.New("执行新建数据库失败："+err.Error()+" 错误sql:"+buf.String()))
					return
				}
			} else {
				res, err = db.QueryString(`desc ` + table_name)
				if err != nil {
					errs = append(errs, errors.New(table_name+":"+err.Error()))
					return
				}
				var pk, sql []string
				var pk_num int
				var res_m = make(map[string]map[string]string, len(res))
				for _, value := range res {
					if value["Key"] == "PRI" {
						pk_num++
					}
					res_m[value["Field"]] = value
				}

				for i := 0; i < r.NumField(); i++ {
					field_t := t.Field(i)
					field := r.Field(i)
					tag := field_t.Tag.Get(`db`)
					if tag == `-` {
						continue
					}
					field_str := field_t.Name
					var is_change int8
					var is_text bool
					var notnull, is_pk bool
					var default_str, varchar_str, extra_str string
					sql_str := make([]string, 5)
					if value, ok := res_m[field_str]; ok {
						extra_str = ""
						sql_str[4] = value["Extra"]
						default_str = value["Default"]
						sql_str[1] = value["Type"]
						if value["Null"] == "YES" {
							sql_str[2] = "NULL"
						} else {
							sql_str[2] = "NOT NULL"
						}

						sql_str[3] = value["Default"]
						if sql_str[3] == "''" {
							sql_str[3] = ""
						}
						if default_str == "''" {
							default_str = ""
						}
						if strings.Contains(tag, "pk") {
							is_pk = true
							notnull = true
						}
						if strings.Contains(tag, "notnull") || strings.Contains(tag, "not null") {
							notnull = true
						}
						if strings.Contains(tag, "index") {
							index[field_str] = true
							notnull = true
							if sql_str[2] == "NULL" {
								sql_str[2] = "NOT NULL"
							}
						}

						if sc, _ := Preg_match_result(`default\((\d+)\)`, tag, 1); len(sc) > 0 {
							default_str = sc[0][1]

						} else if sc, _ := Preg_match_result(`default\('([^']*)'\)`, tag, 1); len(sc) > 0 {
							default_str = sc[0][1]
						}
						if sc, _ := Preg_match_result(`extra\('([^']*)'\)`, tag, 1); len(sc) > 0 {
							extra_str = sc[0][1]
						}
						if field.Kind() == reflect.String {
							switch {
							case strings.Contains(tag, "type:longtext"):
								varchar_str = "longtext"
								notnull = false
							case strings.Contains(tag, "type:mediumtext"):
								varchar_str = "mediumtext"
								notnull = false
							case strings.Contains(tag, "type:tinytext"):
								varchar_str = "tinytext"
								notnull = false
							case strings.Contains(tag, "type:text"):
								varchar_str = "text"
								notnull = false
							default:
								if sc, _ := Preg_match_result(`type:(varchar\(\d+\))`, tag, 1); len(sc) > 0 {
									varchar_str = sc[0][1]
								} else {
									if sc, _ := Preg_match_result(`type:(char\(\d+\))`, tag, 1); len(sc) > 0 {
										varchar_str = sc[0][1]
									}
								}
							}
						} else {
							switch {
							case strings.Contains(tag, "type:longblob"):
								varchar_str = "longblob"
								notnull = false
							case strings.Contains(tag, "type:mediumblob"):
								varchar_str = "mediumblob"
								notnull = false
							case strings.Contains(tag, "type:tinyblob"):
								varchar_str = "tinyblob"
								notnull = false
							case strings.Contains(tag, "type:blob"):
								varchar_str = "blob"
								notnull = false
							case strings.Contains(tag, "type:longtext"):
								varchar_str = "longtext"
								notnull = false
							case strings.Contains(tag, "type:mediumtext"):
								varchar_str = "mediumblob"
								notnull = false
							case strings.Contains(tag, "type:tinytext"):
								varchar_str = "tinytext"
								notnull = false
							case strings.Contains(tag, "type:text"):
								varchar_str = "text"
								notnull = false
							default:
								if sc, _ := Preg_match_result(`type:(varchar\(\d+\))`, tag, 1); len(sc) > 0 {
									varchar_str = sc[0][1]
								} else {
									if sc, _ := Preg_match_result(`type:(char\(\d+\))`, tag, 1); len(sc) > 0 {
										varchar_str = sc[0][1]
									}
								}
							}
						}

						if notnull {
							if value["Null"] == "YES" {
								is_change = 1
								sql_str[2] = "NOT NULL"
							}
						} else {
							if value["Null"] == "NO" {
								is_change = 2
								sql_str[2] = "NULL"
							}
						}
						if strings.Contains(tag, "auto_increment") {
							extra_str = "auto_increment"
							if !strings.Contains(value["Extra"], "auto_increment") {
								is_change = 3
							}
						}

						switch field.Kind() {
						case reflect.Int64, reflect.Int:
							if sql_str[1] != "bigint(20)" {
								is_change = 4
								sql_str[1] = "bigint(20)"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "0"
							}
						case reflect.Uint64, reflect.Uint:
							if sql_str[1] != "bigint(20) unsigned" {
								is_change = 5
								sql_str[1] = "bigint(20) unsigned"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "0"
							}
						case reflect.Float32:
							if sql_str[1] != "float" {
								is_change = 6
								sql_str[1] = "float"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "0"
							}
						case reflect.String:
							if varchar_str != "" {
								if sql_str[1] != varchar_str {
									is_change = 7
									sql_str[1] = varchar_str
								}
								break
							}
							sql_str[3] = default_str
							if strings.Contains(tag, "type:text") {
								is_text = true
								if is_pk {
									if sql_str[1] != text_pk_type_str {
										is_change = 8
										sql_str[1] = "text"
									}
								} else {
									if sql_str[1] != "text" {
										is_change = 9
										sql_str[1] = "text"
									}
								}
							} else {
								is_text = true
								if is_pk {
									if sql_str[1] != text_pk_type_str {
										is_change = 10
										sql_str[1] = "text"
									}
								} else {
									if sql_str[1] != "text" {
										is_change = 11
										sql_str[1] = "text"
									}
								}

							}
						case reflect.Int32:
							if sql_str[1] != "int(11)" {
								is_change = 12
								sql_str[1] = "int(11)"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "0"
							}

						case reflect.Uint32:
							if sql_str[1] != "int(11) unsigned" {
								is_change = 13
								sql_str[1] = "int(11) unsigned"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "0"
							}
						case reflect.Int8:
							if sql_str[1] != "tinyint(3)" {
								is_change = 14
								sql_str[1] = "tinyint(3)"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "0"
							}
						case reflect.Uint8:
							if sql_str[1] != "tinyint(3) unsigned" {
								is_change = 15
								sql_str[1] = "tinyint(3) unsigned"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "0"
							}
						case reflect.Int16:
							if sql_str[1] != "smallint(6)" {
								is_change = 16
								sql_str[1] = "smallint(6)"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "0"
							}
						case reflect.Uint16:
							if sql_str[1] != "smallint(6) unsigned" {
								is_change = 17
								sql_str[1] = "smallint(6) unsigned"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "0"
							}
						case reflect.Bool:
							if sql_str[1] != "tinyint(1)" {
								is_change = 18
								sql_str[1] = "tinyint(1)"
							}
							if default_str == "" || (is_pk && default_str == "NULL") {
								default_str = "1"
							}
						case reflect.Struct:
							switch field.Interface().(type) {
							case time.Time:
								var timestr = "datetime"
								switch {
								case strings.Contains(tag, "type:timestamp"):
									timestr = "timestamp"
								case strings.Contains(tag, "type:time"):
									timestr = "time"
								case strings.Contains(tag, "type:date"):
									timestr = "date"
								}
								if sql_str[1] != timestr {
									is_change = 19
									sql_str[1] = timestr
								}
								if Preg_match(`^\d{4}-\d{2}-\d{2}$`, default_str) {
									default_str += " 00:00:00"
								}
								if default_str == "" || default_str == "NULL" {
									default_str = "current_timestamp()"
								}

							default:
								is_text = true
								if !strings.Contains(sql_str[1], "text") {
									is_change = 20
									sql_str[1] = "text"
								}
								default_str = "NULL"
							}
						default:

							if varchar_str != "" {
								if sql_str[1] != varchar_str {
									is_change = 21
									sql_str[1] = varchar_str
								}
								default_str = "NULL"
							} else {
								is_text = true
								if !strings.Contains(sql_str[1], "text") {
									is_change = 22
									sql_str[1] = "text"
								}
								default_str = "NULL"
							}

						}
						if is_pk {
							pk = append(pk, field_str)
							sql_str[2] = "NOT NULL"
							if !strings.Contains(sql_str[1], "char") {
								if sql_str[3] != "0" && sql_str[3] != "current_timestamp()" {
									sql_str[3] = "NULL"
								}
								if extra_str == "auto_increment" || default_str == "" {
									default_str = "NULL"
								}
							}

							if is_text {
								sql_str[1] = text_pk_type_str
							}

						}
						if sql_str[3] != default_str {
							is_change = 23

							sql_str[3] = default_str
						}
						if sql_str[4] != extra_str {
							is_change = 24
							sql_str[4] = extra_str
						}
						if sql_str[3] != "" {
							switch sql_str[3] {
							case "current_timestamp()":
								sql_str[3] = "Default " + sql_str[3]
							case "AUTO_INCREMENT":
							case "NULL":
								sql_str[3] = "Default NULL"
							default:
								sql_str[3] = "Default '" + strings.Trim(sql_str[3], "'") + "'"
							}

						} else {
							sql_str[3] = "Default ''"
						}

						if is_change > 0 {
							if is_text {
								sql_str[3] = ""
							}
							sql_str[0] = "modify column `" + field_str + "`"
							DEBUG(is_change, sql_str)
							sql = append(sql, strings.Join(sql_str, " "))
						}
					} else {
						var after string
						if i == 0 {
							after = " FIRST"
						}
						for index := i - 1; index > -1; index-- {
							before_field := t.Field(index)
							if before_field.Tag.Get(`db`) == `-` {
								continue
							}
							after = " AFTER `" + before_field.Name + "`"
							break
						}

						switch field.Kind() {
						case reflect.Int64, reflect.Int:
							sql_str[1] = "bigint(20)"
							sql_str[3] = "Default '0'"
						case reflect.Uint64, reflect.Uint:
							sql_str[1] = "bigint(20) unsigned"
							sql_str[3] = "Default '0'"
						case reflect.String:

							sql_str[3] = "Default ''"
							if varchar_str != "" {
								sql_str[1] = varchar_str
								sql_str[3] = "Default ''"
								break
							}
							is_text = true
							sql_str[1] = "text"
						case reflect.Int32:
							sql_str[1] = "int(11)"
							sql_str[3] = "Default '0'"
						case reflect.Uint32:
							sql_str[1] = "int(11) unsigned"
							sql_str[3] = "Default '0'"
						case reflect.Int8:
							sql_str[1] = "tinyint(3)"
							sql_str[3] = "Default '0'"
						case reflect.Uint8:
							sql_str[1] = "tinyint(3) unsigned"
							sql_str[3] = "Default '0'"
						case reflect.Int16:
							sql_str[1] = "smallint(6)"
							sql_str[3] = "Default '0'"
						case reflect.Uint16:
							sql_str[1] = "smallint(6) unsigned"
							sql_str[3] = "Default '0'"

						case reflect.Bool:
							sql_str[1] = "tinyint(1)"
							sql_str[3] = "Default '0'"
						case reflect.Struct:
							switch r.Field(i).Interface().(type) {
							case time.Time:
								switch {
								case strings.Contains(tag, "type:timestamp"):
									sql_str[1] = "timestamp"
								case strings.Contains(tag, "type:time"):
									sql_str[1] = "time"
								case strings.Contains(tag, "type:date"):
									sql_str[1] = "date"
								default:
									sql_str[1] = "datetime"
								}

								sql_str[3] = "Default current_timestamp()"
							default:
								is_text = true
								sql_str[1] = "text"
							}
						default:
							is_text = true
							sql_str[1] = "text"
						}
						if strings.Contains(tag, "auto_increment") {
							if !strings.Contains(value["Extra"], "auto_increment") {
								sql_str[3] = " AUTO_INCREMENT"
							}
						}
						switch {
						case strings.Contains(tag, "type:longblob"):
							sql_str[1] = "longblob"
						case strings.Contains(tag, "type:mediumblob"):
							sql_str[1] = "mediumblob"
						case strings.Contains(tag, "type:tinyblob"):
							sql_str[1] = "tinyblob"
						case strings.Contains(tag, "type:blob"):
							sql_str[1] = "blob"
						case strings.Contains(tag, "type:longtext"):
							sql_str[1] = "longtext"
						case strings.Contains(tag, "type:mediumtext"):
							sql_str[1] = "mediumtext"
						case strings.Contains(tag, "type:tinytext"):
							sql_str[1] = "tinytext"
						case strings.Contains(tag, "type:text"):
							sql_str[1] = "text"

							sql_str[3] = strings.Replace(sql_str[3], " Default NULL", "", 1)
						default:
							if sc, _ := Preg_match_result(`type:(varchar\(\d+\))`, tag, 1); len(sc) > 0 {
								sql_str[1] = sc[0][1]
							} else {
								if sc, _ := Preg_match_result(`type:(char\(\d+\))`, tag, 1); len(sc) > 0 {
									sql_str[1] = sc[0][1]
								}
							}
						}
						if strings.Contains(tag, "notnull") || strings.Contains(tag, "not null") {
							notnull = true
						}
						if strings.Contains(tag, "pk") {
							pk = append(pk, field_str)
							sql_str[2] = "NOT NULL"
							if sql_str[3] != " AUTO_INCREMENT" {
								sql_str[3] = ""
							}
							if sql_str[1] == "text" {
								sql_str[1] = text_pk_type_str
							}
						} else {

							if sc, _ := Preg_match_result(`default\((\d+)\)`, tag, 1); len(sc) > 0 && !is_text {
								sql_str[3] = "Default '" + sc[0][1] + "'"

							}

							if sc, _ := Preg_match_result(`default\('([^']*)'\)`, tag, 1); !is_text && len(sc) > 0 {
								switch sc[0][1] {
								case "current_timestamp()":
									sql_str[3] = "Default " + sc[0][1]
								case "NULL":
									sql_str[3] = "Default NULL"
								default:
									sql_str[3] = "Default '" + strings.Trim(sc[0][1], "'") + "'"
								}

							}
							if sc, _ := Preg_match_result(`extra\('([^']*)'\)`, tag, 1); len(sc) > 0 {
								sql_str[4] = sc[0][1]
							}
							if notnull {
								sql_str[2] = "NOT NULL"
							} else {
								sql_str[2] = "NULL"
							}
						}

						sql_str[0] = "ADD `" + field_str + "`"
						sql = append(sql, strings.Join(sql_str, " ")+after)
					}
				}
				if pk_num != len(pk) {
					if pk_num == 0 {
						sql = append(sql, "ADD PRIMARY KEY (`"+strings.Join(pk, "`,`")+"`)")
					} else if len(pk) == 0 {
						sql = append(sql, "DROP PRIMARY KEY")
					} else {
						sql = append(sql, "DROP PRIMARY KEY,ADD PRIMARY KEY (`"+strings.Join(pk, "`,`")+"`)")
					}
				}
				if len(sql) > 0 {
					s := "ALTER TABLE " + table_name + " " + strings.Join(sql, ",")
					DEBUG(s)
					err := Exec(Str2bytes(s), nil, db, &Transaction{})
					if err != nil {
						errs = append(errs, errors.New(table_name+":"+err.Error()))
						return
					}
				}

				res, err := db.QueryString("select ENGINE,CREATE_OPTIONS from information_schema.tables where table_name=? and TABLE_SCHEMA=?", table_name, db.database)
				if err != nil {
					errs = append(errs, errors.New(table_name+":"+err.Error()))
					return
				}
				if res[0]["ENGINE"] != db.storeEngine.name {

					res[0]["CREATE_OPTIONS"] = ""
					err = Exec([]byte("ALTER TABLE "+table_name+" ENGINE = "+db.storeEngine.name+" transactional=default,row_format=default,checksum=0"), nil, db, &Transaction{})
					if err != nil {
						errs = append(errs, errors.New(table_name+":"+err.Error()))
						return
					}
				}
				sql = nil
				switch db.storeEngine.name {
				case "Aria":
					if strings.Contains(res[0]["CREATE_OPTIONS"], "page_checksum=1") {
						if db.storeEngine.aria == nil || (db.storeEngine.aria != nil && !db.storeEngine.aria.PAGE_CHECKSUM) {
							sql = append(sql, "page_checksum=0")
						}
						res[0]["CREATE_OPTIONS"] = strings.Replace(res[0]["CREATE_OPTIONS"], "page_checksum=1", "", 1)
					} else if db.storeEngine.aria != nil && db.storeEngine.aria.PAGE_CHECKSUM {
						sql = append(sql, "page_checksum=1")
						res[0]["CREATE_OPTIONS"] = strings.Replace(res[0]["CREATE_OPTIONS"], "page_checksum=0", "", 1)
					}
					if strings.Contains(res[0]["CREATE_OPTIONS"], "checksum=1") {
						if db.storeEngine.aria == nil || (db.storeEngine.aria != nil && !db.storeEngine.aria.TABLE_CHECKSUM) {
							sql = []string{"checksum=0"}
						}
					} else if db.storeEngine.aria != nil && db.storeEngine.aria.TABLE_CHECKSUM {
						sql = []string{"checksum=1"}
					}
					if strings.Contains(res[0]["CREATE_OPTIONS"], "transactional=1") {
						if db.storeEngine.aria == nil || (db.storeEngine.aria != nil && !db.storeEngine.aria.TRANSACTIONAL) {
							sql = append(sql, "transactional=0")
						}
					} else if db.storeEngine.aria != nil && db.storeEngine.aria.TRANSACTIONAL {
						sql = append(sql, "transactional=1")
					}
				}
				if sql != nil {
					sql_str := "ALTER TABLE " + table_name + " " + strings.Join(sql, ",")
					DEBUG(sql_str)
					err := Exec([]byte(sql_str), nil, db, &Transaction{})
					if err != nil {
						errs = append(errs, errors.New(table_name+":"+err.Error()))
						return
					}
				}
			}
			if len(index) > 0 {

				res, err = db.QueryString(`show index from ` + table_name)
				if err != nil {
					errs = append(errs, errors.New(table_name+":"+err.Error()))
					return
				}
				keys := map[string]bool{}
				for _, v := range res {
					if v["Key_name"] == "PRIMARY" {
						continue
					}
					keys[v["Key_name"]] = true
				}
				for k := range index {
					if !keys[k] {
						buf.Reset()
						buf.WriteString("ALTER TABLE ")
						buf.WriteString(table_name)
						buf.WriteString(" ADD INDEX ")
						buf.WriteString(k)
						buf.WriteString(" (`")
						buf.WriteString(k)
						buf.WriteString("`)")
						err = Exec(buf.Bytes(), nil, db, &Transaction{})
						if err != nil {
							errs = append(errs, errors.New(table_name+":"+err.Error()))
							return
						}
					}
				}
			}

		}(v)
	}
	wg.Wait()
	return errs
}

//处理sql语句防注入不带'
func Getkey(str_i interface{}) string {
	switch str_i.(type) {
	case string:
		return "`" + key_srp.Replace(str_i.(string)) + "`"
	default:
		r := reflect.TypeOf(str_i)
		for r.Kind() == reflect.Ptr {
			r = r.Elem()
		}
		DEBUG("getkey不支持类型", r.Kind(), r.Name())
	}
	return ""
}
func Getvalue(str_i interface{}) string {
	switch v := str_i.(type) {
	case int8:

		return strconv.Itoa(int(v))
	case int:
		return strconv.Itoa(v)
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.Itoa(int(v))
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case []byte:
		return encodeHex(v)

	case float32, float64:
		return fmt.Sprint(str_i)
	case bool:
		if v {
			return "1"
		} else {
			return "0"
		}
	case time.Time:
		return v.In(mysqlLoc).Format("2006-01-02 15:04:05")
	case string:
		return "'" + value_srp.Replace(v) + "'"
	case []string: //判断是否exp表达式
		if len(v) == 2 && v[0] == "exp" {
			return v[1]
		}
		return marsha1Tostring(str_i)
	default:
		return marsha1Tostring(str_i)

	}
	return ""
}
func marsha1Tostring(i interface{}) (str string) {
	r := reflect.TypeOf(i)
	for r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	if r.Kind() == reflect.Struct || r.Kind() == reflect.Map || r.Kind() == reflect.Slice {
		str = "'" + JsonMarshalString(i) + "'"
	} else {
		str = "'" + fmt.Sprint(i) + "'"
	}
	return
}
func encodeHex(bin []byte) string {
	if len(bin) == 0 {
		return "''"
	}
	return "0x" + hex.EncodeToString(bin)
}

var key_srp = strings.NewReplacer(`\`, `\\`, "`rank`", "rank", "`type`", "type", "`", "\\`")
var value_srp = strings.NewReplacer(`\`, `\\`, "'", `\'`)

func GetvaluefromPtr(ptr uintptr, field reflect.StructField) string {
	switch field.Type.Kind() {
	case reflect.String:
		return "'" + value_srp.Replace(*(*string)(unsafe.Pointer(ptr + field.Offset))) + "'"
	case reflect.Int64:
		return strconv.FormatInt(*(*int64)(unsafe.Pointer(ptr + field.Offset)), 10)
	case reflect.Int32:
		return strconv.FormatInt(int64(*(*int32)(unsafe.Pointer(ptr + field.Offset))), 10)
	case reflect.Int8:
		return strconv.FormatInt(int64(*(*int8)(unsafe.Pointer(ptr + field.Offset))), 10)
	case reflect.Int:
		return strconv.FormatInt(int64(*(*int)(unsafe.Pointer(ptr + field.Offset))), 10)
	case reflect.Int16:
		return strconv.FormatInt(int64(*(*int16)(unsafe.Pointer(ptr + field.Offset))), 10)
	case reflect.Uint:
		return strconv.FormatUint(uint64(*(*uint)(unsafe.Pointer(ptr + field.Offset))), 10)
	case reflect.Uint8:
		return strconv.FormatUint(uint64(*(*uint8)(unsafe.Pointer(ptr + field.Offset))), 10)
	case reflect.Uint16:
		return strconv.FormatUint(uint64(*(*uint16)(unsafe.Pointer(ptr + field.Offset))), 10)
	case reflect.Uint32:
		return strconv.FormatUint(uint64(*(*uint32)(unsafe.Pointer(ptr + field.Offset))), 10)
	case reflect.Uint64:
		return strconv.FormatUint((*(*uint64)(unsafe.Pointer(ptr + field.Offset))), 10)
	case reflect.Float32:
		return fmt.Sprint((*(*float32)(unsafe.Pointer(ptr + field.Offset))))
	case reflect.Float64:
		return fmt.Sprint((*(*float64)(unsafe.Pointer(ptr + field.Offset))))
	case reflect.Bool:
		if *(*bool)(unsafe.Pointer(ptr + field.Offset)) {
			return "1"
		} else {
			return "0"
		}
	case reflect.Struct:
		if field.Type.String() == "time.Time" {

			return "'" + (*(*time.Time)(unsafe.Pointer(ptr + field.Offset))).In(mysqlLoc).Format("2006-01-02 15:04:05") + "'"
		}
		fallthrough

	default:
		t := field.Type
		if t.Kind() == reflect.Ptr {
			if *(*uintptr)(unsafe.Pointer(ptr + field.Offset)) == 0 {
				return "'null'"
			}
		}
		for t.Kind() == reflect.Ptr {

			t = t.Elem()
		}
		i := reflect.NewAt(t, unsafe.Pointer(ptr+field.Offset))
		var str string
		if t.Kind() == reflect.Struct || t.Kind() == reflect.Map || t.Kind() == reflect.Slice {
			str = JsonMarshalString(i.Interface())
		} else {
			str = fmt.Sprint(i)
		}
		return "'" + str + "'"
	}
	return ""
}
func (db *MysqlDB) AriaSetting(TRANSACTIONAL, PAGE_CHECKSUM, TABLE_CHECKSUM bool, ROW_FORMAT string) *MysqlDB {
	db.storeEngine.aria = &ariasetting{
		TRANSACTIONAL:  TRANSACTIONAL,
		PAGE_CHECKSUM:  PAGE_CHECKSUM,
		TABLE_CHECKSUM: TABLE_CHECKSUM,
		ROW_FORMAT:     ROW_FORMAT,
	}
	return db
}
func checkKind(k reflect.Kind) error {
	switch k {
	case reflect.Int:

	case reflect.Int8:

	case reflect.Int16:

	case reflect.Int32:

	case reflect.Int64:

	case reflect.Uint:

	case reflect.Uint8:

	case reflect.Uint16:

	case reflect.Uint32:

	case reflect.Uint64:

	case reflect.Float32:

	case reflect.Float64:

	case reflect.String:

	case reflect.Bool:

	case reflect.Struct:

	case reflect.Slice, reflect.Map:

	case reflect.Ptr:

	default:
		return errors.New("不支持的格式")
	}
	return nil
}
func binaryToStr(column MysqlColumn, data []byte, pos *int, row *MysqlRows) (str string, err error) {
	// Convert to byte-coded string

	switch column.fieldtype {
	case fieldTypeNULL:
		return "NULL", nil
	// Numeric Types
	case fieldTypeTiny:
		if column.fleldflag&flagUnsigned != 0 {
			str = strconv.Itoa(int(data[*pos]))
		} else {
			str = strconv.Itoa(int(int8(data[*pos])))
		}
		*pos++
		return

	case fieldTypeShort, fieldTypeYear:
		if column.fleldflag&flagUnsigned != 0 {
			str = strconv.Itoa(int(binary.LittleEndian.Uint16(data[*pos : *pos+2])))
		} else {
			str = strconv.Itoa(int(int16(binary.LittleEndian.Uint16(data[*pos : *pos+2]))))
		}
		*pos += 2
		return

	case fieldTypeInt24, fieldTypeLong:
		if column.fleldflag&flagUnsigned != 0 {
			str = strconv.Itoa(int(binary.LittleEndian.Uint32(data[*pos : *pos+4])))
		} else {
			str = strconv.Itoa(int(int32(binary.LittleEndian.Uint32(data[*pos : *pos+4]))))
		}
		*pos += 4
		return

	case fieldTypeLongLong:
		if column.fleldflag&flagUnsigned != 0 {
			val := binary.LittleEndian.Uint64(data[*pos : *pos+8])
			if val > math.MaxInt64 {
				str = string(uint64ToString(val))
			} else {
				str = strconv.Itoa(int(val))
			}
		} else {
			str = strconv.Itoa(int(binary.LittleEndian.Uint64(data[*pos : *pos+8])))
		}
		*pos += 8
		return

	case fieldTypeFloat:
		str = strconv.FormatFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(data[*pos:*pos+4]))), 'f', -1, 32)
		*pos += 4
		return

	case fieldTypeDouble:
		str = strconv.FormatFloat(math.Float64frombits(binary.LittleEndian.Uint64(data[*pos:*pos+8])), 'f', -1, 64)
		*pos += 8
		return

	// Length coded Binary Strings
	case fieldTypeDecimal, fieldTypeNewDecimal, fieldTypeVarChar,
		fieldTypeBit, fieldTypeEnum, fieldTypeSet, fieldTypeTinyBLOB,
		fieldTypeMediumBLOB, fieldTypeLongBLOB, fieldTypeBLOB,
		fieldTypeVarString, fieldTypeString, fieldTypeGeometry, fieldTypeJSON:

		msglen, err := ReadLength_Coded_Slice(data[*pos:], pos)
		if err != nil {
			return "", err
		}

		if msglen == 0 {
			str = "NULL"
		} else {
			str = string(data[*pos : *pos+msglen])
		}

		*pos += msglen
	case
		fieldTypeDate, fieldTypeNewDate, // Date YYYY-MM-DD
		fieldTypeTime,                         // Time [-][H]HH:MM:SS[.fractal]
		fieldTypeTimestamp, fieldTypeDateTime: // Timestamp YYYY-MM-DD HH:MM:SS[.fractal]

		n, err := ReadLength_Coded_Slice(data[*pos:], pos)
		if err != nil {
			return "", err
		}

		switch {
		case n == 0:
			str = "NULL"
			return str, nil
		case column.fieldtype == fieldTypeTime:
			// database/sql does not support an equivalent to TIME, return a string
			var dstlen uint8
			switch decimals := column.decimals; decimals {
			case 0x00, 0x1f:
				dstlen = 8
			case 1, 2, 3, 4, 5, 6:
				dstlen = 8 + 1 + decimals
			default:
				return "", fmt.Errorf(
					"protocol error, illegal decimals value %d",
					column.decimals,
				)
			}
			t, err := formatBinaryTime(data[*pos:*pos+int(n)], dstlen)
			if err != nil {
				return "", err
			}
			str = string(t)
			//case columns.conn.parseTime:
		default:
			t, err := parseBinaryDateTime(uint64(n), data[*pos:], row.conn.loc)
			if err != nil {
				return "", err
			}
			str = t.Format("2006-01-02 15:04:05")
			/*default:
			var dstlen uint8
			if column.fieldtype == fieldTypeDate {
				dstlen = 10
			} else {
				switch decimals := column.decimals; decimals {
				case 0x00, 0x1f:
					dstlen = 19
				case 1, 2, 3, 4, 5, 6:
					dstlen = 19 + 1 + decimals
				default:
					return nil, fmt.Errorf(
						"protocol error, illegal decimals value %d",
						column.decimals,
					)
				}
			}
			t, err := formatBinaryDateTime(data[*pos:*pos+int(n)], dstlen)
			if err != nil {
				return nil, err
			}
			record[key] = string(t)*/
		}

		if err == nil {
			*pos += int(n)
		} else {
			return str, err
		}
	// Please report if this happens!
	default:
		return "", fmt.Errorf("unknown field type %d", column.fieldtype)
	}
	return str, err
}

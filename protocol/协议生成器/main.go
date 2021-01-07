package main

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

type ErrCode int16
type HtmlKeyValueStr struct {
	Key   string
	Value string
}

//本代码是给msg加入chan bytes.Buffer
var msg = make(map[int32]string)

func main() {
	BASE_ROOT_PATH, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("获取当前运行目录失败 " + err.Error())
	}
	DS := string(os.PathSeparator)
	if strings.Count(BASE_ROOT_PATH, "mp"+DS+"go-build") > 0 { //go run模式
		BASE_ROOT_PATH, err = os.Getwd()
		if err != nil {
			panic("获取当前运行目录失败 " + err.Error())
		}
	}
	new_gopatch := BASE_ROOT_PATH + DS + ".." + DS
	exist, _ := PathExists(new_gopatch)
	if !exist {
		err := os.Mkdir(new_gopatch, os.ModePerm)
		if err != nil {
			Log("创建文件夹失败 %s", new_gopatch)
			return
		}
	}

	new_jspatch := BASE_ROOT_PATH + DS + "js"
	exist, _ = PathExists(new_jspatch)
	if !exist {
		err := os.Mkdir(new_jspatch, os.ModePerm)
		if err != nil {
			Log("创建文件夹失败 %s", new_jspatch)
			return
		}
	}
	list, _ := ListDir(new_jspatch, "js")
	for _, name := range list {
		err := os.Remove(name)
		if err != nil {
			Log("无法删除文件%s", name)
			return
		}
	}
	h := "\n" //换行符

	out := new(bytes.Buffer)
	for serverid, name := range []string{"0common.go", "", "", "", "4user.go", "5project.go", "6test.go"} {
		func() {
			name = BASE_ROOT_PATH + DS + name
			b, err := ioutil.ReadFile(name)
			if err != nil {
				return
			}
			str := strings.Replace(string(b), "\r\n", h, -1)
			m, err := Preg_match_result(`type ([^\s]+)\s* struct\s*{\s*(((?!})[\s\S])*)}`, str, -1)
			r := []*go_struct{}
			for _, val := range m {
				if val[1] == "" {
					for k, v := range val {
						DEBUG(k, v)
					}
				}

				_struct := &go_struct{name: val[1]}
				val[2] = strings.Trim(val[2], h)
				if len(val[2]) > 5 {
					for _, filed_s := range strings.Split(val[2], h) {

						if _m, err := Preg_match_result(`\s*([^\s]+)\s+([^\n/]+?)\s*(`+"`"+"[^`]+`"+`)?\s*(\/\/[^\n]+)?$`, filed_s, 1); len(_m) > 0 {
							if err != nil {
								DEBUG(err)
							}

							if _m[0][1][:2] == "//" {
								continue
							}

							_struct.field = append(_struct.field, struct {
								name string
								tag  string
								typ  string
							}{name: _m[0][1], typ: strings.Trim(_m[0][2], " "), tag: _m[0][3]})
						} else {
							Log(name+"匹配结构体 %s 成员%s错误", val[1], filed_s)
							continue
						}
					}
				}
				r = append(r, _struct)
			}
			out.Reset()
			out.WriteString("package protocol\n\nimport (\n	\"sync\"\n	\"libraries\"\n")
			if strings.Contains(str, "time.Time") {
				out.WriteString("	\"time\"\n")
			}
			out.WriteString(")\n\n")
			out.WriteString("const (\n")
			for _, _struct := range r {
				out.WriteString("	")
				out.WriteString("CMD_")
				out.WriteString(_struct.name)
				out.WriteString(" = ")
				crc32cmd := int32(crc32.ChecksumIEEE([]byte(_struct.name)))
				//去掉第一位Byte，添加上serverid
				cmd := crc32cmd - crc32cmd&255 + int32(serverid)
				for {
					if _, ok := msg[cmd]; !ok {
						break
					}
					cmd++
				}
				msg[cmd] = _struct.name
				_struct.cmd = strconv.Itoa(int(cmd))
				out.WriteString(_struct.cmd)
				out.WriteString(h)
			}
			out.WriteString(")\n\n")

			for _, _struct := range r {
				out.WriteString("type ")
				out.WriteString(_struct.name)
				out.WriteString(" struct {\n")
				for _, f := range _struct.field {
					out.WriteString("	")
					out.WriteString(f.name)
					out.WriteString(" ")
					out.WriteString(f.typ)
					if f.tag != "" {
						out.WriteString(" ")
						out.WriteString(f.tag)
					}
					out.WriteString(h)
				}
				out.WriteString("}\n\n")

				out.WriteString("var pool_")
				out.WriteString(_struct.name)
				out.WriteString(" = sync.Pool{New: func() interface{} { return &")
				out.WriteString(_struct.name)
				out.WriteString("{} }}\n\n")
				out.WriteString(`func GET_`)
				out.WriteString(_struct.name)
				out.WriteString(`() *`)
				out.WriteString(_struct.name)
				out.WriteString(" {\n	return pool_")
				out.WriteString(_struct.name)
				out.WriteString(".Get().(*")
				out.WriteString(_struct.name)
				out.WriteString(")\n}\n\n")
				out.WriteString("func (data *")
				out.WriteString(_struct.name)
				out.WriteString(") cmd() int32 {\n	return CMD_")
				out.WriteString(_struct.name)
				out.WriteString("\n}\n\nfunc (data *")
				out.WriteString(_struct.name)
				out.WriteString(") Put() {\n")
				for _, f := range _struct.field {
					if strings.Contains(f.typ, "*") {
						if f.typ[:2] == "[]" {

							out.WriteString("	for _,v := range data.")
							out.WriteString(f.name)
							out.WriteString(" {\n		v.Put()\n")
							out.WriteString("	}\n	data.")
							out.WriteString(f.name)
							out.WriteString(" = data.")
							out.WriteString(f.name)
							out.WriteString("[:0]\n")
						} else {
							out.WriteString("	if data.")
							out.WriteString(f.name)
							out.WriteString(" != nil {\n")
							out.WriteString("		data.")
							out.WriteString(f.name)
							out.WriteString(".Put()\n")
							out.WriteString("		data.")
							out.WriteString(f.name)
							out.WriteString(" = nil\n")
							out.WriteString("	}\n")
						}
					} else {
						switch f.typ {
						case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "ErrCode":
							out.WriteString("	data.")
							out.WriteString(f.name)
							out.WriteString(" = 0\n")
						case "string":
							out.WriteString("	data.")
							out.WriteString(f.name)
							out.WriteString(" = ``\n")
						case "bool":
							out.WriteString("	data.")
							out.WriteString(f.name)
							out.WriteString(" = false\n")
						case "time.Time":
							out.WriteString("	data.")
							out.WriteString(f.name)
							out.WriteString(" = time.Unix(0,0)\n")
						default:
							if f.typ[:2] == "[]" {
								out.WriteString("	data.")
								out.WriteString(f.name)
								out.WriteString(" = data.")
								out.WriteString(f.name)
								out.WriteString("[:0]\n")
							} else if len(f.typ) > 4 && f.typ[:4] == "map[" {
								out.WriteString("	data.")
								out.WriteString(f.name)
								out.WriteString(" = nil\n")
							} else {
								DEBUG(_struct.name, f.name, "put未设置类型'", f.typ, "'")
							}

						}
					}
				}
				out.WriteString("	pool_")
				out.WriteString(_struct.name)
				out.WriteString(".Put(data)\n}\n")
				//var pool_MSG_LS2B_CommandResult = sync.Pool{New: func() interface{} { return &MSG_LS2B_CommandResult{} }}
				out.WriteString("func (data *")
				out.WriteString(_struct.name)
				out.WriteString(") write(buf *libraries.MsgBuffer) {\n	WRITE_int32(CMD_")
				out.WriteString(_struct.name)
				out.WriteString(",buf)\n	WRITE_")
				out.WriteString(_struct.name)
				out.WriteString("(data, buf)\n}\n\nfunc WRITE_")
				out.WriteString(_struct.name)
				out.WriteString("(data *")
				out.WriteString(_struct.name)
				out.WriteString(", buf *libraries.MsgBuffer) {\n")
				for _, f := range _struct.field {
					switch {
					case f.typ == "[]byte":
						out.WriteString("	WRITE_int32(int32(len(data.")
						out.WriteString(f.name)
						out.WriteString(")), buf)\n")
						out.WriteString("	buf.Write(data.")
						out.WriteString(f.name)
						out.WriteString(")\n")
					case f.typ == "[][]byte":
						out.WriteString("	WRITE_int32(int32(len(data.")
						out.WriteString(f.name)
						out.WriteString(")), buf)\n")
						out.WriteString("	for _, v := range data.")
						out.WriteString(f.name)
						out.WriteString("{\n")
						out.WriteString("		WRITE_int32(int32(len(v)), buf)\n		buf.Write(v)\n")
						out.WriteString("	}\n")
					case f.typ[:2] == "[]":
						out.WriteString("	WRITE_int32(int32(len(data.")
						out.WriteString(f.name)
						out.WriteString(")), buf)\n")
						out.WriteString("	for _, v := range data.")
						out.WriteString(f.name)
						out.WriteString("{\n")
						out.WriteString("		WRITE_")
						out.WriteString(strings.Replace(f.typ[2:], "*", "", 1))
						out.WriteString("(v, buf)\n")
						out.WriteString("	}\n")
					case f.typ == "time.Time":
						out.WriteString("	WRITE_int64(data.")
						out.WriteString(f.name)
						out.WriteString(".UnixNano(), buf)\n")
					case len(f.typ) > 4 && f.typ[:4] == "map[":
						out.WriteString("	WRITE_map(data.")
						out.WriteString(f.name)
						out.WriteString(",buf)\n")
					default:
						if strings.Contains(f.typ, "*") {
							out.WriteString("	if data.")
							out.WriteString(f.name)
							out.WriteString(" == nil {\n")
							out.WriteString("		WRITE_int8(0, buf)\n")
							out.WriteString("	} else {\n		WRITE_int8(1, buf)\n		WRITE_")
							out.WriteString(f.typ[1:])
							out.WriteString("(data.")
							out.WriteString(f.name)
							out.WriteString(", buf)\n")
							out.WriteString("	}\n")
						} else {
							out.WriteString("	WRITE_")
							out.WriteString(strings.Replace(f.typ, "*", "", 1))
							out.WriteString("(data.")
							out.WriteString(f.name)
							out.WriteString(", buf)\n")
						}
					}
				}

				out.WriteString("}\n\n")
				out.WriteString("func READ_")
				out.WriteString(_struct.name)
				out.WriteString("(buf *libraries.MsgBuffer) *")
				out.WriteString(_struct.name)
				out.WriteString(" {\n	data := pool_")
				out.WriteString(_struct.name)
				out.WriteString(".Get().(*")
				out.WriteString(_struct.name)
				out.WriteString(")\n	data.read(buf)\n	return data\n}\n\nfunc (data *")
				out.WriteString(_struct.name)
				out.WriteString(") read(buf *libraries.MsgBuffer) {\n")
				/*Platform_listlength, bin := READ_int16(bin)
				for i := 0; i < int(*Platform_listlength); i++ {
					data, bin1 := READ_MSG_PlatformInfo(bin)
					bin = bin1
					*Platform_list = append(*Platform_list, *data)
				}*/
				var hasQueryID, hasResultQueryID bool
				for _, f := range _struct.field {
					switch {
					case f.typ == "[]byte":
						l := f.name + "_len"
						out.WriteString("	")
						out.WriteString(l)
						out.WriteString(" := int(READ_int32(buf))\n")
						out.WriteString("	data.")
						out.WriteString(f.name)
						out.WriteString(" = make(")
						out.WriteString(f.typ)
						out.WriteString(", ")
						out.WriteString(l)
						out.WriteString(")\n	copy(data.")
						out.WriteString(f.name)
						out.WriteString(",buf.Next(")
						out.WriteString(l)
						out.WriteString("))\n")
					case f.typ == "[][]byte":
						l := f.name + "_len"
						out.WriteString("	")
						out.WriteString(l)
						out.WriteString(" := int(READ_int32(buf))\n	for i := 0; i < ")
						out.WriteString(l)
						out.WriteString("; i++ {\n		l := READ_int32(buf)\n		b := make([]byte,l)\n		copy(b,buf.Next(int(l)))\n")
						out.WriteString("		data.")
						out.WriteString(f.name)
						out.WriteString(" = append(data.")
						out.WriteString(f.name)
						out.WriteString(", b)\n")
						out.WriteString("	}\n")
					case f.typ[:2] == "[]":
						l := f.name + "_len"
						out.WriteString("	")
						out.WriteString(l)
						out.WriteString(" := int(READ_int32(buf))\n	for i := 0; i < ")
						out.WriteString(l)
						out.WriteString("; i++ {\n")
						out.WriteString("		data.")
						out.WriteString(f.name)
						out.WriteString(" = append(data.")
						out.WriteString(f.name)
						out.WriteString(", READ_")
						out.WriteString(strings.Replace(f.typ[2:], "*", "", 1))
						out.WriteString("(buf))\n")
						out.WriteString("	}\n")
					case f.typ == "time.Time":
						out.WriteString("	data.")
						out.WriteString(f.name)
						out.WriteString(" = time.Unix(0, READ_int64(buf))\n")
					case len(f.typ) > 4 && f.typ[:4] == "map[":
						out.WriteString("	READ_map(&data.")
						out.WriteString(f.name)
						out.WriteString(",buf)\n")
					default:

						if strings.Contains(f.typ, "*") {
							l := f.name + "_len"
							out.WriteString("	")
							out.WriteString(l)
							out.WriteString(" := int(READ_int8(buf))\n")
							out.WriteString("	if ")
							out.WriteString(l)
							out.WriteString(" == 1 {\n")
							out.WriteString("		data.")
							out.WriteString(f.name)
							out.WriteString(" = READ_")
							out.WriteString(strings.Replace(f.typ, "*", "", 1))
							out.WriteString("(buf)\n")
							out.WriteString("	}else{\n		data.")
							out.WriteString(f.name)
							out.WriteString(" = nil\n	}\n")
						} else {
							out.WriteString("	data.")
							out.WriteString(f.name)
							out.WriteString(" = READ_")
							out.WriteString(f.typ)
							out.WriteString("(buf)\n")
							if f.name == "QueryID" && f.typ == "uint32" {
								hasQueryID = true
							} else if f.name == "QueryResultID" && f.typ == "uint32" {
								hasResultQueryID = true
							}
						}

					}

				}
				out.WriteString("\n}\n")
				if hasQueryID {
					out.WriteString("func (data *")
					out.WriteString(_struct.name)
					out.WriteString(") getQueryID() uint32 {\n")
					out.WriteString("	return data.QueryID\n}\n")
					out.WriteString("func (data *")
					out.WriteString(_struct.name)
					out.WriteString(") setQueryID(id uint32) {\n")
					out.WriteString("	data.QueryID = id\n}\n")
				} else if hasResultQueryID {
					out.WriteString("func (data *")
					out.WriteString(_struct.name)
					out.WriteString(") getQueryResultID() uint32 {\n")
					out.WriteString("	return data.QueryResultID\n}\n")
					out.WriteString("func (data *")
					out.WriteString(_struct.name)
					out.WriteString(") setQueryResultID(id uint32) {\n")
					out.WriteString("	data.QueryResultID = id\n}\n")
				}
				out.WriteString("\n")
			}
			new_name := strings.Replace(name, BASE_ROOT_PATH, new_gopatch, 1)
			os.Remove(new_name)
			n, err := os.OpenFile(new_name, os.O_CREATE|os.O_WRONLY, 777)
			if err != nil {
				panic("新建文件错误" + new_name + err.Error())
			}

			n.Write(out.Bytes())
			n.Sync()
			n.Close()
			out.Truncate(0)
			return //js部分未完善
			out.WriteString("var ")
			for _, _struct := range r {
				out.WriteString("")
				out.WriteString("CMD_")
				out.WriteString(_struct.name)
				out.WriteString(" = ")
				out.WriteString(_struct.cmd)
				out.WriteString(",")
			}
			out.Truncate(out.Len() - 1)
			out.WriteString(";\n")

			for _, _struct := range r {
				out.WriteString("function WRITE_")
				out.WriteString(_struct.name)
				out.WriteString("(o){\n	var b=[];\n")
				out.WriteString("	b=b.concat(write_int32(CMD_")
				out.WriteString(_struct.name)
				out.WriteString("));\n")
				out.WriteString("	b=b.concat(write_")
				out.WriteString(_struct.name)
				out.WriteString("(o));\n")
				out.WriteString("	return b;\n")
				out.WriteString("}\n")
				out.WriteString("function write_")
				out.WriteString(_struct.name)
				out.WriteString("(o){\n	var b=[];\n")
				for _, f := range _struct.field {
					if f.typ == "[]byte" {
						out.WriteString("	b=b.concat(write_byte(o.")
						out.WriteString(f.name)
						out.WriteString("));\n")
					} else if f.typ[:2] == "[]" {
						out.WriteString("	if(o.")
						out.WriteString(f.name)
						out.WriteString("){\n")
						out.WriteString("		b=b.concat(write_int32(o.")
						out.WriteString(f.name)
						out.WriteString(".length))\n")
						out.WriteString("		for(var i=0;i<o.")
						out.WriteString(f.name)
						out.WriteString(".length;i++){\n")
						t := strings.Replace(f.typ[2:], "*", "", 1)
						out.WriteString("			b=b.concat(write_")
						out.WriteString(t)
						out.WriteString("(o.")
						out.WriteString(f.name)
						out.WriteString("[i]));\n		}\n	}else{\n		b=b.concat([0,0])\n	}\n")
					} else if strings.Contains(f.typ, "*") {
						out.WriteString("	if(o.")
						out.WriteString(f.name)
						out.WriteString("){\n")
						out.WriteString("		b=b.concat(write_int8(1))\n")
						out.WriteString("		b=b.concat(write_")
						out.WriteString(strings.Replace(f.typ, "*", "", 1))
						out.WriteString("(o.")
						out.WriteString(f.name)
						out.WriteString("));\n")
						out.WriteString("	}else{\n")
						out.WriteString("		b=b.concat(write_int8(0))\n	}\n")
					} else {
						out.WriteString("	b=b.concat(write_")
						out.WriteString(f.typ)
						out.WriteString("(o.")
						out.WriteString(f.name)
						out.WriteString("));\n")
					}
				}
				out.WriteString("	return b\n")
				out.WriteString("}\n")
				out.WriteString("function read_")
				out.WriteString(_struct.name)
				out.WriteString("(b){\n")
				out.WriteString("	var o={},r={};r.b=b;\n")
				for _, f := range _struct.field {
					if f.typ == "[]byte" {
						out.WriteString("	r=read_byte(r.b);o.")
						out.WriteString(f.name)
						out.WriteString("=r.o\n")
					} else if f.typ[:2] == "[]" {
						out.WriteString("	r=read_int32(r.b);var l=r.o;if(l>0)o.")
						out.WriteString(f.name)
						out.WriteString("=[]\n")
						out.WriteString("	for(var i=0;i<l;i++){\n")
						out.WriteString("		r=read_")
						out.WriteString(strings.Replace(f.typ[2:], "*", "", 1))
						out.WriteString("(r.b)\n")
						out.WriteString("		o.")
						out.WriteString(f.name)
						out.WriteString(".push(r.o)\n")
						out.WriteString("	}\n")

					} else if strings.Contains(f.typ, "*") {
						out.WriteString("	r=read_int8(r.b);var l=r.o;\n	if(l>0){\n")
						out.WriteString("		r=read_")
						out.WriteString(strings.Replace(f.typ, "*", "", 1))
						out.WriteString("(r.b);\n")
						out.WriteString("		o.")
						out.WriteString(f.name)
						out.WriteString("=r.o\n")
						out.WriteString("	}\n")
					} else {
						out.WriteString("	r=read_")
						out.WriteString(f.typ)
						out.WriteString("(r.b);\n")
						out.WriteString("	o.")
						out.WriteString(f.name)
						out.WriteString("=r.o\n")
					}
				}
				out.WriteString("	return {o:o,b:r.b}\n}\n")
			}
			/*function read_MSG_LS2B_CommandResult(b){
				var o={},r={};
				r = read_int16(b)
				o.Code = r.o
				b = r.b
				r = read_string(b)
				o.Msg = r.o
				b = r.b
				return o
			}*/
			new_name = strings.Replace(strings.Replace(name, BASE_ROOT_PATH, new_jspatch, 1), ".go", ".js", 1)
			n, _ = os.OpenFile(new_name, os.O_CREATE, 0)
			n.Write(out.Bytes())
			n.Close()

		}()

	}
	out.Reset()
	out.WriteString("package protocol\n\nimport (\n	\"libraries\"\n)\n\nvar cmdMapFunc = map[int32]func(*libraries.MsgBuffer) MSG_DATA{\n")
	for cmd, name := range msg {
		out.WriteString("	")
		out.WriteString(strconv.Itoa(int(cmd)))
		out.WriteString(": func(buf *libraries.MsgBuffer) MSG_DATA {return MSG_DATA(READ_")
		out.WriteString(name)
		out.WriteString("(buf))},\n")
	}
	out.WriteString("}\n\nvar CmdToName = map[int32]string{\n")
	for cmd, name := range msg {
		out.WriteString("	")
		out.WriteString(strconv.Itoa(int(cmd)))
		out.WriteString(": `")
		out.WriteString(name)
		out.WriteString("`,\n")
	}
	out.WriteString("}")
	new_name := new_gopatch + DS + "cmdMapFunc.go"
	os.Remove(new_name)
	n, _ := os.OpenFile(new_name, os.O_CREATE|os.O_WRONLY, 777)
	n.Write(out.Bytes())
	n.Close()
	return //跳过js
	out.Reset()
	out.WriteString("function read_msg(b){\n")
	out.WriteString("	var cmd = read_int32(b),r={};\n	switch(cmd.o){\n")
	for cmd, name := range msg {
		out.WriteString("	case ")
		out.WriteString(strconv.Itoa(int(cmd)))
		out.WriteString(":\n")
		out.WriteString("		r=read_")
		out.WriteString(name)
		out.WriteString("(cmd.b)\n		break\n")
	}
	out.WriteString("	}\n	return {cmd:cmd.o,msg:r.o}\n}")
	n, _ = os.OpenFile(new_jspatch+"/handle.js", os.O_CREATE, 0)
	n.Write(out.Bytes())
	n.Close()
}

type go_struct struct {
	name  string
	cmd   string
	field []struct {
		name string
		tag  string
		typ  string
	}
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func DEBUG(v ...interface{}) {

	_, file, line, ok := runtime.Caller(1)
	if ok {
		v = append([]interface{}{fmt.Sprintf("%s,line %d:", file, line)}, v...)
	}
	fmt.Println(v...)

}
func Log(format string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		v = append([]interface{}{fmt.Sprintf("%s,line %d:", file, line)}, v...)
	}
	fmt.Printf("%s "+format+"\r\n", v...)
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := "/"
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		PthSep = "\\"
	}
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

//返回匹配结果,n=次数
func Preg_match_result(regtext string, text string, n int) ([][]string, error) {

	r, err := regexp2.Compile(regtext, 0)
	if err != nil {
		return nil, err
	}

	m, err := r.FindStringMatch(text)
	if err != nil {
		return nil, err
	}
	var result [][]string
	for m != nil && n != 0 {
		var res_v []string
		for _, v := range m.Groups() {
			res_v = append(res_v, v.String())
		}

		m, _ = r.FindNextMatch(m)
		result = append(result, res_v)
		n--
	}

	return result, nil
}

package codec

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"io"
	"io/ioutil"
	"libraries"
	"math/rand"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/klauspost/compress/gzip"
	"github.com/luyu6056/cache"
	"github.com/luyu6056/gnet"
)

type Httpserver struct {
	session        *cache.Hashvalue
	Request        Request
	c              gnet.Conn
	Out            *libraries.MsgBuffer
	Ws             *WSconn
	data           *bytes.Reader
	OutCode        int
	OutContentType string
	ishttps        bool
	Origin         string
	OutHeader      map[string]string
	OutCookie      map[string]httpcookie
	//StartTime      time.Time
}
type httpcookie struct {
	value   string
	max_age uint32
}
type httpQuery struct {
	key   string
	value []string
}
type Request struct {
	Code             int
	CodeMsg          string
	Proto, Method    string
	path, query, uri string
	RemoteAddr       string
	Connection       string
	Header           map[string]string
	Cookie           map[string]string
	Body             *libraries.MsgBuffer
	queryS           []httpQuery
	postS            []httpQuery
	IsHttps          bool
}

var Httppool = sync.Pool{New: func() interface{} {
	hs := &Httpserver{Out: new(libraries.MsgBuffer)}
	hs.OutHeader = make(map[string]string)
	hs.OutCookie = make(map[string]httpcookie)
	hs.Request.Header = make(map[string]string)
	hs.Request.Body = &libraries.MsgBuffer{}
	hs.data = &bytes.Reader{}
	return hs
}}

var msgbufpool = sync.Pool{New: func() interface{} {
	return new(libraries.MsgBuffer)
}}
var gzippool = sync.Pool{New: func() interface{} {
	w, _ := gzip.NewWriterLevel(nil, 6)
	return w
}}

func (r *Request) GetHeader(key string) string {
	return r.Header[key]
}

func (hs *Httpserver) BeginRequest(route string, c *Context) {

}
func (hs *Httpserver) Write(b *libraries.MsgBuffer) {
	hs.Out.Reset()
	if hs.OutCode != 0 && httpCode(hs.OutCode).String() != "" {
		hs.Out.WriteString("HTTP/1.1 ")
		hs.Out.WriteString(httpCode(hs.OutCode).String())
		hs.Out.WriteString("\r\n")
	} else {
		hs.Out.Write(http1head200)
	}
	hs.Out.Write(http1nocache)
	if hs.OutContentType != "" {
		hs.Out.WriteString("Content-Type: ")
		hs.Out.WriteString(hs.OutContentType)
		hs.Out.WriteString("\r\n")
	} else {
		hs.Out.WriteString("Content-Type: text/html;charset=utf-8\r\n")
	}

	if b.Len() > 9192 && strings.Contains(hs.Request.GetHeader("Accept-Encoding"), "deflate") {
		buf := msgbufpool.Get().(*libraries.MsgBuffer)
		buf.Reset()
		w := CompressNoContextTakeover(buf, 6)
		w.Write(b.Bytes())
		w.Close()
		hs.Out.Write(http1deflate)
		hs.httpsfinish(buf, buf.Len())
		buf.Reset()
		msgbufpool.Put(buf)
	} else {
		hs.httpsfinish(b, b.Len())
	}

}
func (hs *Httpserver) WriteString(str string) {
	hs.Out.Reset()
	if hs.OutCode != 0 && httpCode(hs.OutCode).String() != "" {
		hs.Out.WriteString("HTTP/1.1 ")
		hs.Out.WriteString(httpCode(hs.OutCode).String())
		hs.Out.WriteString("\r\n")
	} else {
		hs.Out.Write(http1head200)
	}
	hs.Out.Write(http1nocache)
	if hs.OutContentType != "" {
		hs.Out.WriteString("Content-Type: ")
		hs.Out.WriteString(hs.OutContentType)
		hs.Out.WriteString("\r\n")
	} else {
		hs.Out.WriteString("Content-Type: text/html;charset=utf-8\r\n")
	}
	buf := msgbufpool.Get().(*libraries.MsgBuffer)
	buf.Reset()
	if len(str) > 9192 && strings.Contains(hs.Request.GetHeader("Accept-Encoding"), "deflate") {
		w := CompressNoContextTakeover(buf, 6)
		w.Write(libraries.Str2bytes(str))
		w.Close()
		hs.Out.Write(http1deflate)
	} else {
		buf.WriteString(str)
	}

	hs.httpsfinish(buf, buf.Len())
	buf.Reset()
	msgbufpool.Put(buf)
}
func (hs *Httpserver) RemoteAddr() string {
	return hs.c.RemoteAddr().String()
}
func (hs *Httpserver) IP() (ip string) {

	if ip = hs.Request.GetHeader("X-Real-IP"); ip == "" {
		ip = hs.c.RemoteAddr().String()
	}
	re3, _ := regexp.Compile(`:\d+$`)
	ip = re3.ReplaceAllString(ip, "")
	return ip
}
func (hs *Httpserver) IsMobile() bool {
	return false
}

func (hs *Httpserver) UserAgent() string {
	return hs.Request.GetHeader("UserAgent")
}
func (hs *Httpserver) URI() string {
	if hs.ishttps {
		return fmt.Sprintf("https://%s%s", hs.Request.Header["Host"], hs.Request.uri)
	}
	return fmt.Sprintf("http://%s%s", hs.Request.Header["Host"], hs.Request.uri)
}
func (hs *Httpserver) Referer() string {
	return hs.Request.Header["Referer"]
}

var errprotocol = errors.New("the client is not using the websocket protocol:")

//http升级为websocket
func (hs *Httpserver) Upgradews() (err error) {
	//
	hs.Out.Reset()
	/*if !(strings.Contains(c.Request.Head, "Connection: Upgrade")) {

		hs.Out.WriteString("HTTP/1.1 400 Error\r\nContent-Type: text/plain\r\nContent-Length: 11\r\nConnection: close\r\n\r\nUnknonw MSG")
		libraries.DebugLog("ws协议没有upgrade")
		return errprotocol
	}*/
	if hs.Request.Method != "GET" {

		hs.Out.WriteString("HTTP/1.1 403 Error\r\nContent-Type: text/plain\r\nContent-Length: 11\r\nConnection: close\r\n\r\nUnknonw MSG")
		libraries.DebugLog("ws协议没有get")
		return errprotocol
	}
	/*libraries.DebugLog(c.Request.Head)
	if !(strings.Contains(c.Request.Head, "Sec-WebSocket-Extensions")) {

		hs.Out.WriteString("HTTP/1.1 400 Error\r\nContent-Type: text/plain\r\nContent-Length: 11\r\nConnection: close\r\n\r\nUnknonw MSG")
		libraries.DebugLog("ws协议没有Extensions")
		return
	}*/

	if hs.Origin != "" && hs.Request.Header["Origin"] != hs.Origin {
		hs.Out.WriteString("HTTP/1.1 403 Error\r\nContent-Type: text/plain\r\nContent-Length: 11\r\nConnection: close\r\n\r\nUnknonw MSG")
		libraries.DebugLog("ws来自错误的Origin")
		return errprotocol
	}
	if hs.Request.Header["Upgrade"] != "websocket" {
		hs.Out.WriteString("HTTP/1.1 403 Error\r\nContent-Type: text/plain\r\nContent-Length: 11\r\nConnection: close\r\n\r\nUnknonw MSG")
		libraries.DebugLog("ws协议没有upgrade")
		return errprotocol
	}

	if hs.Request.Header["Sec-WebSocket-Version"] != "13" {
		hs.Out.WriteString("HTTP/1.1 403 Error\r\nContent-Type: text/plain\r\nContent-Length: 11\r\nConnection: close\r\n\r\nUnknonw MSG")
		libraries.DebugLog("ws协议没有Extensions")
		return errprotocol
	}

	var challengeKey string

	if challengeKey = hs.Request.Header["Sec-WebSocket-Key"]; challengeKey == "" {
		hs.Out.WriteString("HTTP/1.1 403 Error\r\nContent-Type: text/plain\r\nContent-Length: 11\r\nConnection: close\r\n\r\nUnknonw MSG")
		libraries.DebugLog("ws协议没有Extensions")
		return errprotocol
	}
	id := atomic.AddInt32(&ClientId, 1)
	hs.Ws = &WSconn{
		IsServer:   true,
		ReadFinal:  true,
		Http:       hs,
		Conn:       &ClientConn{BeginTime: time.Now().Unix(), IP: hs.IP(), UserAgent: hs.Request.GetHeader("User-Agent"), Id: id},
		Write:      hs.c.WriteNoCodec,
		IsCompress: strings.Contains(hs.Request.Header["Sec-WebSocket-Extensions"], "permessage-deflate"),
		readbuf:    &libraries.MsgBuffer{},
	}
	hs.Ws.Conn.Output_data = hs.Ws.Output_data
	hs.c.SetContext(hs.Ws)
	hs.Out.WriteString("HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: ")
	hs.Out.WriteString(ComputeAcceptKey(challengeKey))
	hs.Out.WriteString("\r\n")
	if hs.Ws.IsCompress {
		hs.Out.WriteString("Sec-Websocket-Extensions: permessage-deflate; server_no_context_takeover; client_no_context_takeover\r\n")
	}
	hs.Out.WriteString("\r\n")
	hs.c.WriteNoCodec(hs.Out.Bytes())

	return nil
}

func (req *Request) Parsereq(data []byte) (int, []byte, error) {
	sdata := string(data)
	var i, s int

	var line string
	var clen int
	var q = -1
	// method, path, proto line
	req.Proto = ""
	i = bytes.IndexByte(data, 32)
	if i == -1 {
		return 0, nil, nil
	}
	req.Method = sdata[:i]
	l := len(sdata)
	for i, s = i+1, i+1; i < l; i++ {
		if data[i] == 63 && q == -1 {
			q = i
		} else if data[i] == 32 {
			if q != -1 {
				req.path = sdata[s:q]
				req.query = sdata[q+1 : i]
			} else {
				req.path = sdata[s:i]
			}
			req.uri = sdata[s:i]
			i++
			s = bytes.Index(data[i:], []byte{13, 10})
			if s > -1 {
				s += i
				req.Proto = sdata[i:s]
			}
			//判断http返回
			if req.Method == "HTTP/1.1" || req.Method == "HTTP/1.0" {
				code, err := strconv.Atoi(req.path)
				if err == nil {

					req.Code = code
					req.CodeMsg = req.Proto

				}
				req.Proto = req.Method
				req.Method = ""
				req.path = ""
			}

			break
		}
	}
	switch req.Proto {
	case "HTTP/1.0":
		req.Connection = "close"
	case "HTTP/1.1":
		req.Connection = "keep-alive"
	default:
		return 0, nil, fmt.Errorf("malformed request")
	}
	//fmt.Println(sdata)
	for s += 2; s < l; s += i + 2 {
		i = bytes.Index(data[s:], []byte{13, 10})
		if i == -1 {
			return 0, nil, nil
		}
		line = sdata[s : s+i]
		if i > 15 {
			switch {
			case line[:15] == "Content-Length:", line[:15] == "Content-length:":
				clen, _ = strconv.Atoi(line[16:])
			case line == "Connection: close", line == "Connection: Close":
				req.Connection = "close"
			default:
				j := bytes.IndexByte(data[s:s+i], 58)
				if j == -1 {
					return 0, nil, nil
				}
				req.Header[line[:j]] = line[j+2:]
			}
		} else if i == 0 {
			s += i + 2
			if clen == 0 && req.Header["Transfer-Encoding"] == "chunked" {
				req.Body.Reset()
				for ; s < l; s += 2 {
					i = bytes.Index(data[s:], []byte{13, 10})
					if i == -1 {
						return 0, nil, nil
					}
					b := make([]byte, 8)
					if i&1 == 0 {
						hex.Decode(b[8-i/2:], data[s:s+i])
					} else {
						tmp, _ := hex.DecodeString("0" + sdata[s:s+i])
						copy(b[7-i/2:], tmp)

					}
					clen = int(b[0])<<56 | int(b[1])<<48 | int(b[2])<<40 | int(b[3])<<32 | int(b[4])<<24 | int(b[5])<<16 | int(b[6])<<8 | int(b[7])
					s += i + 2
					if l-s < clen {
						return 0, nil, nil
					}
					if clen > 0 {
						req.Body.Write(data[s : s+clen])
						s += clen
					} else if l-s == 2 && data[s] == 13 && data[s+1] == 10 {
						req.decodeQueryPost()
						return s + 2, req.Body.Bytes(), nil
					}

				}

			} else {
				if l-s < clen {
					return 0, nil, nil
				}
				req.Body.ResetBuf(data[s : s+clen])
				req.decodeQueryPost()
				return s + clen, req.Body.Bytes(), nil
			}
		} else {
			j := bytes.IndexByte(data[s:s+i], 58)
			req.Header[line[:j]] = line[j+2:]
		}

	}

	// not enough data
	return 0, nil, nil
}
func (req *Request) decodeQueryPost() {
	if req.query != "" {
		for _, str := range strings.Split(req.query, "&") {
			s := strings.Split(str, "=")
			if len(s) == 2 {
				k, err1 := url.QueryUnescape(s[0])
				v, err2 := url.QueryUnescape(s[1])
				if err1 == nil && err2 == nil {
					req.addquery(k, v)
				}
			}
		}
	}

	if strings.Contains(req.Header["Content-Type"], "application/x-www-form-urlencoded") {
		for _, str := range strings.Split(req.Body.String(), "&") {
			if i := strings.Index(str, "="); i > 0 {
				k, err1 := url.QueryUnescape(str[:i])
				v, err2 := url.QueryUnescape(str[i+1:])
				if err1 == nil && err2 == nil {
					req.addpost(k, v)
				}

			}
		}
	}
	if strings.Contains(req.Header["Content-Type"], "multipart/form-data") {
		if i := strings.Index(req.Header["Content-Type"], "boundary="); i > -1 {
			for _, str := range strings.Split(req.Body.String(), "--"+req.Header["Content-Type"][i+9:]+"\r\n") {
				i := strings.Index(str, "\r\n")

				if i > -1 {
					if strings.Contains(str[:i], "Content-Disposition: form-data;") {
						var key, value string
						if j := strings.Index(str[:i], `name="`); j > -1 {
							key, _ = url.QueryUnescape(str[j+6 : i-1])
						}
						if j := strings.Index(str[i+4:], "\r\n"); j > -1 {
							value, _ = url.QueryUnescape(str[i+4 : i+4+j])
						}
						if key != "" {
							req.addpost(key, value)
						}
					}

				}

			}
		}

	}
}
func (req *Request) addquery(name, value string) {
	for k, v := range req.queryS {
		if v.key == name {
			req.queryS[k].value = append(req.queryS[k].value, value)
			return
		}
	}
	oldlen := len(req.queryS)
	if oldlen+1 > cap(req.queryS) {
		req.queryS = append(req.queryS, httpQuery{
			key:   name,
			value: []string{value},
		})
	} else {
		req.queryS = req.queryS[:oldlen+1]
		req.queryS[oldlen].key = name
		req.queryS[oldlen].value = req.queryS[oldlen].value[:0]
		req.queryS[oldlen].value = append(req.queryS[oldlen].value, value)
	}
}
func (req *Request) addpost(name, value string) {
	for k, v := range req.postS {
		if v.key == name {
			req.postS[k].value = append(req.postS[k].value, value)
			return
		}
	}
	oldlen := len(req.postS)
	if oldlen+1 > cap(req.postS) {
		req.postS = append(req.postS, httpQuery{
			key:   name,
			value: []string{value},
		})
	} else {
		req.postS = req.postS[:oldlen+1]
		req.postS[oldlen].key = name
		req.postS[oldlen].value = req.postS[oldlen].value[:0]
		req.postS[oldlen].value = append(req.postS[oldlen].value, value)

	}
}

var (

	//http1origin  = []byte("Access-Control-Allow-Origin: " + config.Server.Origin + "\r\n")
	http1head200 = []byte("HTTP/1.1 200 OK\r\nserver: gnet by luyu6056\r\n")
	http1head206 = []byte("HTTP/1.1 206 Partial Content\r\nserver: gnet by luyu6056\r\n")
	http1head304 = []byte("HTTP/1.1 304 Not Modified\r\nserver: gnet by luyu6056\r\n")
	http1deflate = []byte("Content-encoding: deflate\r\n")
	http1gzip    = []byte("Content-encoding: gzip\r\n")
	http404b, _  = ioutil.ReadFile(static_patch + "/404.html")
	http1cache   = []byte("Cache-Control: max-age=86400\r\n")
	http1nocache = []byte("Cache-Control: no-store, no-cache, must-revalidate, max-age=0, s-maxage=0\r\nPragma: no-cache\r\n")
)

func (hs *Httpserver) StaticHandler() gnet.Action {
	hs.Out.Reset()
	etag := hs.Request.GetHeader("If-None-Match")
	filename := hs.Request.path
	if filename == "/" {
		filename = "/index.html"
	}

	isdeflate := strings.Contains(hs.Request.GetHeader("Accept-Encoding"), "deflate")
	var isgzip bool
	if !isdeflate {
		isgzip = strings.Contains(hs.Request.GetHeader("Accept-Encoding"), "gzip")
	}
	filename = static_patch + filename
	var f_cache *file_cache
	var f_cache_err error
	if cache, ok := static_cache.Load(filename); ok { //这个cache在http2那边
		f_cache = cache.(*file_cache)

		//有缓存，检查文件是否修改
		if !httpIswatcher && f_cache.etag != "" && atomic.CompareAndSwapUint32(&f_cache.check, 0, 1) {
			f_cache_err, f_cache = f_cache.Check(filename)
			time.AfterFunc(time.Second, func() { f_cache.check = 0 })
		}
	} else {
		if httpIswatcher {
			httpWatcher.Add(filename)
		}
		f_cache_err, f_cache = f_cache.Check(filename)

	}

	if f_cache_err == nil {
		if f_cache.etag == etag {
			hs.Out.Write(http1head304)
			hs.data.Reset(nil)
		} else if isdeflate && f_cache.iscompress { //deflate压缩资源
			hs.Out.Write(http1head200)
			hs.Out.WriteString("Content-Type: ")
			hs.Out.WriteString(f_cache.content_type)
			hs.Out.WriteString("\r\n")
			hs.Out.Write(http1deflate)
			hs.data.Reset(f_cache.deflatefile)
		} else if isgzip && f_cache.iscompress { //gzip可压缩资源
			hs.Out.Write(http1head200)
			hs.Out.WriteString("Content-Type: ")
			hs.Out.WriteString(f_cache.content_type)
			hs.Out.WriteString("\r\n")
			hs.Out.Write(http1gzip)
			g := gzippool.Get().(*gzip.Writer)
			buf := msgbufpool.Get().(*libraries.MsgBuffer)
			defer msgbufpool.Put(buf)
			defer gzippool.Put(g)
			buf.Reset()
			g.Reset(buf)
			g.Write(f_cache.file)
			g.Flush()
			hs.data.Reset(buf.Bytes())
		} else { //非压缩资源
			hs.Out.Write(http1head200)
			hs.Out.WriteString("Content-Type: ")
			hs.Out.WriteString(f_cache.content_type)
			hs.Out.WriteString("\r\n")
			hs.data.Reset(f_cache.file)
		}
		hs.Out.WriteString("Etag: ")
		hs.Out.WriteString(f_cache.etag)
		hs.Out.WriteString("\r\n")
		hs.httpsfinish(hs.data, hs.data.Len())
		return gnet.None
	} else {
		switch f_cache_err {
		case file_cache_err_NotFound:
			hs.Out404()
		case file_cache_file_TooLarge:
			f, err := os.Open(filename)
			if err != nil {
				hs.OutErr(err)
				return gnet.None
			}
			defer f.Close()
			fstat, err := f.Stat()
			if err != nil {
				hs.OutErr(err)
				return gnet.None
			}
			hs.RangeDownload(f, fstat.Size(), path.Base(filename))
		default:
			hs.OutErr(errors.New("Unknown Error"))
		}

	}
	return gnet.None
}

type HttpIoReader interface {
	Seek(int64, int) (int64, error)
	Read([]byte) (int, error)
}

func (hs *Httpserver) RangeDownload(b HttpIoReader, size int64, name string) {
	var range_start, range_end int
	if r := hs.Request.GetHeader("range"); strings.Index(r, "bytes=") == 0 {
		if e := strings.Index(r, "-"); e > 6 {
			range_start, _ = strconv.Atoi(r[6:e])
			range_end, _ = strconv.Atoi(r[e+1:])
		}
	}
	if range_start > 0 || range_end > 0 {
		hs.Out.Write(http1head206)
		if range_end == 0 {
			range_end = int(size)
		}
		if _, e := b.Seek(int64(range_start), 0); e != nil {
			hs.OutErr(e)
			return
		}
		hs.Out.WriteString("Content-Type: application/octet-stream\r\nAccept-Ranges: bytes\r\nContent-Range: bytes ")

		hs.Out.WriteString(strconv.Itoa(range_start))
		hs.Out.WriteString("-")
		hs.Out.WriteString(strconv.Itoa(range_end))
		hs.Out.WriteString("/")
		hs.Out.WriteString(strconv.Itoa(int(size)))
		hs.Out.WriteString("\r\n")
		hs.Out.WriteString(`Content-Disposition: attachment; filename*="utf8''` + url.QueryEscape(name) + `"` + "\r\n")
		hs.httpsfinish(b, range_end-range_start)
	} else {
		hs.Out.Write(http1head200)
		hs.Out.WriteString("Content-Type: application/octet-stream\r\n")
		hs.Out.WriteString(`Content-Disposition: attachment; filename*="utf8''` + url.QueryEscape(name) + `"` + "\r\n")
		hs.httpsfinish(b, int(size))
	}
}
func (hs *Httpserver) httpsfinish(b io.Reader, l int) {
	if hs.ishttps {
		hs.Out.WriteString("strict-transport-security: max-age=31536000; includeSubDomains\r\n")
	}
	for k, v := range hs.OutHeader {
		hs.Out.WriteString(k)
		hs.Out.WriteString(": ")
		hs.Out.WriteString(v)
		hs.Out.WriteString("\r\n")
	}

	for k, v := range hs.OutCookie {
		hs.Out.WriteString("Set-Cookie: ")
		hs.Out.WriteString(url.QueryEscape(k))
		hs.Out.WriteString("=")
		hs.Out.WriteString(url.QueryEscape(v.value))
		if v.max_age > 0 {
			hs.Out.WriteString("; Max-age=")
			hs.Out.WriteString(strconv.FormatUint(uint64(v.max_age), 10))
		}
		hs.Out.WriteString("; path=/")
		hs.Out.WriteString("\r\n")
	}

	hs.Out.WriteString("Connection: ")
	hs.Out.WriteString(hs.Request.Connection)

	hs.Out.WriteString("\r\nContent-Length: ")
	hs.Out.WriteString(strconv.Itoa(l))
	hs.Out.WriteString("\r\n\r\n")
	for msglen := l; msglen > 0; msglen = l {
		if msglen > http2initialMaxFrameSize*100-hs.Out.Len() { //切分为一个tls包
			msglen = http2initialMaxFrameSize*100 - hs.Out.Len()
		}
		if _, e := b.Read(hs.Out.Make(msglen)); e != nil {
			libraries.DebugLog("httpsfinish Read错误%v", e)
			hs.Out.Reset()
			hs.OutErr(e)
			hs.c.Close()
			return
		}
		if l > msglen {
			hs.c.FlushWrite(hs.Out.Bytes(), true)
		} else {
			hs.c.WriteNoCodec(hs.Out.Bytes())
		}

		hs.Out.Reset()
		l -= msglen
	}
	if l := hs.Out.Len(); l > 0 {
		hs.c.WriteNoCodec(hs.Out.Next(l))
	}

}
func (hs *Httpserver) Out404() {
	hs.Out.WriteString("HTTP/1.1 404 Not Found\r\nContent-Length: ")
	hs.Out.WriteString(strconv.Itoa(len(http404b)))
	hs.Out.WriteString("\r\n\r\n")
	hs.Out.Write(http404b)
	hs.c.AsyncWrite(hs.Out.Bytes())
}

var Errfunc func(i interface{}, err error) bool

func (hs *Httpserver) OutErr(err error) {
	if Errfunc != nil {
		if Errfunc(hs, err) {
			return
		}
	}
	hs.Out.WriteString("HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/html;charset=utf-8\nContent-Length: ")
	hs.Out.WriteString(strconv.Itoa(len(err.Error())))
	hs.Out.WriteString("\r\n\r\n")
	hs.Out.WriteString(err.Error())
	hs.c.AsyncWrite(hs.Out.Bytes())

}

func init() {

	if http404b == nil {
		http404b = []byte("404 not found")
	}
}
func (hs *Httpserver) Close() {
	hs.session = nil
	hs.c.Close()
}

var sessionID = uint64(rand.NewSource(time.Now().Unix()).Int63())

func (hs *Httpserver) Session() *cache.Hashvalue {
	if hs.session == nil {
		//检查sessionID
		var has bool
		sessionIdKey := hs.Cookie("sessionID")
		if sessionIdKey != "" {
			hs.session, has = cache.Has(sessionIdKey, "session")
		}
		//不存在则创建一个
		if !has {
			has = true
			//循环检查到一个没用过的sessionIdKey
			for has {
				b := make([]byte, 8)
				binary.LittleEndian.PutUint64(b, atomic.AddUint64(&sessionID, 1))
				sessionIdKey = strings.TrimRight(libraries.SHA256_URL_BASE64(strconv.FormatInt(time.Now().UnixNano(), 10)+string(b)), "=")
				_, has = cache.Has(sessionIdKey, "session")
			}
			hs.SetCookie("sessionID", sessionIdKey, 7*86400)
			hs.session = cache.Hget(sessionIdKey, "session")
			hs.session.Set("sessionID", sessionIdKey)
			hs.session.Expire(8 * 3600) //给个临时session
		}
	}
	return hs.session
}
func (hs *Httpserver) DelSession() {
	if hs.session != nil {
		hs.session.Hdel()
	}
}
func (hs *Httpserver) Body() []byte {
	return hs.Request.Body.Bytes()
}
func (hs *Httpserver) Method() string {
	return hs.Request.Method
}
func (hs *Httpserver) Header(name string) string {
	return hs.Request.Header[name]
}
func (hs *Httpserver) SetHeader(name, value string) {
	hs.OutHeader[name] = value
}
func (hs *Httpserver) SetCookie(name, value string, max_age uint32) {
	hs.OutCookie[name] = httpcookie{
		value:   value,
		max_age: max_age,
	}
}
func (hs *Httpserver) Redirect(url string) {
	hs.Out.Reset()
	hs.Out.WriteString("HTTP/1.1 302 OK\r\nserver: gnet by luyu6056\r\nCache-Control: Max-age=0\r\nContent-Type: text/html;charset=utf-8\r\nLocation: ")
	hs.Out.WriteString(url)
	hs.Out.WriteString("\r\n")
	hs.httpsfinish(nil, 0)
}
func (hs *Httpserver) Recovery() {
	hs.Out.Reset()
	for k := range hs.OutHeader {
		delete(hs.OutHeader, k)
	}
	for k := range hs.OutCookie {
		delete(hs.OutCookie, k)
	}
	for k := range hs.Request.Header {
		delete(hs.Request.Header, k)
	}
	hs.Request.Cookie = nil
	hs.session = nil
	hs.Request.queryS = hs.Request.queryS[:0]
	hs.Request.postS = hs.Request.postS[:0]
	hs.OutCode = 0
	hs.OutContentType = ""
	hs.Request.query = ""
}
func (hs *Httpserver) Cookie(name string) string {
	if cookieHead, ok := hs.Request.Header["Cookie"]; ok {
		for _, cookie := range strings.Split(cookieHead, "; ") {
			if i := strings.Index(cookie, "="); i > 0 && cookie[:i] == name {
				v, _ := url.QueryUnescape(cookie[i+1:])
				return v
			}
		}
	}
	return ""
}
func (hs *Httpserver) Path() string {

	return hs.Request.path
}
func (hs *Httpserver) Query(key string) string {
	for _, q := range hs.Request.queryS {
		if q.key == key {
			return q.value[0]
		}
	}
	return ""
}

func (hs *Httpserver) Post(key string) (value string) {
	for _, q := range hs.Request.postS {
		if q.key == key {
			return q.value[0]
		}
	}
	return
}
func (hs *Httpserver) PostSlice(key string) []string {
	for _, q := range hs.Request.postS {
		if q.key == key {
			return q.value
		}
	}
	return nil
}
func (hs *Httpserver) GetAllPost() (res map[string][]string) {
	res = make(map[string][]string, len(hs.Request.postS))
	for _, v := range hs.Request.postS {
		res[v.key] = v.value
	}
	return res
}
func (hs *Httpserver) GetAllQuery() (res map[string][]string) {
	res = make(map[string][]string, len(hs.Request.queryS))
	for _, v := range hs.Request.queryS {
		res[v.key] = v.value
	}
	return res
}
func (hs *Httpserver) AddQuery(name, value string) {
	hs.Request.addquery(name, value)
}

func (hs *Httpserver) SetCode(code int) {
	hs.OutCode = code
}
func (hs *Httpserver) SetContentType(ContentType string) {
	hs.OutContentType = ContentType
}

type httpCode int

func (code httpCode) String() string {
	return map[int]string{
		100: "100 Continue",
		101: "101 Switching Protocols",
		102: "102 Processing",
		200: "200 OK",
		201: "201 Created",
		202: "202 Accepted",
		203: "203 Non-Authoritative Information",
		204: "204 No Content",
		205: "205 Reset Content",
		206: "206 Partial Content",
		207: "207 Multi-Status",
		300: "300 Multiple Choices",
		301: "301 Moved Permanently",
		302: "302 Move Temporarily",
		303: "303 See Other",
		304: "304 Not Modified",
		305: "305 Use Proxy",
		306: "306 Switch Proxy",
		307: "307 Temporary Redirect",
		400: "400 Bad Request",
		401: "401 Unauthorized",
		402: "402 Payment Required",
		403: "403 Forbidden",
		404: "404 Not Found",
		405: "405 Method Not Allowed",
		406: "406 Not Acceptable",
		407: "407 Proxy Authentication Required",
		408: "408 Request Timeout",
		409: "409 Conflict",
		410: "410 Gone",
		411: "411 Length Required",
		412: "412 Precondition Failed",
		413: "413 Request Entity Too Large",
		414: "414 Request-URI Too Long",
		415: "415 Unsupported Media Type",
		416: "416 Requested Range Not Satisfiable",
		417: "417 Expectation Failed",
		418: "418 I'm a teapot",
		421: "421 Misdirected Request",
		422: "422 Unprocessable Entity",
		423: "423 Locked",
		424: "424 Failed Dependency",
		425: "425 Too Early",
		426: "426 Upgrade Required",
		449: "449 Retry With",
		451: "451 Unavailable For Legal Reasons",
		500: "500 Internal Server Error",
		501: "501 Not Implemented",
		502: "502 Bad Gateway",
		503: "503 Service Unavailable",
		504: "504 Gateway Timeout",
		505: "505 HTTP Version Not Supported",
		506: "506 Variant Also Negotiates",
		507: "507 Insufficient Storage",
		509: "509 Bandwidth Limit Exceeded",
		510: "510 Not Extended",
		600: "600 Unparseable Response Headers",
	}[int(code)]
}

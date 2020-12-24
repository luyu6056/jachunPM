package server

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
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"protocol"
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
	session *cache.Hashvalue
	//ClientFd [4]byte //用户的fd标识
	Request Request
	//IsServer bool
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
}
type httpcookie struct {
	value   string
	max_age uint32
}
type Request struct {
	Proto, Method string
	path, query   string
	RemoteAddr    string
	Connection    string
	Header        map[string]string
	Cookie        map[string]string
	Body          *libraries.MsgBuffer
	FormCache     url.Values
	PostForm      url.Values
	QueryCache    url.Values
	MultipartForm *multipart.Form
	Form          url.Values
	URL           *url.URL
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

	hs.data.Reset(b.Bytes())
	if b.Len() > 4096 {
		isdeflate := strings.Contains(hs.Request.GetHeader("Accept-Encoding"), "deflate")
		var isgzip bool
		if !isdeflate {
			isgzip = strings.Contains(hs.Request.GetHeader("Accept-Encoding"), "gzip")
		}
		if isdeflate { //deflate压缩资源
			hs.Out.Write(http1deflate)
			buf := msgbufpool.Get().(*libraries.MsgBuffer)
			defer msgbufpool.Put(buf)
			buf.Reset()
			w := CompressNoContextTakeover(buf, 6)
			w.Write(b.Bytes())
			w.Close()
			hs.data.Reset(buf.Bytes())
		} else if isgzip { //gzip可压缩资源
			hs.Out.Write(http1gzip)
			g := gzippool.Get().(*gzip.Writer)
			buf := msgbufpool.Get().(*libraries.MsgBuffer)
			defer msgbufpool.Put(buf)
			defer gzippool.Put(g)
			buf.Reset()
			g.Reset(buf)
			g.Write(b.Bytes())
			g.Flush()
			hs.data.Reset(buf.Bytes())
		}
	}
	hs.httpsfinish(hs.data, hs.data.Len())

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
	defer msgbufpool.Put(buf)
	buf.Reset()
	buf.WriteString(str)
	hs.httpsfinish(buf, buf.Len())
	buf.Reset()
}
func (hs *Httpserver) RemoteAddr() string {
	return hs.c.RemoteAddr().String()
}
func (hs *Httpserver) IP() (ip string) {

	if ip = hs.Request.GetHeader("X-Real-IP"); ip == "" {
		ip = hs.c.RemoteAddr().String()
	}
	ip, _ = libraries.Preg_replace(`:\d+$`, "", ip)
	return ip
}
func (hs *Httpserver) IsMobile() bool {
	return false
}

func (hs *Httpserver) UserAgent() string {
	return hs.Request.GetHeader("UserAgent")
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
		Write:      hs.c.AsyncWrite,
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
	hs.c.AsyncWrite(hs.Out.Bytes())
	time.AfterFunc(time.Second*10, func() {
		if ctx, ok := hs.c.Context().(*WSconn); !ok || ctx.Conn == nil || ctx.Conn.Session == nil {
			hs.c.Close()
		}
	})
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
			i++
			s = bytes.Index(data[i:], []byte{13, 10})
			if s > -1 {
				s += i
				req.Proto = sdata[i:s]
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

	for s += 2; s < l; s += i + 2 {
		i = bytes.Index(data[s:], []byte{13, 10})
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
						return s + 2, req.Body.Bytes(), nil
					}

				}

			} else {
				if l-s < clen {
					return 0, nil, nil
				}
				req.Body.ResetBuf(data[s : s+clen])
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
	var range_start, range_end int
	if r := hs.Request.GetHeader("range"); strings.Index(r, "bytes=") == 0 {
		if e := strings.Index(r, "-"); e > 6 {
			range_start, _ = strconv.Atoi(r[6:e])
			range_end, _ = strconv.Atoi(r[e+1:])
		}
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
			if range_start > 0 || range_end > 0 {
				hs.Out.Write(http1head206)
				if range_end == 0 {
					range_end = int(fstat.Size())
				}
				f.Seek(int64(range_start), 0)
				hs.Out.WriteString("Content-Type: application/octet-stream\r\nAccept-Ranges: bytes\r\nContent-Range: bytes ")
				hs.Out.WriteString(strconv.Itoa(range_start))
				hs.Out.WriteString("-")
				hs.Out.WriteString(strconv.Itoa(range_end))
				hs.Out.WriteString("/")
				hs.Out.WriteString(strconv.Itoa(int(fstat.Size())))
				hs.httpsfinish(f, range_end-range_start)
			} else {
				hs.Out.Write(http1head200)
				hs.Out.WriteString("Content-Type: application/octet-stream\r\n")
				hs.httpsfinish(f, int(fstat.Size()))
			}
		default:
			hs.OutErr(errors.New("Unknown Error"))
		}

	}
	return gnet.None
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
		if msglen > http2initialMaxFrameSize-hs.Out.Len() { //切分为一个tls包
			msglen = http2initialMaxFrameSize - hs.Out.Len()
		}
		b.Read(hs.Out.Make(msglen))
		hs.c.AsyncWrite(hs.Out.Bytes())
		hs.Out.Reset()
		l -= msglen
	}
	if l := hs.Out.Len(); l > 0 {
		hs.c.AsyncWrite(hs.Out.Next(l))
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
	hs.Out.WriteString("HTTP/1.1 500 Internal Server Error\r\nContent-Length: ")
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
			hs.SetCookie("sessionID", sessionIdKey, protocol.SessionTempExpires)
			hs.session = cache.Hget(sessionIdKey, "session")
			hs.session.Expire(3600) //给个临时session
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
	hs.Out.WriteString("\r\n\r\n")
	hs.c.AsyncWrite(hs.Out.Bytes())
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
	hs.Request.FormCache = nil
	hs.Request.QueryCache = nil
	hs.Request.PostForm = nil
	hs.Request.Form = nil
	hs.OutCode = 0
	hs.OutContentType = ""
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
	hs.getQueryCache()
	if values, ok := hs.Request.QueryCache[key]; ok && len(values) > 0 {
		return values[0]
	}
	return ""
}
func (hs *Httpserver) getQueryCache() {
	if hs.Request.QueryCache == nil {
		hs.Request.QueryCache = make(url.Values)
		hs.Request.QueryCache, _ = url.ParseQuery(hs.Request.query)
	}
}
func (c *Httpserver) Post(key string) (value string) {
	value = c.PostForm(key)
	if value == "" {
		value = c.Query(key)
	}
	return
}
func (hs *Httpserver) GetAllPost() (res map[string][]string) {
	hs.Request.getFormCache()
	res = make(map[string][]string, len(hs.Request.FormCache))
	for k, v := range hs.Request.FormCache {
		res[k] = v
	}
	return res
}
func (hs *Httpserver) GetAllQuery() (res map[string][]string) {
	hs.getQueryCache()
	res = make(map[string][]string, len(hs.Request.QueryCache))
	for k, v := range hs.Request.QueryCache {
		res[k] = v
	}
	return res
}
func (hs *Httpserver) AddQuery(name, value string) {
	hs.getQueryCache()
	for k, _ := range hs.Request.QueryCache {
		if k == name {
			hs.Request.QueryCache[k] = []string{value}
			return
		}
	}
	hs.Request.QueryCache[name] = []string{value}
}
func (hs *Httpserver) SetCode(code int) {
	hs.OutCode = code
}
func (hs *Httpserver) SetContentType(ContentType string) {
	hs.OutContentType = ContentType
}
func (hs *Httpserver) PostForm(key string) string {
	hs.Request.getFormCache()
	if values := hs.Request.FormCache[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}
func (req *Request) getFormCache() {
	if req.FormCache == nil {
		if err := req.ParseForm(); err != nil {
			if err != http.ErrNotMultipart {
				libraries.DebugLog("error on parse multipart form array: %v", err)
			}
		}
		req.FormCache = req.PostForm
	}
}

var multipartByReader = &multipart.Form{
	Value: make(map[string][]string),
	File:  make(map[string][]*multipart.FileHeader),
}

func (r *Request) ParseMultipartForm(maxMemory int64) error {
	if r.MultipartForm == multipartByReader {
		return errors.New("http: multipart handled by MultipartReader")
	}
	if r.Form == nil {
		err := r.ParseForm()
		if err != nil {
			return err
		}
	}
	if r.MultipartForm != nil {
		return nil
	}

	mr, err := r.multipartReader(false)
	if err != nil {
		return err
	}

	f, err := mr.ReadForm(maxMemory)
	if err != nil {
		return err
	}

	if r.PostForm == nil {
		r.PostForm = make(url.Values)
	}
	for k, v := range f.Value {
		r.Form[k] = append(r.Form[k], v...)
		// r.PostForm should also be populated. See Issue 9305.
		r.PostForm[k] = append(r.PostForm[k], v...)
	}

	r.MultipartForm = f

	return nil
}
func (r *Request) ParseForm() error {
	var err error
	if r.PostForm == nil {
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
			r.PostForm, err = parsePostForm(r)
		}
		if r.PostForm == nil {
			r.PostForm = make(url.Values)
		}
	}
	if r.Form == nil {
		if len(r.PostForm) > 0 {
			r.Form = make(url.Values)
			copyValues(r.Form, r.PostForm)
		}
		var newValues url.Values
		if r.URL != nil {
			var e error
			newValues, e = url.ParseQuery(r.URL.RawQuery)
			if err == nil {
				err = e
			}
		}
		if newValues == nil {
			newValues = make(url.Values)
		}
		if r.Form == nil {
			r.Form = newValues
		} else {
			copyValues(r.Form, newValues)
		}
	}
	return err
}
func (r *Request) multipartReader(allowMixed bool) (*multipart.Reader, error) {
	v := r.GetHeader("Content-Type")
	if v == "" {
		return nil, ErrNotMultipart
	}
	d, params, err := mime.ParseMediaType(v)
	if err != nil || !(d == "multipart/form-data" || allowMixed && d == "multipart/mixed") {
		return nil, ErrNotMultipart
	}
	boundary, ok := params["boundary"]
	if !ok {
		return nil, ErrMissingBoundary
	}
	return multipart.NewReader(r.Body, boundary), nil
}
func parsePostForm(r *Request) (vs url.Values, err error) {
	if r.Body == nil {
		err = errors.New("missing form body")
		return
	}
	ct := r.GetHeader("Content-Type")
	// RFC 7231, section 3.1.1.5 - empty type
	//   MAY be treated as application/octet-stream
	if ct == "" {
		ct = "application/octet-stream"
	}
	ct, _, err = mime.ParseMediaType(ct)
	switch {
	case ct == "application/x-www-form-urlencoded":
		var e error
		vs, e = url.ParseQuery(r.Body.String())
		if err == nil {
			err = e
		}
	case ct == "multipart/form-data":
		// handled by ParseMultipartForm (which is calling us, or should be)
		// TODO(bradfitz): there are too many possible
		// orders to call too many functions here.
		// Clean this up and write more tests.
		// request_test.go contains the start of this,
		// in TestParseMultipartFormOrder and others.
	}
	return
}
func copyValues(dst, src url.Values) {
	for k, vs := range src {
		for _, value := range vs {
			dst.Add(k, value)
		}
	}
}

type ProtocolError struct {
	ErrorString string
}

func (pe *ProtocolError) Error() string { return pe.ErrorString }

var (
	// ErrNotSupported is returned by the Push method of Pusher
	// implementations to indicate that HTTP/2 Push support is not
	// available.
	ErrNotSupported = &ProtocolError{"feature not supported"}

	// Deprecated: ErrUnexpectedTrailer is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrUnexpectedTrailer = &ProtocolError{"trailer header without chunked transfer encoding"}

	// ErrMissingBoundary is returned by Request.MultipartReader when the
	// request's Content-Type does not include a "boundary" parameter.
	ErrMissingBoundary = &ProtocolError{"no multipart boundary param in Content-Type"}

	// ErrNotMultipart is returned by Request.MultipartReader when the
	// request's Content-Type is not multipart/form-data.
	ErrNotMultipart = &ProtocolError{"request Content-Type isn't multipart/form-data"}

	// Deprecated: ErrHeaderTooLong is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrHeaderTooLong = &ProtocolError{"header too long"}

	// Deprecated: ErrShortBody is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrShortBody = &ProtocolError{"entity body too short"}

	// Deprecated: ErrMissingContentLength is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrMissingContentLength = &ProtocolError{"missing ContentLength in HEAD response"}
)

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

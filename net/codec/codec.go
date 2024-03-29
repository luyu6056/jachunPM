package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"golang.org/x/net/http2/hpack"

	"github.com/luyu6056/gnet"
	"github.com/luyu6056/tls"
)

//所有的Encode在上一级调用进行编码，以方便循环pool回收利用
var (
	http2clientPreface = []byte(http2ClientPreface)
)

type Tlscodec struct {
	Config *tls.Config
}

func (code *Tlscodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {

	return buf, nil
}

func (code *Tlscodec) Decode(c gnet.Conn) (data []byte, err error) {
	if c.BufferLength() > 0 {
		data = c.Read()

		switch svr := c.Context().(type) {
		case *WSconn:
			msgtype, data, err := svr.ReadMessage(data)
			c.ShiftN(svr.ReadLength)
			if msgtype == BinaryMessage {
				return data, nil
			}
			return nil, err
		case *Http2server:
			for data = c.Read(); len(data) > http2headerlength; data = c.Read() {
				svr.InHeader.Length = int(data[0])<<16 | int(data[1])<<8 | int(data[2])
				if svr.InHeader.Length > http2defaultMaxReadFrameSize {

					return nil, svr.connError(http2ErrCodeFrameSize)
				}
				if len(data) < svr.InHeader.Length+http2headerlength {
					return nil, nil
				}
				c.ShiftN(svr.InHeader.Length + http2headerlength)
				svr.InHeader.Type = data[3]
				if svr.InHeader.Type > 9 {
					return nil, svr.connError(http2ErrCodeProtocol)
				}
				svr.InHeader.Flags = data[4]
				svr.InHeader.StreamID = binary.BigEndian.Uint32(data[5:]) & (1<<31 - 1)

				//帧率控制
				if fps := atomic.AddUint32(&svr.fps, 1); fps == 1 {
					time.AfterFunc(time.Second, func() { svr.fps = 0 })
				} else if fps > http2fpslimit {
					return nil, svr.connError(http2ErrCodeFlowControl)
				}
				switch svr.InHeader.Type {
				case http2FrameData:
					if svr.InHeader.StreamID == 0 || int(svr.InHeader.StreamID) > len(svr.Streams) {

						return nil, svr.connError(http2ErrCodeProtocol)
					}
					stream := svr.Streams[svr.InHeader.StreamID]
					if stream == nil {

						return nil, svr.connError(http2ErrCodeStreamClosed)
					}
					pos := http2headerlength
					padding := 0
					if svr.InHeader.Flags&http2FlagDataPadded == http2FlagDataPadded {
						padding = int(data[pos])
						//DebugLog("设置了padded", padding)
						pos++
					}
					stream.In.Write(data[pos : svr.InHeader.Length+http2headerlength-padding])

					stream.IN_WINDOW_SIZE -= int32(svr.InHeader.Length)
					svr.Streams[0].IN_WINDOW_SIZE -= int32(svr.InHeader.Length)
					if stream.IN_WINDOW_SIZE < http2initialWindowSize*10 {
						http2writeWindow_Update{size: http2initialWindowSize * 10, streamId: stream.Id}.writeFrame(svr.Streams[0])
						svr.c.AsyncWrite(svr.Streams[0].Out.Next(svr.Streams[0].Out.Len()))
						stream.IN_WINDOW_SIZE += http2initialWindowSize * 10
					}
					if svr.Streams[0].IN_WINDOW_SIZE < http2initialWindowSize {
						http2writeWindow_Update{size: http2serverWindowSize, streamId: 0}.writeFrame(svr.Streams[0])
						svr.c.AsyncWrite(svr.Streams[0].Out.Next(svr.Streams[0].Out.Len()))
						svr.Streams[0].IN_WINDOW_SIZE += http2serverWindowSize
					}
					if svr.InHeader.Flags&http2FlagDataEndStream == http2FlagDataEndStream {
						stream.close |= 1
						//处理stream数据前清空缓冲数据
						if l := svr.Streams[0].Out.Len(); l > 0 {
							svr.c.AsyncWrite(svr.Streams[0].Out.Next(l))
						}
						svr.WorkStream = stream
						stream.headerbuf.Reset()
						stream.henc = hpack.NewEncoder(&stream.headerbuf)
						//解析post请求
						if strings.Contains(stream.content_type, "application/x-www-form-urlencoded") {
							for _, str := range strings.Split(stream.In.String(), "&") {
								if i := strings.Index(str, "="); i > 0 {
									k, err1 := url.QueryUnescape(str[:i])
									v, err2 := url.QueryUnescape(str[i+1:])
									if err1 == nil && err2 == nil {
										stream.post[k] = append(stream.post[k], v)
									}

								}
							}
						}
						if strings.Contains(stream.content_type, "multipart/form-data") {
							if i := strings.Index(stream.content_type, "boundary="); i > -1 {
								for _, str := range strings.Split(stream.In.String(), "--"+stream.content_type[i+9:]+"\r\n") {
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
												stream.post[key] = append(stream.post[key], value)
											}
										}
									}
								}
							}
						}
						return stream.In.Bytes(), nil
					}
				case http2FrameHeaders:
					if svr.InHeader.StreamID == 0 || int(svr.InHeader.StreamID)-2 > len(svr.Streams) {
						return nil, svr.connError(http2ErrCodeProtocol)
					}
					if int(svr.InHeader.StreamID) > len(svr.Streams) {
						svr.Streams = append(svr.Streams, make([]*Http2stream, len(svr.Streams))...)
					}
					stream := svr.Streams[svr.InHeader.StreamID]
					if stream == nil {
						stream = stream_pool.Get().(*Http2stream)
						stream.In.Reset()
						stream.Out.Reset()
						stream.Id = svr.InHeader.StreamID
						svr.Streams[stream.Id] = stream
						stream.close = 0
						stream.svr = svr
						stream.query = make(map[string][]string)
						stream.post = make(map[string][]string)
						stream.cookie = make(map[string]string)
						stream.session = nil
						stream.method = ""
						stream.path = ""
						stream.outCode = 200
						stream.OutContentType = ""
						stream.OutHeader = map[string]string{}
						stream.OutCookie = map[string]httpcookie{}
						stream.content_type = ""
						stream.accept_encoding = ""
					}
					svr.last_stream_id = stream.Id
					pos := http2headerlength
					padding := 0
					if svr.InHeader.Flags&http2FlagHeadersPadded == http2FlagHeadersPadded {
						padding = int(data[pos])
						//DebugLog("设置了padded", padding)
						pos++
					}
					if svr.InHeader.Flags&http2FlagHeadersPriority == http2FlagHeadersPriority {
						//DebugLog("设置了Priority", buf[pos:pos+5]) 暂不处理
						pos += 5
					}
					if svr.InHeader.Flags&http2FlagHeadersEndHeaders == http2FlagHeadersEndHeaders {
						//DebugLog("设置了endheader")
					} else {
						DebugLog("有continue")
					}
					stream.OUT_WINDOW_SIZE = svr.Setting.SETTINGS_INITIAL_WINDOW_SIZE
					stream.IN_WINDOW_SIZE = http2initialWindowSize
					svr.ReadMetaHeaders.SetMaxStringLength(svr.Setting.MAX_HEADER_LIST_SIZE)
					stream.Headers, err = svr.ReadMetaHeaders.DecodeFull(data[pos : svr.InHeader.Length+http2headerlength-padding])
					if err != nil {
						return nil, svr.connError(http2ErrCodeCompression)
					}

					//先解析一波header请求
					for _, head := range stream.Headers {
						switch head.Name {
						case ":path":
							stream.path, stream.uri = head.Value, head.Value

							if i := strings.Index(head.Value, "?"); i > 0 {
								stream.path = head.Value[:i]
								for _, str := range strings.Split(head.Value[i+1:], "&") {
									s := strings.Split(str, "=")
									if len(s) == 2 {
										k, err1 := url.QueryUnescape(s[0])
										v, err2 := url.QueryUnescape(s[1])
										if err1 == nil && err2 == nil {
											stream.query[k] = append(stream.query[k], v)
										}
									}
								}
							}

						case "cookie":
							for _, str := range strings.Split(head.Value, ";") {
								if i := strings.Index(str, "="); i > 0 {
									stream.cookie[strings.TrimLeft(str[:i], " ")], _ = url.QueryUnescape(str[i+1:])
								}
							}
						case ":method":
							stream.method = head.Value
						case "content-type":
							stream.content_type = head.Value
						case "accept-encoding":
							stream.accept_encoding = head.Value
						case "referer":
							stream.referer = head.Value
						default:

						}
					}

					if svr.InHeader.Flags&http2FlagHeadersEndStream == http2FlagHeadersEndStream {
						stream.close |= 1
						//处理stream数据前清空缓冲数据
						if l := svr.Streams[0].Out.Len(); l > 0 {
							svr.c.AsyncWrite(svr.Streams[0].Out.Next(l))
						}
						svr.WorkStream = stream
						stream.headerbuf.Reset()
						stream.henc = hpack.NewEncoder(&stream.headerbuf)

						return stream.In.Bytes(), nil
					}
				case http2FrameSettings:
					if svr.InHeader.Flags&http2FlagSettingsAck == http2FlagSettingsAck {
						if svr.InHeader.Length > 0 {
							// When this (ACK 0x1) bit is set, the payload of the
							// SETTINGS frame MUST be empty. Receipt of a
							// SETTINGS frame with the ACK flag set and a length
							// field value other than 0 MUST be treated as a
							// connection error (Section 5.4.1) of type
							// FRAME_SIZE_ERROR.
							return nil, svr.connError(http2ErrCodeFrameSize)
						}
						continue
					}
					if svr.InHeader.StreamID != 0 {
						// SETTINGS frames always apply to a connection,
						// never a single stream. The stream identifier for a
						// SETTINGS frame MUST be zero (0x0).  If an endpoint
						// receives a SETTINGS frame whose stream identifier
						// field is anything other than 0x0, the endpoint MUST
						// respond with a connection error (Section 5.4.1) of
						// type PROTOCOL_ERROR.
						return nil, svr.connError(http2ErrCodeProtocol)
					}
					if svr.InHeader.Length%6 != 0 {
						// Expecting even number of 6 byte settings.
						return nil, svr.connError(http2ErrCodeFrameSize)
					}
					for n := http2headerlength; n < svr.InHeader.Length+http2headerlength; n += 6 {
						value := int(data[n+2])<<24 | int(data[n+3])<<16 | int(data[n+4])<<8 | int(data[n+5])
						switch int(data[n])<<8 | int(data[n+1]) {
						case http2SettingHeaderTableSize:
							svr.Setting.HEADER_TABLE_SIZE = value
						case http2SettingEnablePush:
							switch value {
							case 0:
								svr.Setting.ENABLE_PUSH = false
							case 1:
								svr.Setting.ENABLE_PUSH = true
							default:
								return nil, svr.connError(http2ErrCodeProtocol)
							}
						case http2SettingMaxConcurrentStreams:
							svr.SendPool.Tune(value)
						case http2SettingInitialWindowSize:
							if value > (1<<31)-1 {
								return nil, svr.connError(http2ErrCodeFlowControl)
							}

							svr.Setting.SETTINGS_INITIAL_WINDOW_SIZE = int32(value)
							//h2s.OUT_WINDOW_SIZE = int32(value)
						case http2SettingMaxFrameSize:
							if value < 2^24-1 || value > 16777215 { //协议允许的范围
								return nil, svr.connError(http2ErrCodeProtocol)
							}
							svr.Setting.MAX_FRAME_SIZE = value
							if svr.Setting.MAX_FRAME_SIZE > http2initialMaxFrameSize-http2headerlength { //主动设置不大于一个tls包，避免tls层进行分包
								svr.Setting.MAX_FRAME_SIZE = http2initialMaxFrameSize - http2headerlength
							}
						case http2SettingMaxHeaderListSize:
							svr.Setting.MAX_HEADER_LIST_SIZE = value
						}
					}

					//libraries.Log("%+v", h2s.Setting)
					http2writeSettingsAck{}.writeFrame(svr.Streams[0])
				case http2FrameWindowUpdate:
					b := data[http2headerlength : svr.InHeader.Length+http2headerlength]
					value := int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3])
					if value < 0 {
						return nil, svr.connError(http2ErrCodeFlowControl)
					}

					if stream := svr.Streams[svr.InHeader.StreamID]; stream != nil {
						old := atomic.AddInt32(&stream.OUT_WINDOW_SIZE, value) - value
						if old <= 0 {
							svr.ReadPool.Submit(func() {
								select {
								case stream.sendch <- http2streamflagadd: //等待放行
								case <-time.After(http2WindowsSizeWaitTimeout):
									return
								}
							})
						}
					}

				case http2FramePriority:
				case http2FrameRSTStream: //目前未针对stream进行清理，可以循环使用
				case http2FramePing:
					http2writePing{}.writeFrame(svr.Streams[0])
				case http2FrameGoAway:
					last_stream_id := uint32(data[http2headerlength])<<24 | uint32(data[http2headerlength+1])<<16 | uint32(data[http2headerlength+2])<<8 | uint32(data[http2headerlength+3])
					err_code := uint32(data[http2headerlength+4])<<24 | uint32(data[http2headerlength+5])<<16 | uint32(data[http2headerlength+6])<<8 | uint32(data[http2headerlength+7])
					if err_code > 0 {
						errmsg := fmt.Sprintf("收到GoAway错误，流: %d, 错误码: %d %s,补充信息: %s", last_stream_id, err_code, http2errCodeName[http2ErrCode(err_code)], data[http2headerlength+8:http2headerlength+svr.InHeader.Length])

						return nil, errors.New(errmsg)
					} else {
						//DebugLog("关闭")
					}
					return nil, io.EOF

				default:
					DebugLog("未处理类型%v", svr.InHeader.Type)
				}
			}
			//除了header其他类型合并到一起发送，通常一个tls帧能发送完
			if l := svr.Streams[0].Out.Len(); l > 0 {
				svr.c.AsyncWrite(svr.Streams[0].Out.Next(l))
			}
			return nil, nil
		case *Httpserver:

			shift, d, err := svr.WorkRequest.Parsereq(data)
			c.ShiftN(shift)
			return d, err

		case nil:
			if len(data) < 24 {
				return nil, nil
			}

			switch {
			case bytes.Equal(data[:24], http2clientPreface):
				http2 := NewH2Conn(c)
				c.SetContext(http2)
				c.ShiftN(24)
				return code.Decode(c)
			default:
				c.SetContext(NewHttpServer(c, code.Config != nil))
				return code.Decode(c)
			}
		}
	}
	return nil, err
}

package protocol

import (
	"errors"
	"io/ioutil"

	"github.com/klauspost/compress/zstd"
	"github.com/luyu6056/gnet"
)

const ZstdLevel = zstd.SpeedDefault

var ZstdDict, _ = ioutil.ReadFile("../1.dict")

type RpcCodec struct {
}

var ErrRpcContext = errors.New("错误的rpcContext")

func (code RpcCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	return buf, nil
}

func (code RpcCodec) Decode(c gnet.Conn) (data []byte, err error) {
	if c.BufferLength() > 4 {
		data = c.Read()
		msglen := int(data[0]) | int(data[1])<<8 | int(data[2])<<16 | int(data[3]&127)<<24
		if len(data) < msglen+4 { //消息长度不够
			return nil, nil
		}
		c.ShiftN(msglen + 4)
		//解压缩
		if data[3]>>7 == 1 {

			if ctx, ok := c.Context().(Rpcdecompress); ok {
				return ctx.Decompress(data[4 : msglen+4]), nil
			}
			return nil, ErrRpcContext
		}

		return data[4 : msglen+4], nil
	}
	return nil, nil
}

type Rpcdecompress interface {
	Decompress(in []byte) (out []byte)
}

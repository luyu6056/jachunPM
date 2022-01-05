package protocol

import (
	"errors"

	"github.com/luyu6056/gnet"
)

type RpcCodec struct {
}

var ErrRpcContext = errors.New("错误的rpcContext")

func (code RpcCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	return buf, nil
}

func (code RpcCodec) Decode(c gnet.Conn) (data []byte, err error) {
	if c.BufferLength() > 5 {
		data = c.Read()
		msglen := int(data[0]) | int(data[1])<<8 | int(data[2])<<16 | int(data[3])<<24
		if len(data) < msglen+5 { //消息长度不够
			return nil, nil
		}
		c.ShiftN(msglen + 5)
		//解压缩
		if data[4]>>7 == 1 {
			ctx := c.Context().(Rpcdecompress)
			return ctx.Decompress(data[4 : msglen+5]), nil

			return nil, ErrRpcContext
		}
		return data[4 : msglen+5], nil
	}
	return nil, nil
}

type Rpcdecompress interface {
	Decompress(in []byte) (out []byte)
}

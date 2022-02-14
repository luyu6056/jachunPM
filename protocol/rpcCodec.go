package protocol

import (
	"errors"
	"io/ioutil"

	"github.com/luyu6056/gnet"
	"github.com/valyala/gozstd"
)

var zstdDict, _ = ioutil.ReadFile("../1.dict")
var cdict, _ = gozstd.NewCDictLevel(zstdDict, 1)
var ddict, _ = gozstd.NewDDict(zstdDict)

type RpcCodec struct {
}

var ErrRpcContext = errors.New("错误的rpcContext")

func (code RpcCodec) Encode(c gnet.Conn, data []byte) ([]byte, error) {
	if ctx, ok := c.Context().(RpcCompress); ok {
		buf := ctx.EncodeBuf()
		*buf = (*buf)[:4]
		*buf = gozstd.CompressDict(*buf, data, cdict)
		msglen := len(*buf) - 4
		(*buf)[0], (*buf)[1], (*buf)[2], (*buf)[3] = byte(msglen), byte(msglen>>8), byte(msglen>>16), byte(msglen>>24)
		return *buf, nil
	}
	return nil, ErrRpcContext
}

func (code RpcCodec) Decode(c gnet.Conn) (data []byte, err error) {
	if c.BufferLength() > 4 {
		data = c.Read()
		msglen := uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
		if len(data) < 4+int(msglen) { //消息长度不够
			return nil, nil
		}

		c.ShiftN(4 + int(msglen))
		//解压缩
		if ctx, ok := c.Context().(RpcCompress); ok {
			buf := ctx.DecodeBuf()
			*buf = (*buf)[:0]
			*buf, err = gozstd.DecompressDict(*buf, data[4:4+msglen], ddict)
			return *buf, err
		}
		return nil, ErrRpcContext
	}
	return nil, nil
}

type RpcCompress interface {
	EncodeBuf() (buf *[]byte)
	DecodeBuf() (buf *[]byte)
}

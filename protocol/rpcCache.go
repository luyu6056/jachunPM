package protocol

import (
	"errors"
	"libraries"
	"reflect"
	"time"
)

const (
	cacheTimeOut = time.Second * 5
)

var (
	cacheServerNotStart = errors.New("cache服务未启动")
)

type RpcCache struct {
	Svr MsgServer
}

func (c *RpcCache) Get(name, path string) (b []byte, err error) {
	if c.Svr == nil {
		return nil, cacheServerNotStart
	}
	data := GET_MSG_HOST_CACHE_GET()
	data.Path = path
	data.Name = name
	defer data.Put()
	var resdata *MSG_HOST_CACHE_GET_result
	err = c.Svr.SendMsgWaitResult(0, 0, 0, 0, data, &resdata)
	if err != nil {
		return nil, err
	}
	//libraries.DebugLog("Path:%s Name:%s Value:%s", path, name, libraries.MD5_B(resdata.Value))
	return resdata.Value, nil

}
func (c *RpcCache) GetPath(path string) (value [][]byte, err error) {

	if c.Svr == nil {
		return nil, cacheServerNotStart
	}
	data := GET_MSG_HOST_CACHE_GETPATH()
	data.Path = path
	defer data.Put()
	var resdata *MSG_HOST_CACHE_GETPATH_result
	err = c.Svr.SendMsgWaitResult(0, 0, 0, 0, data, &resdata)
	if err != nil {
		return nil, err
	}
	return resdata.Value, nil
}
func (c *RpcCache) Set(name, path string, value []byte, expire int64) error {
	//libraries.DebugLog("Path:%s Name:%s Value:%s", path, name, libraries.MD5_B(value))
	if c.Svr == nil {
		return cacheServerNotStart
	}
	data := GET_MSG_HOST_CACHE_SET()
	data.Path = path
	data.Name = name
	data.Value = make([]byte, len(value))
	copy(data.Value, value)
	data.Expire = expire
	c.Svr.SendMsg(0, 0, 0, 0, 0, data)
	data.Put()

	return nil
}
func (c *RpcCache) Del(name, path string) error {
	if c.Svr == nil {
		return cacheServerNotStart
	}
	data := GET_MSG_HOST_CACHE_DEL()
	data.Path = path
	data.Name = name
	c.Svr.SendMsg(0, 0, 0, 0, 0, data)
	data.Put()
	return nil
}
func (c *RpcCache) DelPath(path string) error {
	if c.Svr == nil {
		return cacheServerNotStart
	}
	data := GET_MSG_HOST_CACHE_DelPath()
	data.Path = path
	c.Svr.SendMsg(0, 0, 0, 0, 0, data)
	data.Put()
	return nil
}
func HandleCache(in *Msg) {
	switch data := in.Data.(type) {

	default:
		libraries.ReleaseLog("local%d remote%d cache未设置消息%s处理", in.Local, in.GetRemoteID(), reflect.TypeOf(data).Elem().Name())
	}
}

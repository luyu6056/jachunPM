package main

//注册服务
type MSG_COMMON_regServer struct {
	No     uint8
	IpPort string
	Time   int64
	Token  string
	Window int32
}

type MSG_COMMON_regServer_result struct {
	Id uint8
}

//中心服激活ticker
type MSG_COMMON_StartTicker struct{}

type MSG_COMMON_PING struct{}
type MSG_COMMON_PONG struct{}

//增加窗口
type MSG_COMMON_WINDOW_UPDATE struct {
	Add int32
}

type MSG_COMMON_CACHE_GET struct {
	QueryID uint32 //请求id
	Path    string
	Name    string
}
type MSG_COMMON_CACHE_GET_result struct {
	QueryResultID uint32 //返回请求id
	Value         []byte
}
type MSG_COMMON_CACHE_GETPATH struct {
	QueryID uint32 //请求id
	Path    string
}
type MSG_COMMON_CACHE_GETPATH_result struct {
	QueryResultID uint32 //返回请求id
	Value         [][]byte
}

type MSG_COMMON_CACHE_SET struct {
	Path   string
	Name   string
	Value  []byte
	Expire int64
}
type MSG_COMMON_CACHE_DEL struct {
	Path string
	Name string
}
type MSG_COMMON_CACHE_DelPath struct {
	Path string
}

type MSG_COMMON_GET_Msgno struct {
	QueryID uint32 //请求id
}
type MSG_COMMON_GET_Msgno_result struct {
	QueryResultID uint32 //返回请求id
	Msgno         uint32
}

type MSG_COMMON_QueryErr struct {
	QueryResultID uint32 //返回请求id
	Err           string
	Stack         []byte
}
type MSG_COMMON_ResetWindow struct {
	Window int32
}

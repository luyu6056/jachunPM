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
type MSG_COMMON_StartTicker struct {
}

type MSG_COMMON_PING struct {
}

type MSG_COMMON_PONG struct {
}

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

//小文件上传下载
type MSG_FILE_upload struct {
	QueryID    uint32 //请求id
	Name       string
	Data       []byte
	AddBy      int32
	ObjectType string
	ObjectID   int32
	Code       string //如果是项目的文件，需要“产品代号/项目代号”，会保存在对应的文件夹内
	Type       string //bindingFile sourceFile feedbackFile finalFile processFile
}
type MSG_FILE_upload_result struct {
	QueryResultID uint32 //返回请求id
	FileID        int64
}
type MSG_FILE_getByID struct {
	QueryID uint32 //请求id
	FileID  int64
	NoData  bool //true不反回date
}
type MSG_FILE_getByID_result struct {
	QueryResultID uint32 //返回请求id
	Name          string
	Ext           string
	Data          []byte
	Type          string
}
type MSG_FILE_updateMapByWhere struct {
	QueryID uint32 //请求id
	Where   map[string]interface{}
	Update  map[string]interface{}
}

type MSG_FILE_DeleteByID struct { //
	QueryID uint32
	FileID  int64
}
type MSG_COMMON_BeginTransaction struct {
	QueryID       uint32
	TransactionNo uint32
}
type MSG_COMMON_BeginTransaction_result struct {
	QueryResultID uint32
	TransactionNo uint32
}
type MSG_COMMON_Transaction_Commit struct {
	QueryID uint32
	No      uint32
}
type MSG_COMMON_Transaction_RollBack struct {
	QueryID uint32
	No      uint32
}
type MSG_COMMON_Transaction_Check struct {
	QueryID uint32
	No      uint32
}

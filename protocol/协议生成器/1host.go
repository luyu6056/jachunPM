package main

import "time"

//注册服务
type MSG_HOST_regServer struct {
	No     uint8
	IpPort string
	Time   int64
	Token  string
	Window int32
}

type MSG_HOST_regServer_result struct {
	Id uint8
}

//中心服激活ticker
type MSG_HOST_StartTicker struct {
}

type MSG_HOST_PING struct {
	Rand int
}

type MSG_HOST_PONG struct {
	Rand int
}

//增加窗口
type MSG_HOST_WINDOW_UPDATE struct {
	Add int32
}

type MSG_HOST_CACHE_GET struct {
	Path string
	Name string
}
type MSG_HOST_CACHE_GET_result struct {
	Value []byte
}
type MSG_HOST_CACHE_GETPATH struct {
	Path  string
	Names []string
}
type MSG_HOST_CACHE_GETPATH_result struct {
	Value [][]byte
}

type MSG_HOST_CACHE_SET struct {
	Path   string
	Name   string
	Value  []byte
	Expire int64
}
type MSG_HOST_CACHE_DEL struct {
	Path string
	Name string
}
type MSG_HOST_CACHE_DelPath struct {
	Path string
}

type MSG_HOST_GET_Msgno struct {
	Uid int32
}
type MSG_HOST_GET_Msgno_result struct {
	Msgno uint32
}

type MSG_HOST_QueryErr struct {
	Err   string
	Stack []byte
}
type MSG_HOST_ResetWindow struct {
	Window int32
}

//小文件上传下载
type MSG_FILE_upload struct {
	Name       string
	Title      string
	Data       []byte
	AddBy      int32
	ObjectType string
	ObjectID   int32
	Code       string //如果是项目的文件，需要“产品代号/项目代号”，会保存在对应的文件夹内
	Type       string //bindingFile sourceFile feedbackFile finalFile processFile
}
type MSG_FILE_upload_result struct {
	FileID int64
}
type MSG_FILE_getByID struct {
	FileID int64
	NoData bool //true不反回date
}
type MSG_FILE_getByID_result struct {
	FileID     int64
	Name       string
	Ext        string
	Size       int64
	Type       string
	AddedDate  time.Time
	ObjectType string
	ObjectID   int32
}
type MSG_FILE_updateMapByWhere struct {
	Where  map[string]interface{}
	Update map[string]interface{}
}

type MSG_FILE_DeleteByID struct {
	FileID int64
}
type MSG_HOST_BeginTransaction struct {
	TransactionNo uint32
}
type MSG_HOST_BeginTransaction_result struct {
	TransactionNo uint32
}
type MSG_HOST_Transaction_Commit struct {
	No uint32
}
type MSG_HOST_Transaction_RollBack struct {
	No uint32
}
type MSG_HOST_Transaction_Check struct {
	No uint32
}
type MSG_FILE_getByObject struct {
	ObjectType string
	ObjectID   int32
}
type MSG_FILE_getByObject_result struct {
	List []*MSG_FILE_getByID_result
}
type MSG_HOST_GET_MsgUserId struct {
}
type MSG_HOST_GET_MsgUserId_result struct {
	Uid int32
}

//上传大临时文件
type MSG_FILE_uploadTmp struct {
	Name      string
	Data      []byte
	Index     int //序号
	BlockSize int //单块大小
}
type MSG_FILE_RangeDown struct {
	FileID int64
	Start  int64
	End    int64
}
type MSG_FILE_RangeDown_result struct {
	Byte []byte
}

type MSG_FILE_updateTmp struct {
	Files []*MSG_FILE_upload
}

type MSG_FILE_edit struct {
	FileID int64
	Name   string
}
type MSG_FILE_getByWhere struct {
	Where   map[string]interface{}
	Page    int
	PerPage int
	Total   int
}
type MSG_FILE_getByWhere_result struct {
	List  []*MSG_FILE_getByID_result
	Total int
}

//打包下载
type MSG_FILE_download_byIds struct {
	Ids []int64
}

//生产一个临时fileID
type MSG_FILE_download_byIds_result struct {
	FileID int64
}

type MSG_HOST_getCenterSvrId struct {
	No uint8
}
type MSG_HOST_getCenterSvrId_result struct {
	Id uint16
}

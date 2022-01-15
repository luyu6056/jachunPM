package handler

//封装一个用于common fetch里面用的ws服务
import (
	"codec"
	"libraries"
	"protocol"
	"sync"

	"github.com/luyu6056/cache"
	"github.com/luyu6056/gnet"
)

type CommonFetch struct {
	buf        *libraries.MsgBuffer
	queryCache map[string][]string
	path       string
	oldws      HttpRequest
}

var fetchpool = sync.Pool{New: func() interface{} {
	return &CommonFetch{buf: &libraries.MsgBuffer{}}
}}

func getFetchInterface(oldws HttpRequest, path string, user *protocol.MSG_USER_INFO_cache) *TemplateData {
	f := fetchpool.Get().(*CommonFetch)
	f.oldws = oldws
	f.path = path
	return templateDataInit(f, user)
}
func putFetchInterface(f *CommonFetch) {
	f.queryCache = nil
	f.buf.Reset()
	fetchpool.Put(f)
}
func (f *CommonFetch) AddQuery(k, value string) {
	if f.queryCache == nil {
		f.queryCache = make(map[string][]string)
	}
	f.queryCache[k] = []string{value}
}
func (f *CommonFetch) Body() []byte {
	return nil
}
func (f *CommonFetch) Cookie(key string) string {
	return f.oldws.Cookie(key)
}
func (f *CommonFetch) DelSession() {

}
func (f *CommonFetch) GetAllPost() map[string][]string {
	return nil
}
func (f *CommonFetch) GetAllQuery() map[string][]string {
	return f.queryCache
}
func (f *CommonFetch) Header(name string) string {
	return f.oldws.Header(name)
}

func (f *CommonFetch) RemoteAddr() string {
	return f.oldws.RemoteAddr()
}
func (f *CommonFetch) IP() string {
	return f.oldws.RemoteAddr()
}
func (f *CommonFetch) Path() string {
	return f.path
}
func (f *CommonFetch) Query(key string) string {
	if f.queryCache == nil || len(f.queryCache[key]) == 0 {
		return ""
	}
	return f.queryCache[key][0]
}
func (f *CommonFetch) Post(key string) string {
	return ""
}
func (f *CommonFetch) PostSlice(key string) []string {
	return nil
}
func (f *CommonFetch) Session() *cache.Hashvalue {
	return f.oldws.Session()
}
func (f *CommonFetch) Method() string {
	return f.oldws.Method()
}
func (f *CommonFetch) SetCode(i int) {

}
func (f *CommonFetch) SetContentType(_ string) {

}
func (f *CommonFetch) SetCookie(name, value string, max_age uint32) {
	f.oldws.SetCookie(name, value, max_age)
}
func (f *CommonFetch) SetHeader(name, value string) {
	f.oldws.SetHeader(name, value)
}
func (f *CommonFetch) StaticHandler() gnet.Action {
	return gnet.None
}
func (f *CommonFetch) Write(b *libraries.MsgBuffer) {
	f.buf.Write(b.Bytes())
}
func (f *CommonFetch) WriteString(s string) {
	f.buf.WriteString(s)
}
func (f *CommonFetch) Redirect(url string) {
	f.oldws.Redirect(url)
}
func (f *CommonFetch) OutBuffer() []byte {
	return f.buf.Bytes()
}
func (f *CommonFetch) Close() {
	f.oldws.Close()
}
func (f *CommonFetch) RangeDownload(r codec.HttpIoReader, size int64, name string) {
	f.oldws.RangeDownload(r, size, name)
}
func (f *CommonFetch) URI() string {
	return f.oldws.URI()
}
func (f *CommonFetch) Referer() string {
	return f.oldws.Referer()
}

package libraries

//阿里云cdn
import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Cdn struct {
	param           map[string]string
	u               string
	accessKeySecret string
}

//填充cdn公共参数
func Newcdn(access_id string, accessKeySecret string) *Cdn {
	c := new(Cdn)
	c.accessKeySecret = accessKeySecret
	param := make(map[string]string)
	param["Format"] = "JSON"
	param["Version"] = "2014-11-11"
	param["AccessKeyId"] = access_id
	param["SignatureMethod"] = "HMAC-SHA1"
	param["SignatureVersion"] = "1.0"
	c.param = param
	c.u = "https://cdn.aliyuncs.com/"
	return c
}

//查询服务状态
func (this *Cdn) DescribeCdnService() (result map[string]interface{}, err error) {
	param := this.param
	param["Action"] = "DescribeCdnService"
	u := this.SHA1_signa(param)

	body, err := Http_get(this.u + "?" + u)
	if err == nil {
		err = JsonUnmarshalStr(body, &result)
	}
	return
}

//刷新缓存
func (this *Cdn) RefreshObjectCaches(ObjectPath string, ObjectType string) (result bool, err error) {
	param := this.param
	param["Action"] = "RefreshObjectCaches"
	param["ObjectPath"] = ObjectPath
	param["ObjectType"] = ObjectType
	u := this.SHA1_signa(param)
	str, err := Http_get(this.u + "?" + u)
	result = false
	if err == nil {
		var res map[string]interface{}
		err = JsonUnmarshalStr(str, &res)
		if res != nil && res["RefreshTaskId"] != nil {
			result = true
		}
	}
	return
}

//组装时间戳
func (this *Cdn) TimeStamp() string {
	timestamp := time.Now().UTC()
	t := timestamp.Format("2006-01-02T15:04:05Z")
	return t
}

//签名并组装url
func (this *Cdn) SHA1_signa(param map[string]string) string {
	//组装时间戳
	param["TimeStamp"] = this.TimeStamp()
	//生成随机数
	rand := md5.Sum([]byte(Microtime() + strconv.Itoa(Rand(100000, 999999))))
	param["SignatureNonce"] = fmt.Sprintf("%x", rand)
	data := []string{}
	for k, v := range param {
		data = append(data, k+"="+url.QueryEscape(v))
	}
	//对请求参数排序
	sort.Strings(data)
	//组装sha1签名
	u := "GET&%2F&" + url.QueryEscape(strings.Join(data, "&"))
	//HMAC的key
	key := []byte(this.accessKeySecret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(u))
	//对HMAC结果转base64
	b64 := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	data = append(data, "Signature="+url.QueryEscape(b64))
	return strings.Join(data, "&")
}

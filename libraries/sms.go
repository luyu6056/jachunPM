package libraries

import (
	"errors"
)

type Yunpian_result struct {
	Code  int     `json:"code"`  //0代表发送成功，其他code代表出错，详细见"返回值说明"页面
	Msg   string  `json:"msg"`   //例如""发送成功""，或者相应错误信息
	Count int     `json:"count"` //发送成功短信的计费条数(计费条数：70个字一条，超出70个字时按每67字一条计费)
	Fee   float64 `json:"fee"`   //扣费金额，单位：元，类型：双精度浮点型/double
	//Unit   string  `json:"unit"`   //计费单位；例如：“RMB”
	//Mobile string  `json:"mobile"` //发送手机号
	Sid int64 `json:"sid"` //短信id，64位整型， 对应Java和C#的long，不可用int解析
}

//云片短信
func Sms_send(mobile_number, msg, apikey string) (result *Yunpian_result, err error) {
	res, err := Http_post("https://sms.yunpian.com/v1/sms/send.json", map[string]string{
		"apikey": apikey,
		"text":   msg,
		"mobile": mobile_number,
	})
	if err != nil {
		return
	}
	err = JsonUnmarshalStr(res, &result)
	if err != nil || result == nil || result.Code != 0 {
		err = errors.New("云片发送短信返回解析失败，" + res)
	}
	return
}

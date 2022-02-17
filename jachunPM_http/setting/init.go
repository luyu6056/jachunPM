package setting

import (
	"io/ioutil"
	"libraries"
	"log"
)

var Setting struct {
	ListenHttp   string
	Origin       string
	HttpsTLScert string
	HttpsTLSkey  string
	HttpsTLSca   string
	HostIP       string
	MysqlDsn     string
	MysqlMaxConn int32
	RedisIP      string
	TokenKey     string //注册服务时候用于验证是否有效
	Version      string //前端版本
	Debug        bool
}

func init() {
	data, err := ioutil.ReadFile("./setting.json")
	if err != nil {
		log.Fatalf("无法读取配置文件config.json,错误%v", err)
	}
	err = libraries.JsonUnmarshal(data, &Setting)
	if err != nil {
		log.Fatalf("无法解析配置json,错误%v", err)
	}
	if Setting.Origin == "" {
		log.Fatalf("setting.json未设置Orgin")
	}

}

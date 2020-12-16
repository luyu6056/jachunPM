package config

import (
	"io/ioutil"
	"libraries"
	"log"
	"protocol"
)

var Config struct {
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
var Lang = make(map[protocol.CountryNo]map[string]map[string]interface{}) //map[name][key]=value

func init() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("无法读取配置文件config.json,错误%v", err)
	}
	err = libraries.JsonUnmarshal(data, &Config)
	if err != nil {
		log.Fatalf("无法解析配置json,错误%v", err)
	}
	//Lang[protocol.EN] = make(map[string]map[string]interface{})
	Lang[protocol.ZH_CN] = make(map[string]map[string]interface{})
	//LangEnInit()
	LangZH_CNInit()
}

package config

import (
	"io/ioutil"
	"libraries"
	"log"
)

var Config struct {
	ListenRpc    string
	MysqlDsn     string
	MysqlMaxConn int32
	TokenKey     string //注册服务时候用于验证是否有效
}

func init() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("无法读取配置文件config.json,错误%v", err)
	}
	err = libraries.JsonUnmarshal(data, &Config)
	if err != nil {
		log.Fatalf("无法解析配置json,错误%v", err)
	}
}

module jachunPM_commom

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace server => ../server

replace protocol => ../protocol

require (
	github.com/aliyun/aliyun-oss-go-sdk v2.1.4+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/klauspost/compress v1.11.4
	github.com/luyu6056/cache v1.1.4
	github.com/luyu6056/gnet v1.2.8
	github.com/luyu6056/tls v0.15.1
	github.com/panjf2000/ants/v2 v2.4.3
	github.com/panjf2000/gnet v1.3.0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.1.0 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
)

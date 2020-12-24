module jachunPM_http

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace server => ../server

replace protocol => ../protocol

require (
	github.com/dlclark/regexp2 v1.4.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/luyu6056/cache v1.1.2
	github.com/luyu6056/gnet v1.2.8
	github.com/luyu6056/tls v0.15.1
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/panjf2000/gnet v1.3.2 // indirect
	github.com/rubenfonseca/fastimage v0.0.0-20170112075114-7e006a27a95b
	github.com/vmihailenco/msgpack/v5 v5.1.0 // indirect
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
)

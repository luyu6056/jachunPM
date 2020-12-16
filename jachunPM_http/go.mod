module jachunPM_http

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace server => ../server

replace protocol => ../protocol

require (
	github.com/dlclark/regexp2 v1.4.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/luyu6056/cache v1.0.6
	github.com/luyu6056/gnet v1.2.8
	github.com/luyu6056/tls v0.15.1
	github.com/modern-go/reflect2 v1.0.1 // indirect
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
)

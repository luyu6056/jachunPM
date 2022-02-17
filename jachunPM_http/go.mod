module jachunPM_http

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace jachunPM/image => ../libraries/image

replace protocol => ../protocol

replace codec => ../net/codec

replace github.com/luyu6056/gnet => ../net/gnet

replace config => ../config

require (
	codec v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	github.com/dlclark/regexp2 v1.4.0
	github.com/fsnotify/fsnotify v1.5.1
	github.com/luyu6056/cache v1.1.12
	github.com/luyu6056/gnet v1.3.2
	github.com/luyu6056/reflect2 v1.0.2
	github.com/luyu6056/tls v0.15.1
	github.com/panjf2000/ants/v2 v2.4.7
	github.com/rubenfonseca/fastimage v0.0.0-20170112075114-7e006a27a95b
	github.com/valyala/gozstd v1.16.0 // indirect
	github.com/xuri/excelize/v2 v2.5.0
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb
	jachunPM/image v0.0.0-00010101000000-000000000000
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000

)

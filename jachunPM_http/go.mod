module jachunPM_http

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace codec => ../codec

replace jachunPM/image => ../image

replace protocol => ../protocol

replace github.com/luyu6056/gnet => ../gnet

require (
	codec v0.0.0-00010101000000-000000000000
	github.com/dlclark/regexp2 v1.4.0
	github.com/fsnotify/fsnotify v1.5.1
	github.com/luyu6056/cache v1.1.11
	github.com/luyu6056/gnet v1.3.2
	github.com/luyu6056/reflect2 v1.0.2
	github.com/luyu6056/tls v0.15.1
	github.com/rubenfonseca/fastimage v0.0.0-20170112075114-7e006a27a95b
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5
	jachunPM/image v0.0.0-00010101000000-000000000000
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000

)

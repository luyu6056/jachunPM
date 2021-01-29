module jachunPM_http

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace server => ../server

replace jachunPM/image => ../image

replace protocol => ../protocol

require (
	github.com/dlclark/regexp2 v1.4.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/kr/pretty v0.1.0 // indirect
	github.com/luyu6056/cache v1.1.4
	github.com/luyu6056/gnet v1.2.8
	github.com/luyu6056/tls v0.15.1
	github.com/rubenfonseca/fastimage v0.0.0-20170112075114-7e006a27a95b
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5
	golang.org/x/sys v0.0.0-20201116194326-cc9327a14d48 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	jachunPM/image v0.0.0-00010101000000-000000000000
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
)

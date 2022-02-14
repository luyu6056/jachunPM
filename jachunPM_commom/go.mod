module jachunPM_commom

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace codec => ../codec

replace protocol => ../protocol

replace github.com/luyu6056/gnet => ../gnet

require (
	github.com/kr/pretty v0.1.0 // indirect
	github.com/luyu6056/cache v1.1.12
	github.com/luyu6056/gnet v1.3.2
	github.com/panjf2000/ants/v2 v2.4.7
	github.com/valyala/gozstd v1.16.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
)

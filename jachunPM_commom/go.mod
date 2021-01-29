module jachunPM_commom

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace server => ../server

replace protocol => ../protocol

require (
	github.com/klauspost/compress v1.11.4
	github.com/kr/pretty v0.1.0 // indirect
	github.com/luyu6056/cache v1.1.4
	github.com/luyu6056/gnet v1.2.8
	github.com/panjf2000/ants/v2 v2.4.3
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
)

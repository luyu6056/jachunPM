module jachunPM_log

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace github.com/luyu6056/gnet => ../net/gnet

replace protocol => ../protocol

replace config => ../config

require (
	config v0.0.0-00010101000000-000000000000
	github.com/luyu6056/gnet v1.3.4 // indirect
	github.com/valyala/gozstd v1.16.0 // indirect
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
)

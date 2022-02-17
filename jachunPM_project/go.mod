module jachunPM_project

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace protocol => ../protocol

replace github.com/luyu6056/gnet => ../net/gnet

replace config => ../config

require (
	config v0.0.0-00010101000000-000000000000
	github.com/luyu6056/cache v1.1.6 // indirect
	github.com/valyala/gozstd v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
)

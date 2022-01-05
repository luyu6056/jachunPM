module jachunPM_project

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace server => ../server

replace protocol => ../protocol

replace github.com/luyu6056/gnet => ../gnet

require (
	github.com/luyu6056/cache v1.1.6 // indirect
	libraries v0.0.0-00010101000000-000000000000
	mysql v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
)

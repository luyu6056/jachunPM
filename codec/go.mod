module codec

go 1.15

replace mysql => ../mysql

replace libraries => ../libraries

replace codec => ../codec

replace protocol => ../protocol

replace github.com/luyu6056/gnet => ../gnet

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/klauspost/compress v1.14.1
	github.com/luyu6056/cache v1.1.4
	github.com/luyu6056/gnet v1.2.8
	github.com/panjf2000/ants/v2 v2.4.3
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	libraries v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
)

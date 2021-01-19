module server

go 1.15

require (
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/klauspost/compress v1.11.4
	github.com/luyu6056/cache v1.1.4
	github.com/luyu6056/gnet v1.2.8
	github.com/panjf2000/ants/v2 v2.4.3
	github.com/rubenfonseca/fastimage v0.0.0-20170112075114-7e006a27a95b // indirect
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5 // indirect
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102
	libraries v0.0.0-00010101000000-000000000000
	protocol v0.0.0-00010101000000-000000000000
)

replace libraries => ../libraries

replace protocol => ../protocol

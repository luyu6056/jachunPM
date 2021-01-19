module protocol

go 1.15

replace libraries => ./libraries

replace mysql => ./mysql

require (
	github.com/klauspost/compress v1.11.4
	github.com/luyu6056/cache v1.1.4
	github.com/luyu6056/gnet v1.2.8
	github.com/luyu6056/reflect2 v1.0.2
)

module github.com/gomsa/goods-api

go 1.12

require (
	github.com/bigrocs/barcode v0.0.0-20190809083253-fae6bd05b0de // indirect
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/gomsa/tools v0.0.0-20190808000024-44c2fc0e494a
	github.com/gomsa/user v0.0.0-20190713014816-bcd837144465
	github.com/gomsa/user-api v0.0.0-20190713033333-7a0896d1f7a3 // indirect
	github.com/micro/go-micro v1.7.0
	github.com/micro/kubernetes v0.7.0
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
)

replace github.com/micro/go-micro => github.com/micro/go-micro v1.7.0

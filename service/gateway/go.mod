module muxi-workbench-gateway

// exclude github.com/micro/go-plugins v1.5.1

go 1.14

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v2.0.1+incompatible
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	//github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	//github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	github.com/opentracing/opentracing-go v1.2.0
	github.com/qiniu/api.v7/v7 v7.8.2
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v3.21.8+incompatible
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.3.3
	github.com/swaggo/swag v1.7.4
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/willf/pad v0.0.0-20190207183901-eccfe5d84172
	go.etcd.io/bbolt v1.3.3 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210908191846-a5e095526f91 // indirect
	golang.org/x/sys v0.0.0-20210909193231-528a39cd75f3 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.5 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	muxi-workbench v0.0.0-00010101000000-000000000000
	muxi-workbench-attention v0.0.0-00010101000000-000000000000
	muxi-workbench-feed v0.0.0-00010101000000-000000000000
	muxi-workbench-project v0.0.0-00010101000000-000000000000
	muxi-workbench-status v0.0.0-00010101000000-000000000000
	muxi-workbench-team v0.0.0-00010101000000-000000000000
	muxi-workbench-user v0.0.0-00010101000000-000000000000
)

replace muxi-workbench-gateway => ./

replace muxi-workbench-feed => ../feed

replace muxi-workbench-attention => ../attention

replace muxi-workbench-status => ../status

replace muxi-workbench-project => ../project

replace muxi-workbench => ../../

replace muxi-workbench-user => ../user

replace muxi-workbench-team => ../team

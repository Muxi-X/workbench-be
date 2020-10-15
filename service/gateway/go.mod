module muxi-workbench-gateway

go 1.13

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v2.19.11+incompatible
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	github.com/urfave/cli/v2 v2.2.0 // indirect
	github.com/willf/pad v0.0.0-20190207183901-eccfe5d84172
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20201010224723-4f7140c49acb // indirect
	golang.org/x/tools v0.0.0-20201013053347-2db1cd791039 // indirect
	gopkg.in/go-playground/validator.v9 v9.30.2 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
	muxi-workbench v0.0.0-00010101000000-000000000000
	muxi-workbench-feed v0.0.0-00010101000000-000000000000
	muxi-workbench-project v0.0.0-00010101000000-000000000000
	muxi-workbench-status v0.0.0-00010101000000-000000000000
	muxi-workbench-user v0.0.0-00010101000000-000000000000
)

replace muxi-workbench-gateway => ./

replace muxi-workbench-feed => ../feed

replace muxi-workbench-status => ../status

replace muxi-workbench-project => ../project

replace muxi-workbench => ../../

replace muxi-workbench-user => ../user

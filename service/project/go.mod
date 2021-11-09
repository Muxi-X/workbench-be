module muxi-workbench-project

replace muxi-workbench => ../../

replace muxi-workbench-user => ../user

replace muxi-workbench-attention => ../attention

exclude github.com/micro/go-plugins v1.5.1

go 1.13

require (
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.15
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	github.com/opentracing/opentracing-go v1.2.0
	github.com/spf13/viper v1.7.1
	github.com/swaggo/swag v1.7.4
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
	gorm.io/gorm v1.22.2
	muxi-workbench v0.0.0-00010101000000-000000000000
	muxi-workbench-user v0.0.0-00010101000000-000000000000
	muxi-workbench-attention v0.0.0-00010101000000-000000000000
)

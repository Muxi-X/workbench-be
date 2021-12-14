module muxi-workbench-attention

replace muxi-workbench => ../../

replace muxi-workbench-project => ../project

replace muxi-workbench-user => ../user

replace muxi-workbench-team => ../team

replace muxi-workbench-attention => ./

exclude github.com/micro/go-plugins v1.5.1

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
	muxi-workbench v0.0.0-00010101000000-000000000000
	muxi-workbench-project v0.0.0-00010101000000-000000000000
	muxi-workbench-user v0.0.0-00010101000000-000000000000
)

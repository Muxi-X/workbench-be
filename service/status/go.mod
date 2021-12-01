module muxi-workbench-status

replace muxi-workbench => ../../

exclude github.com/micro/go-plugins v1.5.1

go 1.13

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.15
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	github.com/opentracing/opentracing-go v1.2.0
	github.com/spf13/viper v1.7.1
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
	muxi-workbench v0.0.0-00010101000000-000000000000
)

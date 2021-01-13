module muxi-workbench-team

replace muxi-workbench => ../../

replace muxi-workbench-user => ../user

go 1.14

require (
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.15
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/spf13/viper v1.7.1 // indirectopentracing "github.com/opentracing/opentracing-go"
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	muxi-workbench v0.0.0-00010101000000-000000000000
	muxi-workbench-user v0.0.0-00010101000000-000000000000
)

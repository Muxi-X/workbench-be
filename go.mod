module github.com/Muxi-X/workbench-be

replace muxi-workbench => ./

go 1.12

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro v1.18.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/spf13/viper v1.6.2
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	go.uber.org/zap v1.14.1
	muxi-workbench v0.0.0-00010101000000-000000000000
)

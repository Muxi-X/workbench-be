module github.com/Muxi-X/workbench-be

replace muxi-workbench => ./

go 1.12

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.1 // indirect
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/jinzhu/gorm v1.9.15
	github.com/micro/go-micro v1.18.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.7.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.16.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	muxi-workbench v0.0.0-00010101000000-000000000000
)

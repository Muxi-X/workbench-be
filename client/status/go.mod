module muxi-workbench-status-client

replace muxi-workbench => ../../

replace muxi-workbench-status => ../../service/status

go 1.13

require (
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	google.golang.org/grpc v1.25.1
	muxi-workbench v0.0.0-00010101000000-000000000000
	muxi-workbench-status v0.0.0-00010101000000-000000000000
)

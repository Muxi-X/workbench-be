module muxi-workbench-project-client

replace muxi-workbench => ../../

replace muxi-workbench-project => ../../service/project

go 1.13

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	muxi-workbench v0.0.0-00010101000000-000000000000
	muxi-workbench-project v0.0.0-00010101000000-000000000000
)

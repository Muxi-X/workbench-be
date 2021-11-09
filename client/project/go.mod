module muxi-workbench-project-client

replace muxi-workbench => ../../

replace muxi-workbench-project => ../../service/project

replace (
	muxi-workbench-user => ../../service/user
)

go 1.13

require (
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	muxi-workbench v0.0.0-00010101000000-000000000000
	muxi-workbench-project v0.0.0-00010101000000-000000000000
	muxi-workbench-user v0.0.0-00010101000000-000000000000
)

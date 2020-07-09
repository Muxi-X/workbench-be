module muxi-workbench-team-client

replace muxi-workbench => ../../

replace muxi-workbench-feed => ../../service/feed

replace muxi-workbench-project => ../../service/project

replace muxi-workbench-user => ../../service/user

replace muxi-workbench-team => ../../service/team

go 1.14

require (
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.24.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	muxi-workbench v0.0.0-00010101000000-000000000000 // indirect
	muxi-workbench-team v0.0.0-00010101000000-000000000000 // indirect
)

module muxi-workbench-attention

replace muxi-workbench => ../../

replace muxi-workbench-project => ../project

replace muxi-workbench-user => ../user

replace muxi-workbench-team => ../team

replace muxi-workbench-attention => ./

exclude github.com/micro/go-plugins v1.5.1

go 1.13

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/golang/protobuf v1.3.5
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.2 // indirect
	github.com/jinzhu/now v1.1.3 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
	golang.org/x/tools v0.1.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	muxi-workbench v0.0.0-00010101000000-000000000000
)

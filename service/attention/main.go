package main

import (
	"flag"
	"log"

	pb "muxi-workbench-attention/proto"
	s "muxi-workbench-attention/service"
	"muxi-workbench/config"
	logger "muxi-workbench/log"
	"muxi-workbench/model"
	"muxi-workbench/pkg/handler"
	"muxi-workbench/pkg/tracer"

	_ "github.com/micro/go-plugins/registry/kubernetes"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

// 使用--sub运行subscribe服务
// 否则默认运行feed服务
var subFg = flag.Bool("sub", false, "use subscribe service mode")

func init() {
	s.UserInit()
	s.ProjectInit()
}

// 包含两个服务：feed服务和subscribe服务
// subscribe服务 --> 异步将feed数据写入数据库
func main() {
	flag.Parse()

	var err error

	// init config
	if !*subFg {
		// feed-service
		err = config.Init("./conf/config.yaml", "WORKBENCH_FEED")
	} else {
		// sub-service
		err = config.Init("./conf/config_sub.yaml", "WORKBENCH_SUB")
	}

	if err != nil {
		panic(err)
	}

	t, io, err := tracer.NewTracer(viper.GetString("local_name"), viper.GetString("tracing.jager"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	defer logger.SyncLogger()

	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	// init db
	model.DB.Init()
	defer model.DB.Close()

	srv := micro.NewService(
		micro.Name(viper.GetString("local_name")),
		micro.WrapHandler(
			opentracingWrapper.NewHandlerWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapHandler(handler.ServerErrorHandlerWrapper()),
		micro.Flags(cli.BoolFlag{
			Name:   "sub",
			Usage:  "use subscribe service mode",
			Hidden: false,
		}),
		//micro.Broker(bro),
	)

	// Init will parse the command line flags.
	srv.Init()

	if !*subFg {
		// feed-service

		// init redis db
		model.RedisDB.Init()
		defer model.RedisDB.Close()

		pb.RegisterFeedServiceHandler(srv.Server(), &s.FeedService{})

		// Run the server
		if err := srv.Run(); err != nil {
			logger.Error(err.Error())
		}
	}
	// } else {
	// 	// sub-service
	//
	// 	// init redis pub/sub client
	// 	model.PubSubClient.Init(m.RdbChan)
	// 	defer model.PubSubClient.Close()
	//
	// 	logger.Info("Subscribe service start...")
	// 	s.SubServiceRun()
	// }
}

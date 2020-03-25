package main

import (
	"flag"
	"github.com/micro/cli"
	"log"

	m "muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	s "muxi-workbench-feed/service"
	"muxi-workbench/config"
	logger "muxi-workbench/log"
	"muxi-workbench/model"
	"muxi-workbench/pkg/handler"
	"muxi-workbench/pkg/tracer"

	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

// 使用--sub运行subscribe服务
// 否则默认运行feed服务
var (
	subFg = flag.Bool("sub", false, "use subscribe service mode")
	//subFg *bool
)

func Init() {
}

func main() {
	flag.Parse()

	var err error

	// init config
	if *subFg {
		err = config.Init("./conf/config_sub.yaml", "WORKBENCH_SUB")
	} else {
		err = config.Init("./conf/config.yaml", "WORKBENCH_FEED")
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

	// init redis db for subscribe service
	if *subFg {
		m.PubRdb = m.OpenRedisClient()
		m.SubRdb = m.OpenRedisClient().Subscribe(m.RdbChan)
		defer m.PubRdb.Close()
		defer m.SubRdb.Close()
	}

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
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	_ = pb.RegisterFeedServiceHandler(srv.Server(), &s.FeedService{})

	// Run the server
	if err := srv.Run(); err != nil {
		logger.Error(err.Error())
	}
}

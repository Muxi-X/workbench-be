package main

import (
	"flag"
	"log"

	m "muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	s "muxi-workbench-feed/service"
	"muxi-workbench/config"
	logger "muxi-workbench/log"
	"muxi-workbench/model"
	"muxi-workbench/pkg/handler"
	"muxi-workbench/pkg/tracer"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

// 使用--sub运行subscribe服务
// 否则默认运行feed服务
var (
	subFg = flag.Bool("sub", false, "use subscribe service mode")
)

// 包含两个服务：feed服务和subscribe服务
// subscribe服务 --> 异步将feed数据写入数据库
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
	rdbClient := m.OpenRedisClient()
	if _, err = rdbClient.Ping().Result(); err != nil {
		panic(err)
	}

	if !*subFg {
		m.PubRdb = rdbClient
		defer m.PubRdb.Close()
	} else {
		m.SubRdb = rdbClient.Subscribe(m.RdbChan)
		defer m.SubRdb.Close()
	}

	//redis broker
	//bro := redis.NewBroker(broker.Addrs("redis://user:root@localhost:6379"))
	//if err := bro.Init(); err != nil {
	//	panic(err)
	//}
	//
	//if err := bro.Connect(); err != nil {
	//	panic(err)
	//}
	//bro.Subscribe()

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
		// Register handler
		pb.RegisterFeedServiceHandler(srv.Server(), &s.FeedService{})
	} else {
		go s.SubServiceRun()
	}

	// Run the server
	if err := srv.Run(); err != nil {
		logger.Error(err.Error())
	}
}

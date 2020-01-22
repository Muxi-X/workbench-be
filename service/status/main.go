package main

import (
	"log"

	pb "muxi-workbench-status/proto"
	s "muxi-workbench-status/service"
	"muxi-workbench/config"
	logger "muxi-workbench/log"
	"muxi-workbench/model"
	"muxi-workbench/pkg/handler"
	tracer "muxi-workbench/pkg/tracer"

	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

func main() {

	t, io, err := tracer.NewTracer("workbench.service.status", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	defer logger.SyncLogger()

	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	// init config
	if err := config.Init("", "WORKBENCH_STATUS"); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	srv := micro.NewService(
		micro.Name(viper.GetString("local_name")),
		micro.WrapHandler(
			opentracingWrapper.NewHandlerWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapHandler(handler.ServerErrorHandlerWrapper()),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterStatusServiceHandler(srv.Server(), &s.StatusService{})

	// Run the server
	if err := srv.Run(); err != nil {
		logger.Error(err.Error())
	}
}

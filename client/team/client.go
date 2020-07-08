package main

import (
	"github.com/opentracing/opentracing-go"
	"muxi-workbench-team-client/tracer"
	"muxi-workbench/pkg/handler"
	"github.com/micro/go-micro"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	t, io, err := tracer.NewTracer("workbench.cli.team", address)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(micro.Name("workbench.cli.team"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()))
	service.Init()

	client := pb.NewProjectServiceClient("workbench.service.team", service.Client())

}

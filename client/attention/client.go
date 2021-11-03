package main

import (
	"context"
	"fmt"
	"log"

	"muxi-workbench-attention-client/tracer"
	pb "muxi-workbench-attention/proto"
	"muxi-workbench/pkg/handler"

	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

const address = "localhost:50051"

func main() {
	t, io, err := tracer.NewTracer("workbench.cli.attention", address)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("workbench.cli.attention"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	service.Init()

	client := pb.NewAttentionServiceClient("workbench.service.attention", service.Client())

	// 获取attention
	req := &pb.ListRequest{
		LastId: 167,
		Limit:  5,
		UserId: 100,
	}
	resp, err := client.List(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("11111111111")
	fmt.Println(resp)
	fmt.Println(resp.List[0])

	// 新增attention
	addReq := &pb.PushRequest{
		UserId: 100,
		DocId:  6666, // status id
	}
	addResp, err := client.Create(context.Background(), addReq)
	if err != nil {
		panic(err)
	}
	fmt.Println(addResp)
}

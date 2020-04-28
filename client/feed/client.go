package main

import (
	"context"
	"fmt"
	"log"

	"muxi-workbench-feed-client/tracer"
	pb "muxi-workbench-feed/proto"
	"muxi-workbench/pkg/handler"

	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

const address = "localhost:50051"

func main() {
	t, io, err := tracer.NewTracer("workbench.cli.feed", address)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(micro.Name("workbench.cli.feed"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	service.Init()

	client := pb.NewFeedServiceClient("workbench.service.feed", service.Client())

	// 获取feed
	req := &pb.ListRequest{
		LastId: 67,
		Limit:  5,
		Role:   3,
		UserId: 53,
		Filter: &pb.Filter{
			UserId:  0,
			GroupId: 3,
		},
	}

	resp, err := client.List(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// 新增feed
	//addReq := &pb.PushRequest{
	//	Action: "创建",
	//	UserId: 2333,
	//	Source: &pb.Source{
	//		Kind:        6,
	//		Id:          6666,   // status id
	//		Name:        "测试数据", // 进度标题
	//		ProjectId:   0,      // 固定数据
	//		ProjectName: "",     // 固定数据
	//	},
	//}
	//addResp, err := client.Push(context.Background(), addReq)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(addResp)
}

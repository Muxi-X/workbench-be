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

const (
	address = "localhost:50051"
)

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

	req := &pb.ListRequest{
		Page:   2,
		Size:   5,
		LastId: 400,
	}

	resp, err := client.List(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	//fmt.Println(resp.PageMax, resp.RowsNum)

	reqForUser := &pb.PersonalListRequest{
		UserId: 1,
		Page:   2,
		Size:   5,
		LastId: 400,
	}

	respForUser, err := client.PersonalList(context.Background(), reqForUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(respForUser)
	fmt.Println(respForUser.PageMax, respForUser.RowsNum)

	// 新feed测试数据，创建status
	addReq := &pb.AddRequest{
		Action: "创建",
		User: &pb.User{
			Name:      "测试",
			Id:        2333,
			AvatarUrl: "",
		},
		Source: &pb.Source{
			KindId:      6,
			ObjectId:    6666, // status id
			ObjectName:  "测试数据", // 进度标题
			ProjectId:   -1, // 固定数据
			ProjectName: "noname", // 固定数据
		},
	}
	addResp, err := client.Add(context.Background(), addReq)
	if err != nil {
		panic(err)
	}
	fmt.Println(addResp)
}

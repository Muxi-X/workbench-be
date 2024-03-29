package main

import (
	"context"
	"fmt"
	"log"

	"muxi-workbench-user-client/tracer"
	pb "muxi-workbench-user/proto"
	"muxi-workbench/pkg/handler"

	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

const address = "localhost:50051"

func main() {
	t, io, err := tracer.NewTracer("workbench.cli.user", address)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(micro.Name("workbench.cli.user"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	service.Init()

	client := pb.NewUserServiceClient("workbench.service.user", service.Client())

	// get user list
	// func() {
	// 	req := &pb.ListRequest{
	// 		LastId: 0,
	// 		Offset: 5,
	// 		Limit:  10,
	// 		Team:   1,
	// 		Group:  2,
	// 	}

	// 	resp, err := client.List(context.Background(), req)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(resp)
	// }()

	// register
	func() {
		req := &pb.RegisterRequest{
			Email:    "muxi@304.com",
			Name:     "muxi",
			Password: "muxi",
		}

		resp, err := client.Register(context.Background(), req)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}()

	// login
	// func() {
	// 	req := &pb.LoginRequest{
	// 		OauthCode: "NKPSPF0IOWWZSACLZ0OKAQ",
	// 	}

	// 	resp, err := client.Login(context.Background(), req)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(resp)
	// }()
}

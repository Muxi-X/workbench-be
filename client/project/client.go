package main

import (
	"context"
	"fmt"
	"log"

	"github.com/opentracing/opentracing-go"

	tracer "muxi-workbench-project-client/tracer"

	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"

	pb "muxi-workbench-project/proto"

	handler "muxi-workbench/pkg/handler"

	micro "github.com/micro/go-micro"
)

const (
	address = "localhost:50051"
)

func main() {
	t, io, err := tracer.NewTracer("workbench.cli.project", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(micro.Name("workbench.cli.project"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()))
	service.Init()

	client := pb.NewProjectServiceClient("workbench.service.project", service.Client())

	req := &pb.GetProjectListRequest{
		UserId:     22,
		Offset:     0,
		Limit:      20,
		Lastid:     0,
		Pagination: false,
	}

	resp, err := client.GetProjectList(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	fmt.Println(resp.List)

	// new request context
	// ctx := metadata.NewContext(context.Background(), map[string]string{"key1": "val1", "key2": "val2"})
	// add key-value pairs of metadata to context

	// req := &pb.CreateRequest{
	// 	UserId:  0,
	// 	Title:   "哈哈哈哈😁",
	// 	Content: "后废物废物分为",
	// }

	// _, err = client.Create(context.Background(), req)

	// // span.SetTag("req", req)
	// // span.SetTag("resp", resp)

	// if err != nil {
	// 	// span.SetTag("err", err)
	// 	log.Fatalf("Could not greet: %v", err)
	// }

	// resp, err := client.Get(context.Background(), &pb.GetRequest{
	// 	Id: 1,
	// })

	// if err != nil {
	// 	log.Fatalf("Could not greet: %v", err)
	// }

	// fmt.Println(resp.Status.Title)

	// resp, err := client.List(context.Background(), &pb.ListRequest{
	// 	Offset: 0,
	// 	Limit:  20,
	// 	Lastid: 162,
	// 	Group:  3,
	// 	Uid:    0,
	// })

	// if err != nil {
	// 	log.Fatalf("Could not greet: %v", err)
	// }

	// fmt.Println(resp.List, resp.Count)

	//
	//req := &pb.CreateCommentRequest{
	//	UserId:  0,
	//	StatusId: 3488,
	//	Content: "后废物废物分为",
	//}
	//
	//_, err = client.CreateComment(context.Background(), req)

}

package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"log"

	tracer "muxi-workbench-status-client/tracer"

	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"

	pb "muxi-workbench-status/proto"

	handler "muxi-workbench/pkg/handler"

	micro "github.com/micro/go-micro"
)

const (
	address = "localhost:50051"
)

func main() {
	t, io, err := tracer.NewTracer("workbench.cli.status", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(micro.Name("workbench.cli.status"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()))
	service.Init()

	client := pb.NewStatusServiceClient("workbench.service.status", service.Client())

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


	req := &pb.CommentListRequest{
		StatusId: 3488,
		Offset: 0,
		Limit: 20,
		Lastid: 0,
	}

	resp, err := client.ListComment(context.Background(), req)
	fmt.Println(resp.List, resp.Count)
}

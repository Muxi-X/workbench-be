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

	// getDocChildren
	// req := &pb.GetRequest{
	// 	Id: 173,
	// }
	// resp, err := client.GetDocChildren(context.Background(), req)
	// if err != nil {
	// 	log.Fatal("Could not greet: %v", err)
	// }
	// fmt.Println(resp)

	// getDocDetail
	// TODO: 测试删除能否获取
	// req := &pb.GetFileDetailRequest{
	// 	Id:       uint32(173),
	// 	FatherId: uint32(3),
	// }
	// resp, err := client.GetDocDetail(context.Background(), req)
	// if err != nil {
	// 	log.Fatal("Could not greet: %v", err)
	// }
	// fmt.Printf("%+v\n", *resp)

	// getDocFolderInfoList
	// req := &pb.GetInfoByIdsRequest{
	// 	List:     []uint32{173, 174, 176, 190},
	// 	FatherId: 3,
	// }
	// resp, err := client.GetDocFolderInfoList(context.Background(), req)
	// if err != nil {
	// 	log.Fatal("Could not greet: %v", err)
	// }
	// fmt.Printf("%+v\n", *resp)

	// getDocInfoList
	// req := &pb.GetInfoByIdsRequest{
	// 	List:     []uint32{6, 70, 88, 93, 99, 103, 105, 107, 153},
	// 	FatherId: 173,
	// }
	// resp, err := client.GetDocInfoList(context.Background(), req)
	// if err != nil {
	// 	log.Fatal("Could not greet: %v", err)
	// }
	// fmt.Printf("%+v\n", *resp)

	// getFileChildren
	req := &pb.GetRequest{
		Id: 29,
	}
	resp, err := client.GetFileChildren(context.Background(), req)
	if err != nil {
		log.Fatal("Could not greet: %v", err)
	}
	fmt.Printf("%+v\n", *resp)

	// getFileDetail
	//req := &pb.GetFileDetailRequest{
	//	Id:       69,
	//	FatherId: 29,
	//}
	//resp, err := client.GetFileDetail(context.Background(), req)
	//if err != nil {
	//	log.Fatal("Could not greet: %v", err)
	//}
	//fmt.Printf("%+v\n", *resp)

	// getFileFolderInfoList
	//req := &pb.GetInfoByIdsRequest{
	//	List: []uint32{},
	//	FatherId: uint32(),
	//}

	// 创建文档
	// req := &pb.CreateDocRequest{
	// 	Title:                 "测试",
	// 	Content:               "测试",
	// 	ProjectId:             2,
	// 	UserId:                1,
	// 	TeamId:                1,
	// 	FatherId:              1,
	// 	ChildrenPositionIndex: 0,
	// }
	// resp, err := client.CreateDoc(context.Background(), req)
	// if err != nil {
	// 	log.Fatalf("Could not greet: %v", err)
	// }
	// fmt.Println(resp)

}

package main

import (
	"context"
	"fmt"
	"log"

	"muxi-workbench-team-client/tracer"
	pb "muxi-workbench-team/proto"
	"muxi-workbench/pkg/handler"

	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
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
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	service.Init()

	client := pb.NewTeamServiceClient("workbench.service.team", service.Client())

    //列举applicationlist
	req := &pb.ApplicationListRequest{
		Offset:               1,
		Limit:                2,
		Pagination:           true,
	}
	resp, err := client.GetApplications(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range resp.List{
		fmt.Println(item.Id,item.Name,item.Email)
	}
	fmt.Println(resp.Count)


	/*
    //列举group中的members
	req := &pb.MemberListRequest{
		GroupId:    28,
		Offset:     1,
		Limit:      3,
		Pagination: false,
	}
    resp, err := client.GetMemberList(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range resp.List{
		fmt.Println(item.Name, item.TeamId, item.GroupId, item.GroupName, item.Email, item.Id)
	}
	fmt.Println(resp.Count)
	 */

    /*
    //列举grouplist
	req := &pb.GroupListRequest{
		Offset:               1,
		Limit:                2,
		Pagination:           false,
	}

	resp, err := client.GetGroupList(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.List)
	fmt.Println(resp.Count)
     */

}

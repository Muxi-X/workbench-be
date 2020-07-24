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

    //update members
	req := &pb.UpdateMembersRequest{
		GroupId:              37,
		Role:                 7,
		UserList:             []uint32{1,4,5},
	}
	resp, err := client.UpdateMembersForGroup(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)


	/*
	//list all members
	req := &pb.MemberListRequest{
		GroupId:              37,
		Offset:               0,
		Limit:                2,
		Pagination:           true,
	}
	resp, err := client.GetMemberList(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	//list all groups
	req := &pb.GroupListRequest{
		Offset:               2,
		Limit:                3,
		Pagination:           true,
	}
	resp, err := client.GetGroupList(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	//test updategroupinfo
	req := &pb.UpdateGroupInfoRequest{
		GroupId:              37,
		NewName:              "产品",
		Role:                 7,
	}
	resp, err := client.UpdateGroupInfo(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	//test deletegroup
	req := &pb.DeleteGroupRequest{
		GroupId:              38,
		Role:                 7,
	}
	resp, err := client.DeleteGroup(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */


	/*
    //test createGroup
    req := &pb.CreateGroupRequest{
		GroupName:            "产品",
		Role:                 7,
		UserList:             []uint32{4,5},
	}
	resp, err := client.CreateGroup(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */



}

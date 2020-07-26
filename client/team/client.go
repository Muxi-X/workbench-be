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


	//remove from team
	req := &pb.RemoveRequest{
		UserId:               5,
		TeamId:               3,
		Role:                 7,
	}
	resp, err := client.Remove(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	/*
	//join team
	req := &pb.JoinRequest{
		UserId:               5,
		TeamId:               3,
		Role:                 7,
	}
	resp, err := client.Join(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	//update teaminfo
	req := &pb.UpdateTeamInfoRequest{
		TeamId:               3,
		NewName:              "muxi2",
		Role:                 7,
	}
	resp, err := client.UpdateTeamInfo(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	// drop team
	req := &pb.DropTeamRequest{
		TeamId:               2,
		Role:                 7,
	}
	resp, err := client.DropTeam(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	//create team
	req := &pb.CreateTeamRequest{
		TeamName:             "muxi2",
		CreatorId:            1,
		Role:                 0,
	}
	resp, err := client.CreateTeam(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	//parse invitation
	req := &pb.ParseInvitationRequest{Hash:"gDrc51BJZF3m6hRigZ6WRg=="}
	resp, err := client.ParseInvitation(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	//create invitation
	req := &pb.CreateInvitationRequest{
		TeamId:               1,
		Expired:              3600,
	}
	resp, err  := client.CreateInvitation(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */


	/*
	//delete apply
	req := &pb.ApplicationRequest{UserId: 1}
	resp, err := client.DeleteApplication(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	//create apply
	req := &pb.ApplicationRequest{
		UserId:               1,
	}
	resp, err := client.CreateApplication(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
	//get applys
	req := &pb.ApplicationListRequest{
		Offset:     1,
		Limit:      0,
		Pagination: true,
	}
	resp, err := client.GetApplications(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */

	/*
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

	//test updategroupinfo
	req := &pb.UpdateGroupInfoRequest{
		GroupId:              28,
		NewName:              "产",
		Role:                 7,
	}
	resp, err := client.UpdateGroupInfo(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	 */
}

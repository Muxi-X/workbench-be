package service

import (
	"context"
	"log"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
)

//Create …… 生成邀请码
func (ts *TeamService) CreateInvitation(ctx context.Context, req *pb.CreateInvitationRequest, res *pb.CreateInvitationResponse) error {
	log.Print(req.TeamId)
	res.Hash = model.CreateInvitation(req.TeamId, req.Expired)
	return nil
}

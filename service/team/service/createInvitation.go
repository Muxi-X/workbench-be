package service

import (
	"context"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
)

//Create …… 生成邀请码
func (ts *TeamService) CreateInvitation(ctx context.Context, req *pb.CreateInvitationRequest, res *pb.CreateInvitationResponse) error {
	res.Hash = model.CreateInvitation(req.TeamId, req.Expired)
	return nil
}

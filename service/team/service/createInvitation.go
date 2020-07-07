package service

import (
	"context"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
)

func (ts *TeamService) CreateInvitation(ctx context.Context, req *pb.CreateInvitationRequest, res *pb.CreateInvitationResponse) error {
	res.Hash = model.CreateInvitation(req.TeamId, req.Expired)

	return nil
}


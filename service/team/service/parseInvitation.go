package service

import (
	"context"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
)

func (ts *TeamService) ParseInvitation(ctx context.Context, req *pb.ParseInvitationRequest, res *pb.ParseInvitationResponse) error {
	teamid, err := model.ParseInvitation(req.Hash)
	if err != nil {

	}
	res.TeamId = teamid
	return nil
}



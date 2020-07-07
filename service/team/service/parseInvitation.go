package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	"github.com/Muxi-X/workbench-be/service/team/errno"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
)

func (ts *TeamService) ParseInvitation(ctx context.Context, req *pb.ParseInvitationRequest, res *pb.ParseInvitationResponse) error {
	teamid, err := model.ParseInvitation(req.Hash)
	if err != nil {
		return e.ServerErr(errno.ErrLinkExpiration, "链接已过期")
	}
	res.TeamId = teamid
	return nil
}



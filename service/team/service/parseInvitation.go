package service

import (
	"context"
	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

//Parse …… 解析hash
func (ts *TeamService) ParseInvitation(ctx context.Context, req *pb.ParseInvitationRequest, res *pb.ParseInvitationResponse) error {
	teamid, err := model.ParseInvitation(req.Hash)
	if err != nil {
		return e.ServerErr(errno.ErrLinkExpiration, "链接已过期")
	}
	res.TeamId = teamid
	return nil
}

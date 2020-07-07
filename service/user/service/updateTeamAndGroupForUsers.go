package service

import (
	"context"

	"muxi-workbench-user/errno"
	"muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
)

func (s *UserService) UpdateTeamAndGroupForUsers(ctx context.Context, req *pb.UpdateTeamGroupRequest, res *pb.Response) error {
	if err := model.UpdateTeamAndGroup(req.Ids, req.Value, req.Kind); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

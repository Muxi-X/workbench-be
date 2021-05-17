package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// CheckProjectForUser ... 查 project 权限
func (s *Service) CheckProjectForUser(ctx context.Context, req *pb.CheckProjectRequest, res *pb.CheckProjectResponse) error {
	// 查 user2project
	var err error
	res.IfValid, err = model.CheckUser2ProjectRecord(req.UserId, req.ProjectId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

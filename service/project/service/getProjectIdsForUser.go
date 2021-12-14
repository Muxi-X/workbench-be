package service

import (
	"context"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"

	errno "muxi-workbench-project/errno"
	e "muxi-workbench/pkg/err"
)

// GetProjectIdsForUser ... 获取项目信息
func (s *Service) GetProjectIdsForUser(ctx context.Context, req *pb.GetRequest, res *pb.ProjectIdsResponse) error {

	list, err := model.GetUserToProjectByUser(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]uint32, 0)

	for _, item := range list {
		resList = append(resList, item.ProjectID)
	}

	res.List = resList

	return nil
}

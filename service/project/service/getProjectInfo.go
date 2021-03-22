package service

import (
	"context"
	"fmt"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetProjectInfo ... 获取项目信息
// 新增获取其子节点
func (s *Service) GetProjectInfo(ctx context.Context, req *pb.GetRequest, res *pb.ProjectInfo) error {
	// 判断自己 id 是否被删
	target := fmt.Sprintf("%d-%d", req.Id, constvar.ProjectCode)
	isDeleted, err := m.SIsmembersFromRedis(constvar.Trashbin, target)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted {
		return e.BadRequestErr(errno.ErrDatabase, "This file has been deleted.")
	}

	project, err := model.GetProject(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	count, err := model.GetProjectUserCount(req.Id)

	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = project.ID
	res.Name = project.Name
	res.Intro = project.Intro
	res.UserCount = count
	res.DocChildren = project.DocChildren
	res.FileChildren = project.FileChildren

	return nil
}

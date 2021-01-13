package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteProject ... 删除项目
func (s *Service) DeleteProject(ctx context.Context, req *pb.GetRequest, res *pb.ProjectNameAndIDResponse) error {
	var name string
	var err error

	// 先找到记录，再删除
	// TODO:test service
	if name, err = model.GetProjectName(req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// TODO：软删除，DB 要添加 deleted_at 字段
	if err = model.DeleteProject(req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Name = name

	return nil
}

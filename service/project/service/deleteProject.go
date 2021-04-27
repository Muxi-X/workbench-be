package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteProject ... 删除项目
// project 不需要放到回收站和 redis ，获取 project 的东西必须带 projectId
// gateway 经过验证 projectId 才能访问。
func (s *Service) DeleteProject(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	// 软删除
	if err := model.DeleteProject(req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

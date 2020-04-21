package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteProject ... 删除项目
func (s *Service) DeleteProject(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	// TODO：软删除，DB 要添加 deleted_at 字段
	if err := model.DeleteProject(req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

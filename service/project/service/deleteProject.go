package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteProject ... 删除项目 使用 gorm 软删除，直接删即可
func (s *Service) DeleteProject(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	// 软删除
	if err := model.DeleteProject(req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

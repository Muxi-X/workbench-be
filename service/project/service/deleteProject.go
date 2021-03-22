package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	e "muxi-workbench/pkg/err"
)

// DeleteProject ... 删除项目
// 插入回收站 找到所有子文件夹同步 redis
func (s *Service) DeleteProject(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	// 软删除
	// 还是要删，回收站也要建。因为有 list project，不删能被找到
	// 恢复的时候直接 update deleted_at
	if err := model.DeleteProject(m.DB.Self, req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

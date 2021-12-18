package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	e "muxi-workbench/pkg/err"
)

// UpdateTrashbin ... 从回收站移除文件
// 需要事务
// 删除回收站表 同步删除 redis 恢复文件树
func (s *Service) UpdateTrashbin(ctx context.Context, req *pb.RemoveTrashbinRequest, res *pb.Response) error {
	if err := model.RecoverTrashbin(m.DB.Self, req.Id, uint8(req.Type),
		req.IsFatherProject, req.FatherId, req.ChildrenPositionIndex, req.ProjectId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

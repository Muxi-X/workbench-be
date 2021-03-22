package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateDocChildren ... 更新任意文档夹的文档树
func (s *Service) UpdateDocChildren(ctx context.Context, req *pb.UpdateChildrenRequest, res *pb.Response) error {

	item, err := model.GetFolderForDocModel(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	item.Children = req.Children

	if err := item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

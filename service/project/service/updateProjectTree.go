package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateProjectChildren ... 更新任意文档夹的文档树
// 禁用
func (s *Service) UpdateProjectChildren(ctx context.Context, req *pb.UpdateProjectChildrenRequest, res *pb.Response) error {

	item, err := model.GetProject(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// type=0->doc type=1->file
	if req.Type {
		item.FileChildren = req.Children
	} else {
		item.DocChildren = req.Children
	}

	if err := item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

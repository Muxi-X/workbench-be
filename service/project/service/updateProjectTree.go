package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateProjectTree ... 更新任意文档夹的文档树
func (s *Service) UpdateProjectTree(ctx context.Context, req *pb.UpdateProjectTreeRequest, res *pb.Response) error {

	item, err := model.GetProject(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// type=0->doc type=1->file
	if req.Type {
		item.FileChildren = req.Tree
	} else {
		item.DocChildren = req.Tree
	}

	if err := item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

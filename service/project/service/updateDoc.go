package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateDoc ... 更新文档
func (s *Service) UpdateDoc(ctx context.Context, req *pb.UpdateDocRequest, res *pb.ProjectIDResponse) error {

	item, err := model.GetDoc(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	item.Name = req.Title
	item.Content = req.Content

	if err := item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = item.ProjectID

	return nil
}

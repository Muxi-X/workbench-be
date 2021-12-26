package service

import (
	"context"
	"time"

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

	if item.ProjectID != req.ProjectId {
		return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
	}

	item.Name = req.Name
	item.Content = req.Content
	item.EditorID = req.EditorId

	t := time.Now()
	item.LastEditTime = t.Format("2006-01-02 15:04:05")

	if err := item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = item.ProjectID

	return nil
}

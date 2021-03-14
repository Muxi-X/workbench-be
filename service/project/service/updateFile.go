package service

import (
	"context"

	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateFile ... 更新文件
func (s *Service) UpdateFile(ctx context.Context, req *pb.UpdateFileRequest, res *pb.ProjectIDResponse) error {

	item, err := model.GetFile(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	item.Name = req.Name
	item.URL = req.Url

	if err := item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = item.ProjectID

	return nil
}
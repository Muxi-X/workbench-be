package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateDocFolder ... 修改文档夹，改名
func (s *Service) UpdateDocFolder(ctx context.Context, req *pb.UpdateFolderRequest, res *pb.Response) error {
	item, err := model.GetFolderForDocModel(req.FolderId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	item.Name = req.Name

	if err = item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteDocFolder ... 删除文档，只是修改 re 字段
func (s *Service) DeleteDocFolder(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	item, err := model.GetFolderForDocModel(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	item.Re = true

	if err = item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

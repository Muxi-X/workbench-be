package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateTrashbin ... 从回收站移除文件
func (s *Service) UpdateTrashbin(ctx context.Context, req *pb.EditTrashbinRequest, res *pb.Response) error {
	var err error

	switch req.Type {
	case "0":
		err = model.RemoveProjectTrashbin(req.Id)
	case "1":
		err = model.RemoveDocTrashbin(req.Id)
	case "2":
		err = model.RemoveFileTrashbin(req.Id)
	case "3":
		err = model.RemoveDocFolderTrashbin(req.Id)
	case "4":
		err = model.RemoveFileFolderTrashbin(req.Id)
	}
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

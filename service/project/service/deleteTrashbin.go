package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteTrashbin ... 从回收站删除文件
func (s *Service) DeleteTrashbin(ctx context.Context, req *pb.EditTrashbinRequest, res *pb.Response) error {
	var err error

	switch req.Type {
	case "0":
		err = model.DeleteProjectTrashbin(req.Id)
	case "1":
		err = model.DeleteDocTrashbin(req.Id)
	case "2":
		err = model.DeleteFileTrashbin(req.Id)
	case "3":
		err = model.DeleteDocFolderTrashbin(req.Id)
	case "4":
		err = model.DeleteFileTrashbin(req.Id)
	}
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

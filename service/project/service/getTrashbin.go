package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetTrashbin ... 获取回收站资源
func (s *Service) GetTrashbin(ctx context.Context, req *pb.GetTrashbinRequest, res *pb.GetTrashbinResponse) error {
	var item []*pb.Trashbin
	var err error

	// 对 type 做判断
	switch req.Type {
	case "0":
		item, err = model.GetProjectTrashbin()
	case "1":
		item, err = model.GetDocTrashbin()
	case "2":
		item, err = model.GetFileTrashbin()
	case "3":
		item, err = model.GetDocFolderTrashbin()
	case "4":
		item, err = model.GetFileFolderTrashbin()
	}
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.List = item

	return nil
}

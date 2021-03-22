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

	list, err := model.ListTrashbin(req.Offset, req.Limit)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	for _, v := range list {
		item = append(item, &pb.Trashbin{
			Id:   v.FileId,
			Type: uint32(v.FileType),
			Name: v.Name,
		})
	}

	res.List = item

	return nil
}

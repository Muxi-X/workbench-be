package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteDoc ... 删除文档
func (s *Service) DeleteDoc(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	if err := model.DeleteDoc(req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

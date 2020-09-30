package service

import (
	"context"
	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
)

// Delete ... 删除动态
func (s *StatusService) Delete(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	err := model.DeleteStatus(req.Id, req.Uid)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

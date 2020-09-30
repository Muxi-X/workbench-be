package service

import (
	"context"
	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteComment ... 删除评论
func (s *StatusService) DeleteComment(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	err := model.DeleteComment(req.Id, req.Uid)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

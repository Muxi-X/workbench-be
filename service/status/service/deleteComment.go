package service

import (
	"context"

	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	m "muxi-workbench/model"
	e "muxi-workbench/pkg/err"
)

// DeleteComment ... 删除评论
func (s *StatusService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest, res *pb.Response) error {
	status, err := model.GetStatus(req.StatusId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	status.Comment = status.Comment - 1

	// 事务
	err = model.DeleteStatusComment(m.DB.Self, req.CommentId, req.UserId, status)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

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

	if req.Kind == 0 {
		// 事务
		err := model.DeleteStatusComment(m.DB.Self, req.CommentId, req.UserId)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	} else {
		var comment model.CommentModel
		if err := comment.Create(m.DB.Self); err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	}

	return nil
}

package service

import (
	"context"
	"muxi-workbench-status/model"

	errno "muxi-workbench-status/errno"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
)

func (s *StatusService) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest, res *pb.Response) error {
	var comment model.CommentModel
	comment.GetModelById(req.CommentId)

	// 权限判定
	if req.UserId != comment.Creator {
		return e.BadRequestErr(errno.ErrPermissionDenied, "cannot update comment not created by yourself")
	}

	if err := comment.Update(req.Content); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

func (s *Service) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest, res *pb.Response) error {
	comment, err := getCommentType(uint8(req.TypeId))
	if err != nil {
		return e.BadRequestErr(errno.ErrBind, err.Error())
	}

	if err := comment.GetModelById(req.CommentId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 权限判定
	if !comment.Verify(req.UserId) {
		return e.BadRequestErr(errno.ErrPermissionDenied, "cannot update comment not created by yourself")
	}

	if err := comment.Update(req.Content); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

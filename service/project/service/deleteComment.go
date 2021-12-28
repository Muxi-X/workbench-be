package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteComment ... 删除文档/文件评论
func (s *Service) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest, res *pb.Response) error {

	comment, err := getCommentType(uint8(req.TypeId))
	if err != nil {
		return e.BadRequestErr(errno.ErrBind, err.Error())
	}

	if err := comment.GetModelById(req.CommentId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	if err := comment.Delete(req.CommentId, req.UserId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

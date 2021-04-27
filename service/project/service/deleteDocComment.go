package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteDocComment ... 删除文档评论
func (s *Service) DeleteDocComment(ctx context.Context, req *pb.DeleteDocCommentRequest, res *pb.Response) error {
	err := model.DeleteComment(req.CommentId, req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateDocComment ... 建立文档评论
func (s *Service) CreateDocComment(ctx context.Context, req *pb.CreateDocCommentRequest, res *pb.Response) error {
	t := time.Now()

	comment := &model.CommentsModel{
		Creator: req.UserId,
		Kind:    1,
		Content: req.Content,
		Time:    t.Format("2006-01-02 15:04:05"),
		DocID:   req.DocId,
	}

	if err := comment.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

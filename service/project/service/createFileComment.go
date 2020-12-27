package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateFileComment ... 创建文件评论
func (s *Service) CreateFileComment(ctx context.Context, req *pb.CreateCommentRequest, res *pb.Response) error {

	t := time.Now()

	docComment := model.CommentsModel{
		Content: req.Content,
		Creator: req.UserId,
		DocID:   req.TargetId,
		Time:    t.Format("2006-01-02 15:04:05"),
		Kind:    1,
	}

	if err := docComment.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

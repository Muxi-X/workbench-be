package service

import (
	"context"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	"time"

	errno "muxi-workbench-status/errno"
	e "muxi-workbench/pkg/err"
)

// CreateComment ... 创建进度的评论
func (s *StatusService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest, res *pb.IdResponse) error {
	t := time.Now()

	comment := model.CommentsModel{
		Creator:  req.UserId,
		Kind:     0,
		Content:  req.Content,
		Time:     t.Format("2006-01-02 15:04:05"),
		StatusID: req.StatusId,
	}

	if err := comment.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

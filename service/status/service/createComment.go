package service

import (
	"context"
	"time"

	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	m "muxi-workbench/model"
	e "muxi-workbench/pkg/err"
)

// CreateComment ... 创建进度的评论
func (s *StatusService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest, res *pb.IdResponse) error {
	status, err := model.GetStatus(req.StatusId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	status.Comment = status.Comment + 1

	t := time.Now()

	comment := &model.CommentsModel{
		Creator:  req.UserId,
		Kind:     0,
		Content:  req.Content,
		Time:     t.Format("2006-01-02 15:04:05"),
		StatusID: req.StatusId,
	}

	// 事务
	if err := model.CreateStatusComment(m.DB.Self, comment, status); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

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
func (s *StatusService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest, res *pb.Response) error {
	t := time.Now()

	comment := &model.CommentModel{
		Creator:  req.UserId,
		Kind:     req.Kind,
		Content:  req.Content,
		Time:     t.Format("2006-01-02 15:04:05"),
		TargetID: req.TargetId,
	}

	if req.Kind == 0 {
		status, err := model.GetStatus(req.TargetId)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}

		status.Comment += 1

		// 事务
		if err := model.CreateStatusComment(m.DB.Self, comment, status); err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	} else {
		err := comment.Create(m.DB.Self)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	}
	return nil
}

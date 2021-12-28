package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateComment ... 创建文档/文件评论
func (s *Service) CreateComment(ctx context.Context, req *pb.CreateCommentRequest, res *pb.Response) error {
	t := time.Now()

	kind := req.TypeId / 3 // 1,2 代表一级评论；3,4 代表二级评论
	commentModel := &model.CommentModel{
		Creator:  req.UserId,
		Kind:     kind,
		Content:  req.Content,
		Time:     t.Format("2006-01-02 15:04:05"),
		TargetID: req.TargetId,
	}

	comment, err := getCommentType(uint8(req.TypeId))
	if err != nil {
		return e.BadRequestErr(errno.ErrBind, err.Error())
	}

	if kind == 1 {
		err := comment.GetModelById(req.TargetId)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	}

	if err := comment.Create(*commentModel); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

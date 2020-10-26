package service

import (
	"context"
	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// Create ... 创建动态
func (s *StatusService) Create(ctx context.Context, req *pb.CreateRequest, res *pb.IdResponse) error {
	t := time.Now()

	status := model.StatusModel{
		UserID:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
		Time:    t.Format("2006-01-02 15:04:05"),
	}

	if err := status.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = status.ID

	return nil
}

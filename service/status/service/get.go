package service

import (
	"context"

	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
)

// Get ... 获取动态
func (s *StatusService) Get(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {

	status, err := model.GetStatus(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Status = &pb.Status{
		Id:      status.ID,
		Title:   status.Title,
		Content: status.Content,
		UserId:  status.UserID,
		Time:    status.Time,
	}

	return nil
}

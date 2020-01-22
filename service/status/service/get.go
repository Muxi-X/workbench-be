package service

import (
	"context"
	"fmt"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
)

// Get ... 获取动态
func (s *StatusService) Get(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {

	status, err := model.GetStatus(req.Id)
	if err != nil {
		fmt.Println(err)
		return err
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

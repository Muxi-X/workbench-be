package service

import (
	"context"
	"fmt"
	"workbench-status-service/model"
	pb "workbench-status-service/proto"
)

// Update ... 更新动态
func (s *StatusService) Update(ctx context.Context, req *pb.UpdateRequest, res *pb.Response) error {

	status, err := model.GetStatus(req.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	status.Title = req.Title
	status.Content = req.Content

	if err := status.Update(); err != nil {
		return err
	}

	return nil
}

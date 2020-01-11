package service

import (
	"context"
	"fmt"
	"workbench-status-service/model"
	pb "workbench-status-service/proto"
)

// Delete ... 删除动态
func (s *StatusService) Delete(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	err := model.DeleteStatus(req.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

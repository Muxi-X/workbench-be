package service

import (
	"context"
	"fmt"
	pb "muxi-workbench-status/proto"
	"muxi-workbench-status/model"
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

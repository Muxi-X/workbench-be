package service

import (
	"context"
	"fmt"
	"workbench-status-service/model"
	pb "workbench-status-service/proto"
)

// List ... 动态列表
func (s *StatusService) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {

	list, count, err := model.ListStatus(0, 0, 20, 0)
	if err != nil {
		fmt.Println(err)
		return err
	}

	resList := make([]*pb.Status, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.Status{
			Id:      item.ID,
			Content: item.Content,
			Title:   item.Title,
			Time:    item.Time,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

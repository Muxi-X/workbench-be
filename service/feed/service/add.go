package service

import (
	"context"
	"fmt"

	"muxi-workbench-feed/model"
	logger "muxi-workbench/log"

	pb "muxi-workbench-feed/proto"
)

func (s *StatusService) Add(ctx context.Context, req *pb.AddRequest, res *pb.AddResponse) error {
	err := model.PubRdb.Publish(model.RdbChan, req).Err()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *StatusService) SubForDB() error {
	sub := model.SubRdb.Subscribe(model.RdbChan)
	ch := sub.Channel()
	for msg := range ch {
		fmt.Println(msg.Payload)
		// DB写入
	}

	return nil
}

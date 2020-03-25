package service

import (
	"context"

	"muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	logger "muxi-workbench/log"
)

func (s *FeedService) Add(ctx context.Context, req *pb.AddRequest, res *pb.AddResponse) error {
	err := model.PubRdb.Publish(model.RdbChan, req).Err()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

package service

import (
	"context"
	"encoding/json"
	"time"

	"muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	logger "muxi-workbench/log"
)

// 新增feed
func (s *FeedService) Add(ctx context.Context, req *pb.AddRequest, res *pb.Response) error {
	feed := &model.FeedModel{
		UserId:            req.User.Id,
		Username:          req.User.Name,
		UserAvatar:        req.User.AvatarUrl,
		Action:            req.Action,
		SourceKindId:      req.Source.KindId,
		SourceObjectName:  req.Source.ObjectName,
		SourceObjectId:    req.Source.ObjectId,
		SourceProjectName: req.Source.ProjectName,
		SourceProjectId:   req.Source.ProjectId,
		TimeDay:           time.Now().Format("2006/01/02"),
		TimeHm:            time.Now().Format("15:04"),
	}

	msg, err := json.Marshal(feed)

	err = model.PubRdb.Publish(model.RdbChan, msg).Err()
	if err != nil {
		logger.Error("Publish data error")
		return err
	}
	logger.Info("Publish data OK")
	return nil
}

package service

import (
	"context"
	"encoding/json"
	"time"

	"muxi-workbench-feed/errno"
	"muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	logger "muxi-workbench/log"
	e "muxi-workbench/pkg/err"
)

// Push ... 异步新增feed
func (s *FeedService) Push(ctx context.Context, req *pb.PushRequest, res *pb.Response) error {
	// get username and avatar by userId from user-service
	username, avatar, err := GetInfoFromUserService(req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
	}

	feed := &model.FeedModel{
		UserId:            req.UserId,
		Username:          username,
		UserAvatar:        avatar,
		Action:            req.Action,
		SourceKindId:      req.Source.Kind,
		SourceObjectName:  req.Source.Name,
		SourceObjectId:    req.Source.Id,
		SourceProjectName: req.Source.ProjectName,
		SourceProjectId:   req.Source.ProjectId,
		TimeDay:           time.Now().Format("2006/01/02"),
		TimeHm:            time.Now().Format("15:04"),
	}

	msg, err := json.Marshal(feed)
	if err != nil {
		return e.ServerErr(errno.ErrJsonMarshal, err.Error())
	}

	if err := model.PublishMsg(msg); err != nil {
		logger.Error("Publish data error")
		return e.ServerErr(errno.ErrPublishMsg, err.Error())
	}
	logger.Info("Publish data OK")
	return nil
}

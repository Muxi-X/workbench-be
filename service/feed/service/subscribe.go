package service

import (
	"time"

	"muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	logger "muxi-workbench/log"
)

func (s *SubService) SubForDB() error {
	//sub := model.SubRdb.Subscribe(model.RdbChan)
	//ch := sub.Channel()
	//for msg := range ch {
	//	fmt.Println(msg.Payload)
	//	// DB写入
	//}
	msg, err := model.SubRdb.Receive()
	if err != nil {
		return err
	}
	req := msg.(pb.AddRequest)

	return FeedDBWrite(&req)

	//for {
	//	msg, err := model.SubRdb.Receive()
	//	if err != nil {
	//		return err
	//	}
	//	req := msg.(pb.AddRequest)
	//
	//	go FeedDBWrite(&req)
	//}
}

func FeedDBWrite(req *pb.AddRequest) error {

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
		TimeHm:            time.Now().Format("15:04:05"),
	}

	if err := feed.Create(); err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

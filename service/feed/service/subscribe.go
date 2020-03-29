package service

import (
	"fmt"
	"time"

	"muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	logger "muxi-workbench/log"
)

func SubServiceRun() {
	//sub := model.SubRdb.Subscribe(model.RdbChan)
	//ch := sub.Channel()
	//for msg := range ch {
	//	fmt.Println(msg.Payload)
	//	// DB写入
	//}

	for {
		msg, err := model.SubRdb.Receive()
		if err != nil {
			logger.Error("sub receive error")
			continue
		}
		req := msg.(pb.AddRequest)

		if err := FeedDBWrite(&req); err != nil {
			logger.Error("feed db write error")
		}
	}
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

	fmt.Println(feed)

	//if err := feed.Create(); err != nil {
	//	logger.Error(err.Error())
	//	return err
	//}
	return nil
}

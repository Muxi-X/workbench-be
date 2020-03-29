package service

import (
	"encoding/json"

	"muxi-workbench-feed/model"
	logger "muxi-workbench/log"
)

func SubServiceRun() {
	var feed = &model.FeedModel{}

	ch := model.SubRdb.Channel()
	for msg := range ch {
		logger.Info("received")

		if err := json.Unmarshal([]byte(msg.Payload), feed); err != nil {
			panic(err)
		}
		//fmt.Println(feed)

		if err := feed.Create(); err != nil {
			logger.Error(err.Error())
		}
	}

	//for {
	//	msg, err := model.SubRdb.ReceiveMessage()
	//	if err != nil {
	//		logger.Error("sub receive error")
	//		continue
	//	}
	//	//logger.Info("received")
	//	if err := json.Unmarshal([]byte(msg.Payload), feed); err != nil {
	//		panic(err)
	//	}
	//
	//	//fmt.Println(feed)
	//	if err := feed.Create(); err != nil {
	//		logger.Error(err.Error())
	//	}
	//}
}

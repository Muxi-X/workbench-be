package service

import (
	"encoding/json"
	"muxi-workbench-feed/model"
	logger "muxi-workbench/log"
	m "muxi-workbench/model"
)

// SubServiceRun ... 写入feed数据
func SubServiceRun() {
	var feed = &model.FeedModel{}

	ch := m.PubSubClient.Self.Channel()
	for msg := range ch {
		logger.Info("received")

		if err := json.Unmarshal([]byte(msg.Payload), feed); err != nil {
			panic(err)
		}

		if err := feed.Create(); err != nil {
			logger.Error(err.Error())
		}
	}
}

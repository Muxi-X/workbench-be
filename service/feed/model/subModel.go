package model

import (
	m "muxi-workbench/model"
)

//var PubRdb *redis.Client
//var SubRdb *redis.PubSub

const RdbChan = "sub"

func PublishMsg(msg []byte) error {
	return m.Rdb.Self.Publish(RdbChan, msg).Err()
}
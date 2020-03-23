package model

import (
	. "muxi-workbench-feed/proto"
	m "muxi-workbench/model"
)

type FeedModel struct {
	Id                uint32 `json:"id" gorm:"id"`
	UserId            uint32 `json:"user_id" gorm:"column:userid"`
	Username          string `json:"username" gorm:"column:username"`
	UserAvatar        string `json:"user_avatar" gorm:"useravatar"`
	Action            string `json:"action" gorm:"action"`
	SourceKindId      uint32 `json:"source_kind_id" gorm:"source_kindid"`
	SourceObjectName  string `json:"source_object_name" gorm:"source_objectname"`
	SourceObjectId    uint32 `json:"source_object_id" gorm:"source_objectid"`
	SourceProjectName string `json:"source_project_name" gorm:"source_projectname"`
	SourceProjectId   int32  `json:"source_project_id" gorm:"source_projectid"` // 存在-1
	TimeDay           string `json:"time_day" gorm:"timeday"`
	TimeHm            string `json:"time_hm" gorm:"timehm"`
}

func (*FeedModel) TableName() string {
	return "feeds"
}

func (f *FeedModel) Create() error {
	return m.DB.Self.Create(f).Error
}

func GetRowsSum() (uint32, error) {
	var count uint32
	d := m.DB.Self.Table("feeds").Count(count)
	return count, d.Error
}

func FeedList(offset, limit uint32) ([]*SingleData, error) {
	//var data []*FeedModel
	var result []*SingleData

	//d := m.DB.Self.Find(&data).Order("id desc").Offset(offset).Limit(limit)
	d := m.DB.Self.Table("feeds").Order("id dsc").Offset(offset).Limit(limit).Scan(result)

	return result, d.Error
}

func PersonalFeedList(uid, offset, limit uint32) ([]*SingleData, error) {
	var result []*SingleData
	return result, nil
}

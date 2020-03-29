package model

import (
	"github.com/jinzhu/gorm"
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
	d := m.DB.Self.Table("feeds").Count(&count)
	return count, d.Error
}

func GetPersonalRowsSum(uid uint32) (uint32, error) {
	var count uint32
	d := m.DB.Self.Table("feeds").Where("userid = ?", uid).Count(&count)
	return count, d.Error
}

func FeedList(lastId, limit uint32) ([]*FeedModel, error) {
	var data []*FeedModel
	var d *gorm.DB

	// 判断是否为0, 0为第一次查询
	if lastId != 0 {
		d = m.DB.Self.Where("id < ?", lastId).Order("id desc").Limit(limit).Find(&data)
	} else {
		d = m.DB.Self.Order("id desc").Limit(limit).Find(&data)
	}

	if d.RecordNotFound() {
		return data, nil
	}
	return data, d.Error
}

func PersonalFeedList(uid, lastId, limit uint32) ([]*FeedModel, error) {
	var data []*FeedModel
	var d *gorm.DB

	// 判断是否为0, 0为第一次查询
	if lastId != 0 {
		d = m.DB.Self.Where("userid = ? AND id < ?", uid, lastId).Order("id desc").Limit(limit).Find(&data)
	} else {
		d = m.DB.Self.Where("userid = ?", uid).Order("id desc").Limit(limit).Find(&data)
	}

	if d.RecordNotFound() {
		return data, nil
	}

	return data, d.Error
}

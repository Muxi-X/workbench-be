package model

import (
	"github.com/jinzhu/gorm"
	m "muxi-workbench/model"
)

type FeedModel struct {
	Id                uint32 `json:"id" gorm:"column:id"`
	UserId            uint32 `json:"user_id" gorm:"column:userid"`
	Username          string `json:"username" gorm:"column:username"`
	UserAvatar        string `json:"user_avatar" gorm:"column:useravatar"`
	Action            string `json:"action" gorm:"column:action"`
	SourceKindId      uint32 `json:"source_kind_id" gorm:"column:source_kindid"`
	SourceObjectName  string `json:"source_object_name" gorm:"column:source_objectname"`
	SourceObjectId    uint32 `json:"source_object_id" gorm:"column:source_objectid"`
	SourceProjectName string `json:"source_project_name" gorm:"column:source_projectname"`
	SourceProjectId   int32  `json:"source_project_id" gorm:"column:source_projectid"` // 存在-1
	TimeDay           string `json:"time_day" gorm:"column:timeday"`
	TimeHm            string `json:"time_hm" gorm:"column:timehm"`
}

func (*FeedModel) TableName() string {
	return "feeds"
}

// Create a new feed
func (f *FeedModel) Create() error {
	return m.DB.Self.Create(f).Error
}

// 获取全部feed数量
func GetRowsSum() (uint32, error) {
	var count uint32
	d := m.DB.Self.Table("feeds").Count(&count)
	return count, d.Error
}

// 获取个人feed总数
func GetPersonalRowsSum(uid uint32) (uint32, error) {
	var count uint32
	d := m.DB.Self.Table("feeds").Where("userid = ?", uid).Count(&count)
	return count, d.Error
}

// 查找feed
// 参数：last-->上一次查询的最后一个feed的id；limit-->预期返回的数据数
func GetFeedList(lastId, limit uint32) ([]*FeedModel, error) {
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

// 查找个人feed
// 参数：uid-->user id；last-->上一次查询的最后一个feed的id；limit-->预期返回的数据数
func GetPersonalFeedList(uid, lastId, limit uint32) ([]*FeedModel, error) {
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

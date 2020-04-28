package model

import (
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
	SourceProjectId   uint32 `json:"source_project_id" gorm:"column:source_projectid"`
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

// GetFeedList ... 查找feed
func GetFeedList(lastId, limit, uid uint32, filter []uint32, filterRequired bool) ([]*FeedModel, error) {
	var data []*FeedModel

	query := m.DB.Self.Table("feeds").Order("id desc").Limit(limit)

	// 查找用户的feed
	if uid != 0 {
		query = query.Where("userid = ?", uid)
	}

	// 0为第一次查询
	if lastId != 0 {
		query = query.Where("id < ?", lastId)
	}

	if filterRequired {
		query = query.Where("source_projectid in (?) OR source_projectid = 0", filter)
	}

	if err := query.Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// GetPersonalFeedList ... 查找用户所属feed
//func GetPersonalFeedList(lastId, limit, uid uint32) ([]*FeedModel, uint32, error) {
//	var data []*FeedModel
//	var count uint32
//
//	query := m.DB.Self.Table("feeds").Where("userid = ?", uid).Order("id desc").Limit(limit)
//
//	// 0为第一次查询
//	if lastId != 0 {
//		query = query.Where("id < ?", lastId)
//	}
//
//	if err := query.Scan(&data).Count(&count).Error; err != nil {
//		return nil, 0, err
//	}
//
//	return data, count, nil
//}

// GetFeedListForGroup ... 根据组别查找feed
func GetFeedListForGroup(lastId, limit, groupId uint32, filter []uint32, filterRequired bool) ([]*FeedModel, error) {
	var data []*FeedModel
	//var count uint32

	query := m.DB.Self.Table("feeds").
		Select("feeds.*").
		Where("users.group_id = ?", groupId).
		Joins("left join users on users.id = feeds.userid").
		Order("feeds.id desc").
		Limit(limit)

	if lastId != 0 {
		query = query.Where("feeds.id < ?", lastId)
	}

	if filterRequired {
		query = query.Where("source_projectid in (?) OR source_projectid = 0", filter)
	}

	if err := query.Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

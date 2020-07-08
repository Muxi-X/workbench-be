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

// FilterParams provide filter's params.
type FilterParams struct {
	UserId     uint32
	GroupId    uint32
	ProjectIds []uint32 // 可访问的 projects' id
}

// GetFeedList ... 查找feed
func GetFeedList(lastId, limit uint32, filter *FilterParams) ([]*FeedModel, error) {
	var data []*FeedModel

	query := m.DB.Self.Table("feeds").Select("feeds.*").Order("feeds.id desc").Limit(limit)

	// 查找用户的feed
	if filter.UserId != 0 {
		query = query.Where("feeds.userid = ?", filter.UserId)
	}

	// 组别筛选
	if filter.GroupId != 0 {
		query = query.Where("users.group_id = ?", filter.GroupId).Joins("left join users on users.id = feeds.userid")
	}

	// 项目权限过滤
	if len(filter.ProjectIds) != 0 {
		query = query.Where("source_projectid in (?) OR source_projectid = 0", filter.ProjectIds)
	}

	// 分页
	if lastId != 0 {
		query = query.Where("feeds.id < ?", lastId)
	}

	if err := query.Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

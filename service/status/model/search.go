package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type SearchResult struct {
	Id       uint32 `json:"id"`
	Title    string `json:"title"`
	UserName string `json:"user_name" gorm:"column:name;not null"`
	Content  string `json:"content"`
	Time     string `json:"time"`
}

func Search(filter *FilterParams, offset, limit, lastID uint32, pagination bool) ([]*SearchResult, uint32, error) {
	var count uint32
	var record []*SearchResult
	query := m.DB.Self.Table("status").Select("status.id, title, time, u.name, content").
		Joins("LEFT JOIN users u ON u.id = status.user_id").
		Where("title LIKE ? OR content LIKE ?", filter.Key, filter.Key)

	if filter.UserName != "" {
		query = query.Where("u.name LIKE ?", filter.UserName)
	}

	if filter.GroupId != 0 {
		query = query.Where("u.group_id = ?", filter.GroupId)
	}

	if pagination {
		if limit == 0 {
			limit = constvar.DefaultLimit
		}
		query = query.Offset(offset).Limit(limit)

		if lastID != 0 {
			query = query.Where("status.id < ?", lastID)
		}
	}

	err := query.Scan(&record).Count(&count).Error

	return record, count, err
}

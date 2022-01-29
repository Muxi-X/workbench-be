package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type SearchResult struct {
	Id          uint32 `json:"id"`
	Type        uint8  `json:"type"`
	Title       string `json:"title" gorm:"column:filename;not null"`
	UserName    string `json:"user_name" gorm:"column:name;not null"`
	Content     string `json:"content"`
	ProjectName string `json:"project_name" gorm:"column:project_name;not null"`
	Time        string `json:"time" gorm:"column:create_time"`
}

func SearchDoc(projectIDs []uint32, keyword string, offset, limit, lastID uint32, pagination bool) ([]*SearchResult, uint32, error) {
	var count uint32
	var record []*SearchResult
	key := "%" + keyword + "%"
	query := m.DB.Self.Table("docs").Select("docs.id, filename, create_time, u.name, content, p.name project_name").
		Joins("LEFT JOIN users u ON u.id = docs.editor_id").
		Joins("LEFT JOIN projects p ON p.id = project_id").
		Where("re = 0 AND project_id IN (?) AND (filename LIKE ? OR content LIKE ?)", projectIDs, key, key)

	if pagination {
		if limit == 0 {
			limit = constvar.DefaultLimit
		}
		query = query.Offset(offset).Limit(limit)

		if lastID != 0 {
			query = query.Where("docs.id < ?", lastID)
		}
	}

	for _, r := range record {
		r.Type = constvar.DocCode
	}

	err := query.Scan(&record).Count(&count).Error

	return record, count, err
}

func SearchFile(projectIDs []uint32, keyword string, offset, limit, lastID uint32, pagination bool) ([]*SearchResult, uint32, error) {
	var count uint32
	var record []*SearchResult
	key := "%" + keyword + "%"
	query := m.DB.Self.Table("files").Select("files.id, realname filename, create_time, u.name, url content, p.name project_name").
		Joins("LEFT JOIN users u ON u.id = files.creator_id").
		Joins("LEFT JOIN projects p ON p.id = project_id").
		Where("re = 0 AND project_id IN (?) AND realname LIKE ?", projectIDs, key)

	if pagination {
		if limit == 0 {
			limit = constvar.DefaultLimit
		}
		query = query.Offset(offset).Limit(limit)

		if lastID != 0 {
			query = query.Where("files.id < ?", lastID)
		}
	}

	for _, r := range record {
		r.Type = constvar.FileCode
	}

	err := query.Scan(&record).Count(&count).Error

	return record, count, err
}

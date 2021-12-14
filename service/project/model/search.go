package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type SearchResult struct {
	Id uint32 `json:"id"`
	// Type        uint32 `json:"type"`
	Title string `json:"title" gorm:"column:filename;not null"`
	// Url         string `json:"url"`
	UserName    string `json:"user_name" gorm:"column:name;not null"`
	Content     string `json:"content"`
	ProjectName string `json:"project_name" gorm:"column:project_name;not null"`
	Time        string `json:"time" gorm:"column:create_time"`
}

func SearchTitle(projectIDs []uint32, keyword string, offset, limit, lastID uint32, pagination bool) ([]*SearchResult, uint32, error) {
	var count uint32
	var record []*SearchResult
	key := "%" + keyword + "%"
	query := m.DB.Self.
		Raw("SELECT d.id, filename, create_time, u.name, content, p.name project_name FROM docs d "+
			"LEFT JOIN users u ON u.id = d.editor_id "+
			"LEFT JOIN projects p ON p.id = project_id "+
			"WHERE project_id in (?) AND d.filename like ? "+
			"UNION ALL SELECT f.id, filename, create_time, u.name, url, p.name project_name FROM files f "+
			"LEFT JOIN users u ON u.id = f.creator_id "+
			"LEFT JOIN projects p ON p.id = project_id "+
			"WHERE project_id in (?) AND f.filename like ? ", projectIDs, key, projectIDs, key)

	if pagination {
		if limit == 0 {
			limit = constvar.DefaultLimit
		}
		query = query.Offset(offset).Limit(limit)

		// if lastID != 0 {
		// 	query = query.Where("projects.id < ?", lastID) // TODO
		// }
	}

	err := query.Scan(&record).Count(&count).Error

	return record, count, err
}

func SearchContent(projectIDs []uint32, keyword string, offset, limit, lastID uint32, pagination bool) ([]*SearchResult, uint32, error) {
	var count uint32
	var record []*SearchResult
	key := "%" + keyword + "%"
	query := m.DB.Self.
		Raw("SELECT d.id, filename, create_time, u.name, content, p.name project_name FROM docs d "+
			"LEFT JOIN users u ON u.id = d.editor_id "+
			"LEFT JOIN projects p ON p.id = project_id "+
			"WHERE project_id in (?) AND d.filename like ? "+
			"UNION ALL SELECT f.id, filename, create_time, u.name, url, p.name project_name FROM files f "+
			"LEFT JOIN users u ON u.id = f.creator_id "+
			"LEFT JOIN projects p ON p.id = project_id "+
			"WHERE project_id in (?) AND f.filename like ? ", projectIDs, key, projectIDs, key)

	if pagination {
		if limit == 0 {
			limit = constvar.DefaultLimit
		}
		query = query.Offset(offset).Limit(limit)

		// if lastID != 0 {
		// 	query = query.Where("projects.id < ?", lastID) // TODO
		// }
	}

	err := query.Scan(&record).Count(&count).Error

	return record, count, err
}

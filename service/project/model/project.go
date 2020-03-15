package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type ProjectModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name     string `json:"name" gorm:"column:name;" binding:"required"`
	Intro    string `json:"intro" gorm:"column:intro;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Count    uint32 `json:"count" gorm:"column:count;" binding:"required"`
	TeamID   uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	FileTree string `json:"fileTree" gorm:"column:fileTree;" binding:"required"`
	DocTree  string `json:"docTree" gorm:"column:docTree;" binding:"required"`
}

type ProjectListItem struct {
	ID   uint32 `json:"id" gorm:"column:project_id;not null" binding:"required"`
	Name string `json:"name" gorm:"column:name;" binding:"required"`
}

func (u *ProjectModel) TableName() string {
	return "projects"
}

// ListProject list all project
func ListProject(userID, offset, limit, lastID uint32, pagination bool) ([]*ProjectListItem, uint64, error) {
	var count uint64

	projectList := make([]*ProjectListItem, 0)

	query := m.DB.Self.Table("user2projects").Where("user2projects.user_id = ?", userID).Select("user2projects.*, projects.name").Joins("left join projects on user2projects.project_id = projects.id").Order("projects.id")

	if pagination {
		if limit == 0 {
			limit = constvar.DefaultLimit
		}

		query = query.Offset(offset).Limit(limit).Count(&count)

		if lastID != 0 {
			query = query.Where("projects.id < ?", lastID)
		}
	}

	if err := query.Scan(&projectList).Error; err != nil {
		return projectList, count, err
	}

	return projectList, count, nil
}

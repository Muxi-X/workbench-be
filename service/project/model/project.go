package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

// ProjectModel project table's structure
type ProjectModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name     string `json:"name" gorm:"column:name;" binding:"required"`
	Intro    string `json:"intro" gorm:"column:intro;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Count    uint32 `json:"count" gorm:"column:count;" binding:"required"`
	TeamID   uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	FileTree string `json:"fileTree" gorm:"column:fileTree;" binding:"required"`
	DocTree  string `json:"docTree" gorm:"column:docTree;" binding:"required"`
	Children string `json:"children" gorm:"column:children;" binding:"required"`
}

// ProjectListItem ProjectList service item
type ProjectListItem struct {
	ID   uint32 `json:"id" gorm:"column:project_id;not null" binding:"required"`
	Name string `json:"name" gorm:"column:name;" binding:"required"`
}

// ProjectName save the name of project
type ProjectName struct {
	Name string `json:"name" gorm:"column:name;" binding:"required"`
}

// ChildrenItem ... 数据库提取出来的 children 经过序列化形成结构体切片，此为结构体
type ChildrenItem struct {
	ID   uint32 `json:"id"`
	Type int    `json:"type"` // 类型标志 0->project 1->文档夹 2->文件夹 3->文档 4->文件
}

// TableName return table name
func (u *ProjectModel) TableName() string {
	return "projects"
}

// Create ... 创建项目
func (u *ProjectModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// DeleteProject ... 删除项目
func DeleteProject(id uint32) error {
	doc := &ProjectModel{}
	doc.ID = id
	return m.DB.Self.Delete(&doc).Error
}

// GetProjectName return project's name
func GetProjectName(id uint32) (string, error) {
	record := &ProjectName{}
	err := m.DB.Self.Table("projects").Where("id=?", id).Select("name").First(record).Error
	return record.Name, err
}

// Update ... 更新项目
func (u *ProjectModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetProject ... 获取项目
func GetProject(id uint32) (*ProjectModel, error) {
	s := &ProjectModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
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

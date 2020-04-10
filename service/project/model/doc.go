package model

import (
	m "muxi-workbench/model"
)

type DocModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	Content    string `json:"content" gorm:"column:content;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	Top        bool   `json:"top" gorm:"column:top;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	DeleteTime string `json:"deleteTime" gorm:"column:delete_time;" binding:"required"`
	TeamID     uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	EditorID   uint32 `json:"editorId" gorm:"column:editor_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

func (u *DocModel) TableName() string {
	return "docs"
}

// Create doc
func (u *DocModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// DeleteDoc ... 删除文档
func DeleteDoc(id uint32) error {
	doc := &DocModel{}
	doc.ID = id
	return m.DB.Self.Delete(&doc).Error
}

// Update doc
func (u *DocModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetDoc ... 获取文档
func GetDoc(id uint32) (*DocModel, error) {
	s := &DocModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

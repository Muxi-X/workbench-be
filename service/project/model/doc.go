package model

import (
	m "muxi-workbench/model"
)

// DocDetail ... 文档详情
type DocDetail struct {
	Creator string `json:"creator" gorm:"column:creator;not null" binding:"required"`
	Editor  string `json:"editor" gorm:"column:editor;" binding:"required"`
	DocModel
}

// DocInfo ... 文档信息
type DocInfo struct {
	ID   uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name string `json:"name" gorm:"column:filename;" binding:"required"`
}

// DocModel ... 文档物理模型
type DocModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	Content    string `json:"content" gorm:"column:content;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	Top        bool   `json:"top" gorm:"column:top;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	DeleteTime string `json:"deleteTime" gorm:"column:delete_time;" binding:"required"`
	TeamID     uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	CreatorID  uint32 `json:"creatorId" gorm:"column:creator_id;" binding:"required"`
	EditorID   uint32 `json:"editorId" gorm:"column:editor_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

// TableName ... 物理表名
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

// GetDocInfo ... 获取文档信息
func GetDocInfo(id uint32) (*DocInfo, error) {
	s := &DocModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	info := &DocInfo{}
	info.ID = s.ID
	info.Name = s.Name
	return info, d.Error
}

// GetDocInfoByIds ... 获取文档信息列表
func GetDocInfoByIds(ids []uint32) ([]*DocInfo, error) {
	s := make([]*DocInfo, 0)
	d := m.DB.Self.Table("docs").Where("id IN (?)", ids).Find(&s)
	return s, d.Error
}

// GetDocDetail ... 获取文档详情
func GetDocDetail(id uint32) (*DocDetail, error) {
	s := &DocDetail{}
	// multiple left join
	d := m.DB.Self.Table("docs").Where("docs.id = ?", id).Select("docs.*, c.name as creator, e.name as editor").Joins("left join users c on c.id = docs.creator_id").Joins("left join users e on e.id = docs.editor_id").First(&s)
	return s, d.Error
}

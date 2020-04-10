package model

import m "muxi-workbench/model"

// FolderForDocModel ... 文档文件夹模型
type FolderForDocModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	CreatorID  string `json:"creatorID" gorm:"column:create_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

// TableName ... 物理表名
func (u *FolderForDocModel) TableName() string {
	return "foldersformds"
}

// Create ... 创建文档文件夹
func (u *FolderForDocModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// DeleteFolderForDoc ... 删除文档文件夹
func DeleteFolderForDoc(id uint32) error {
	doc := &FolderForDocModel{}
	doc.ID = id
	return m.DB.Self.Delete(&doc).Error
}

// Update ... 更新文档文件夹
func (u *FolderForDocModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetFolderForDocModel ... 获取文档文件夹
func GetFolderForDocModel(id uint32) (*FolderForDocModel, error) {
	s := &FolderForDocModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

package model

import m "muxi-workbench/model"

// FolderForFileModel ... 文件文件夹模型
type FolderForFileModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	CreatorID  string `json:"creatorID" gorm:"column:create_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

// TableName ... 物理表名
func (u *FolderForFileModel) TableName() string {
	return "foldersforfiles"
}

// Create ... 创建文件文件夹
func (u *FolderForFileModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// DeleteFolderForFile ... 删除文件文件夹
func DeleteFolderForFile(id uint32) error {
	doc := &FolderForFileModel{}
	doc.ID = id
	return m.DB.Self.Delete(&doc).Error
}

// Update ... 更新文件文件夹
func (u *FolderForFileModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetFolderForFileModel ... 获取文件文件夹
func GetFolderForFileModel(id uint32) (*FolderForFileModel, error) {
	s := &FolderForFileModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

package model

import m "muxi-workbench/model"

type FileModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	RealName   string `json:"realName" gorm:"column:realname;" binding:"required"`
	URL        string `json:"url" gorm:"column:url;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	Top        bool   `json:"top" gorm:"column:top;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	DeleteTime string `json:"deleteTime" gorm:"column:delete_time;" binding:"required"`
	TeamID     uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

func (u *FileModel) TableName() string {
	return "files"
}

// 创建文件
func (u *FileModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// DeleteFile ... 删除文件
func DeleteFile(id uint32) error {
	doc := &FileModel{}
	doc.ID = id
	return m.DB.Self.Delete(&doc).Error
}

// Update doc
func (u *FileModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetFile ... 获取文件
func GetFile(id uint32) (*FileModel, error) {
	s := &FileModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

package model

import (
	"github.com/jinzhu/gorm"
	m "muxi-workbench/model"
)

// FileDetail ... 文件/文档详情
type FileDetail struct {
	Creator      string `json:"creator" gorm:"column:creator;" binding:"required"`
	Editor       string `json:"editor" gorm:"column:editor;" binding:"required"`
	ID           uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string `json:"name" gorm:"column:filename;" binding:"required"`
	Content      string `json:"content" gorm:"column:content;" binding:"required"`
	URL          string `json:"url" gorm:"column:url;" binding:"required"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	ProjectID    uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	LastEditTime string `json:"lastEditTime" gorm:"column:last_edit_time;" binding:"required"`
}

// FileInfo ... 文件/文档信息
type FileInfo struct {
	ID        uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name      string `json:"name" gorm:"column:name;" binding:"required"`
	ProjectID uint32 `json:"project_id" gorm:"column:project_id;" binding:"required"`
}

// FileModel ... 文件物理模型
type FileModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	RealName   string `json:"realName" gorm:"column:realname;" binding:"required"`
	URL        string `json:"url" gorm:"column:url;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	Top        bool   `json:"top" gorm:"column:top;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	DeleteTime string `json:"deleteTime" gorm:"column:delete_time;" binding:"required"`
	CreatorID  uint32 `json:"creatorId" gorm:"column:creator_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	FatherId   uint32 `json:"father_id" gorm:"column:father_id;" binding:"required"`
}

// TableName ... 物理表名
func (u *FileModel) TableName() string {
	return "files"
}

// Create ... 创建文件
func (u *FileModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// Update ... 更新文件
func (u *FileModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetFileInfoByIds ... 获取文件信息列表
func GetFileInfoByIds(ids []uint32) ([]*FileInfo, error) {
	s := make([]*FileInfo, 0)
	d := m.DB.Self.Table("files").Select("id, project_id, filename name").Where("id IN (?) AND re = 0", ids).Find(&s)
	return s, d.Error
}

// GetFileDetail ... 获取文件详情
func GetFileDetail(id uint32) (*FileDetail, error) {
	s := &FileDetail{}
	d := m.DB.Self.Table("files").Where("files.id = ? AND re = 0", id).Select("realname filename, files.*, users.name as creator").Joins("left join users on users.id = files.creator_id").First(&s)
	return s, d.Error
}

func CreateFile(db *gorm.DB, file *FileModel, childrenPositionIndex uint32) (uint32, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(file).Error; err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	// 获取 fatherId
	fatherId := file.FatherId
	isFatherProject := false
	if file.FatherId == 0 {
		isFatherProject = true
		fatherId = file.ProjectID
	}

	if err := AddChildren(tx, isFatherProject, fatherId, childrenPositionIndex, file); err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	return file.ID, tx.Commit().Error
}

func GetFile(id uint32) (*FileModel, error) {
	s := &FileModel{}
	d := m.DB.Self.Where("id = ? AND re = 0", id).First(&s)
	return s, d.Error
}

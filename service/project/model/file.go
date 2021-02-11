package model

import (
	"fmt"
	m "muxi-workbench/model"

	"github.com/jinzhu/gorm"
)

// FileDetail ... 文件详情
type FileDetail struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:realname;" binding:"required"`
	URL        string `json:"url" gorm:"column:url;" binding:"required"`
	Creator    string `json:"creator" gorm:"column:creator;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
}

// FileInfo ... 文件信息
type FileInfo struct {
	ID   uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name string `json:"name" gorm:"column:realname;" binding:"required"`
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
	TeamID     uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	CreatorID  uint32 `json:"creatorId" gorm:"column:creator_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

// TableName ... 物理表名
func (u *FileModel) TableName() string {
	return "files"
}

// Create ... 创建文件
func (u *FileModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// DeleteFile ... 删除文件
func DeleteFile(id uint32) error {
	doc := &FileModel{}
	doc.ID = id
	return m.DB.Self.Delete(&doc).Error
}

// Update ... 更新文件
func (u *FileModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetFileInfoByIds ... 获取文件信息列表
func GetFileInfoByIds(ids []uint32) ([]*FileInfo, error) {
	s := make([]*FileInfo, 0)
	d := m.DB.Self.Table("files").Where("id IN (?)", ids).Find(&s)
	return s, d.Error
}

// GetFileDetail ... 获取文件详情
func GetFileDetail(id uint32) (*FileDetail, error) {
	s := &FileDetail{}
	d := m.DB.Self.Table("files").Where("files.id = ?", id).Select("files.*, users.name as creator").Joins("left join users on users.id = files.creator_id").First(&s)
	return s, d.Error
}

func CreateFile(db *gorm.DB, file *FileModel, fatherId uint32, fatherType bool) (uint32, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := file.Create(); err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	if fatherType {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			tx.Rollback()
			return uint32(0), err
		}
		item.FileChildren = fmt.Sprintf("%s,%d-%d", item.FileChildren, file.ID, 0)
		if err := item.Update(); err != nil {
			tx.Rollback()
			return uint32(0), err
		}
	} else {
		item, err := GetFolderForFileModel(fatherId)
		if err != nil {
			tx.Rollback()
			return uint32(0), err
		}
		item.Children = fmt.Sprintf("%s,%d-%d", item.Children, file.ID, 0)
		if err := item.Update(); err != nil {
			tx.Rollback()
			return uint32(0), err
		}
	}

	return file.ID, tx.Commit().Error
}

func GetFile(id uint32) (*FileModel, error) {
	s := &FileModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

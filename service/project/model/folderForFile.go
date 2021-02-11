package model

import (
	"fmt"
	m "muxi-workbench/model"

	"github.com/jinzhu/gorm"
)

// FolderForFileInfo ... 文件文件夹信息
type FolderForFileInfo struct {
	ID   uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name string `json:"name" gorm:"column:filename;" binding:"required"`
}

type FolderForFileChildren struct {
	Children string `json:"children" gorm:"column:children;" binding:"required"`
}

// FolderForFileModel ... 文件文件夹模型
type FolderForFileModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	CreatorID  uint32 `json:"creatorID" gorm:"column:create_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	Children   string `json:"children" gorm:"column:children;" binding:"required"`
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

// GetFolderForFileInfoByIds ... 获取文件文件夹信息列表
func GetFolderForFileInfoByIds(ids []uint32) ([]*FolderForFileInfo, error) {
	s := make([]*FolderForFileInfo, 0)
	d := m.DB.Self.Table("foldersforfiles").Where("id IN (?)", ids).Find(&s)
	return s, d.Error
}

// GetFileChildrenById ... 获取子文件
func GetFileChildrenById(id uint32) (*FolderForFileChildren, error) {
	s := &FolderForFileChildren{}
	d := m.DB.Self.Table("foldersforfiles").Where("id = ?", id).Find(&s)
	return s, d.Error
}

// CreateFileFolder ... 事务
func CreateFileFolder(db *gorm.DB, folder *FolderForFileModel, fatherId uint32, fatherType bool) (uint32, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := folder.Create(); err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	if fatherType { // 1->project 0->file folder
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			tx.Rollback()
			return uint32(0), err
		}
		item.DocChildren = fmt.Sprintf("%s,%d-%d", item.DocChildren, folder.ID, 1)
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
		item.Children = fmt.Sprintf("%s,%d-%d", item.Children, folder.ID, 1)
		if err := item.Update(); err != nil {
			tx.Rollback()
			return uint32(0), err
		}
	}

	return folder.ID, tx.Commit().Error
}

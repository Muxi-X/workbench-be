package model

import (
	"github.com/jinzhu/gorm"
	m "muxi-workbench/model"
)

// FolderForFileModel ... 文件文件夹模型
type FolderForFileModel struct {
	FolderModel
}

// TableName ... 物理表名
func (u *FolderForFileModel) TableName() string {
	return "foldersforfiles"
}

// Create ... 创建文件文件夹
func (u *FolderForFileModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// Update ... 更新文件文件夹
func (u *FolderForFileModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetFolderForFileDetail ... 获取文件文件夹详情
func GetFolderForFileDetail(id uint32) (*FolderDetail, error) {
	s := &FolderDetail{}
	d := m.DB.Self.Table("foldersforfiles").Select("foldersforfiles.*, u.name as creator").Joins("left join users u on u.id = foldersforfiles.create_id").Where("foldersforfiles.id = ? AND re = 0", id).First(&s)
	return s, d.Error
}

// GetFolderForFileModel ... 获取文件文件夹
func GetFolderForFileModel(id uint32) (*FolderModel, error) {
	s := &FolderModel{}
	d := m.DB.Self.Table("foldersforfiles").Where("id = ? AND re = 0", id).First(&s)
	return s, d.Error
}

// GetFolderForFileInfoByIds ... 获取文件文件夹信息列表
func GetFolderForFileInfoByIds(ids []uint32) ([]*FolderInfo, error) {
	s := make([]*FolderInfo, 0)
	d := m.DB.Self.Table("foldersforfiles").Where("id IN (?) AND re = 0", ids).Find(&s)
	return s, d.Error
}

// GetFileChildrenById ... 获取子文件
func GetFileChildrenById(id uint32) (*FolderChildren, error) {
	s := &FolderChildren{}
	d := m.DB.Self.Table("foldersforfiles").Where("id = ? AND re = 0", id).Find(&s)
	return s, d.Error
}

// CreateFileFolder ... 事务
func CreateFileFolder(db *gorm.DB, folder *FolderForFileModel, childrenPositionIndex uint32) (uint32, error) {
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

	// 获取 fatherId
	fatherId := folder.FatherId
	isFatherProject := false
	if folder.FatherId == 0 {
		isFatherProject = true
		fatherId = folder.ProjectID
	}

	if err := AddChildren(tx, isFatherProject, fatherId, childrenPositionIndex, folder); err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	return folder.ID, tx.Commit().Error
}

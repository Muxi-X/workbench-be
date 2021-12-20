package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// FolderForFileDetail ... 文件文件夹详情
type FolderForFileDetail struct {
	Creator string `json:"creator" gorm:"column:creator;not null" binding:"required"`
	FolderForFileModel
}

// FolderForFileInfo ... 文件文件夹信息
type FolderForFileInfo struct {
	ID        uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name      string `json:"name" gorm:"column:filename;" binding:"required"`
	ProjectID uint32 `json:"project_id" gorm:"column:project_id;" binding:"required"`
}

type FolderForFileChildren struct {
	Children  string `json:"children" gorm:"column:children;" binding:"required"`
	ProjectID uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

// FolderForFileModel ... 文件文件夹模型
type FolderForFileModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:name;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	CreatorID  uint32 `json:"creatorID" gorm:"column:create_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	Children   string `json:"children" gorm:"column:children;" binding:"required"`
	FatherId   uint32 `json:"father_id" gorm:"column:father_id;" binding:"required"`
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
func GetFolderForFileDetail(id uint32) (*FolderForFileDetail, error) {
	s := &FolderForFileDetail{}
	d := m.DB.Self.Select("foldersforfiles.*, u.name as creator").Joins("left join users u on u.id = foldersforfiles.create_id").Where("foldersforfiles.id = ? AND re = 0", id).First(&s)
	return s, d.Error
}

// GetFolderForFileModel ... 获取文件文件夹
func GetFolderForFileModel(id uint32) (*FolderForFileModel, error) {
	s := &FolderForFileModel{}
	d := m.DB.Self.Where("id = ? AND re = 0", id).First(&s)
	return s, d.Error
}

// GetFolderForFileInfoByIds ... 获取文件文件夹信息列表
func GetFolderForFileInfoByIds(ids []uint32) ([]*FolderForFileInfo, error) {
	s := make([]*FolderForFileInfo, 0)
	d := m.DB.Self.Table("foldersforfiles").Where("id IN (?) AND re = 0", ids).Find(&s)
	return s, d.Error
}

// GetFileChildrenById ... 获取子文件
func GetFileChildrenById(id uint32) (*FolderForFileChildren, error) {
	s := &FolderForFileChildren{}
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

	if err := AddFileChildren(tx, isFatherProject, fatherId, childrenPositionIndex, folder); err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	return folder.ID, tx.Commit().Error
}

func DeleteFileFolder(db *gorm.DB, trashbin *TrashbinModel, fatherId uint32, isFatherProject bool) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 获取时间
	day := viper.GetInt("trashbin.expired")
	t := time.Now().Unix()
	trashbin.ExpiresAt = t + int64(time.Hour*24*time.Duration(day))

	if err := tx.Create(trashbin).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 获取子文件，同步 redis
	var res []string
	if err := GetFileChildFolder(trashbin.FileId, &res); err != nil {
		tx.Rollback()
		return err
	}
	if len(res) != 0 {
		if err := m.SAddToRedis(constvar.Trashbin, res); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := DeleteFileChildren(tx, isFatherProject, fatherId, trashbin.FileId, constvar.IsFolderCode); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

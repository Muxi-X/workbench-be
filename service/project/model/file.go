package model

import (
	"fmt"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// FileDetail ... 文件详情
type FileDetail struct {
	Creator string `json:"creator" gorm:"column:creator;" binding:"required"`
	FileModel
}

// FileInfo ... 文件信息
type FileInfo struct {
	ID        uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name      string `json:"name" gorm:"column:realname;" binding:"required"`
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
	d := m.DB.Self.Table("files").Where("id IN (?) AND re = 0", ids).Find(&s)
	return s, d.Error
}

// GetFileDetail ... 获取文件详情
func GetFileDetail(id uint32) (*FileDetail, error) {
	s := &FileDetail{}
	d := m.DB.Self.Table("files").Where("files.id = ? AND re = 0", id).Select("files.*, users.name as creator").Joins("left join users on users.id = files.creator_id").First(&s)
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

	if err := AddFileChildren(tx, isFatherProject, fatherId, childrenPositionIndex, file); err != nil {
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

func DeleteFile(db *gorm.DB, trashbin *TrashbinModel, fatherId uint32, isFatherProject bool) error {
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

	// 同步 redis
	// 不需要找子文件夹
	if err := m.SAddToRedis(constvar.Trashbin,
		fmt.Sprintf("%d-%d", trashbin.FileId, constvar.FileCode)); err != nil {
		tx.Rollback()
		return err
	}

	if err := DeleteFileChildren(tx, isFatherProject, fatherId, trashbin.FileId, constvar.NotFolderCode); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

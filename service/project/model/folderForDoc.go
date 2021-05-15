package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// FolderForDocInfo ... 文档文件夹信息
type FolderForDocInfo struct {
	ID   uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name string `json:"name" gorm:"column:name;" binding:"required"`
}

// FolderForDocChildren ... 子文档/夹
type FolderForDocChildren struct {
	Children string `json:"children" gorm:"column:children;" binding:"required"`
}

// FolderForDocModel ... 文档文件夹模型
type FolderForDocModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:name;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	CreatorID  uint32 `json:"creatorID" gorm:"column:create_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	Children   string `json:"children" gorm:"column:children;" binding:"required"`
	FatherId   uint32 `json:"fahter_id" gorm:"column:father_id;" binding:"required"`
}

// TableName ... 物理表名
func (u *FolderForDocModel) TableName() string {
	return "foldersformds"
}

// Create ... 创建文档文件夹
func (u *FolderForDocModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// Update ... 更新文档文件夹
func (u *FolderForDocModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetFolderForDocModel ... 获取文档文件夹
func GetFolderForDocModel(id uint32) (*FolderForDocModel, error) {
	s := &FolderForDocModel{}
	d := m.DB.Self.Where("id = ? AND re = 0", id).First(&s)
	return s, d.Error
}

// GetFolderForDocInfoByIds ... 获取文档文件夹信息列表
func GetFolderForDocInfoByIds(ids []uint32) ([]*FolderForDocInfo, error) {
	s := make([]*FolderForDocInfo, 0)
	d := m.DB.Self.Table("foldersformds").Where("id IN (?) AND re = 0", ids).Find(&s)
	return s, d.Error
}

// GetDocChildrenById ... 获取子文档/夹
func GetDocChildrenById(id uint32) (*FolderForDocChildren, error) {
	s := &FolderForDocChildren{}
	d := m.DB.Self.Table("foldersformds").Select("children").Where("id = ? AND re = 0", id).Find(&s)
	return s, d.Error
}

// CreateDocFolder ... 事务
func CreateDocFolder(db *gorm.DB, folder *FolderForDocModel, childrenPositionIndex uint32) (uint32, error) {
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

	if err := AddDocChildren(tx, isFatherProject, fatherId, childrenPositionIndex, folder); err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	return folder.ID, tx.Commit().Error
}

func DeleteDocFolder(db *gorm.DB, trashbin *TrashbinModel, fatherId uint32, isFatherProject bool) error {
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

	// 插入回收站
	if err := trashbin.Create(); err != nil {
		tx.Rollback()
		return err
	}

	// 获取子文件，同步 redis
	var res []string
	if err := GetDocChildFolder(trashbin.FileId, &res); err != nil {
		tx.Rollback()
		return err
	}
	if len(res) != 0 {
		if err := m.SAddToRedis(constvar.Trashbin, res); err != nil {
			tx.Rollback()
			return err
		}
	}

	// 修改文件树
	if err := DeleteDocChildren(tx, isFatherProject, fatherId, trashbin.FileId, constvar.IsFolderCode); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

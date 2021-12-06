package model

import (
	"fmt"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DocDetail ... 文档详情
type DocDetail struct {
	Creator string `json:"creator" gorm:"column:creator;not null" binding:"required"`
	Editor  string `json:"editor" gorm:"column:editor;" binding:"required"`
	DocModel
}

// DocInfo ... 文档信息
type DocInfo struct {
	ID   uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name string `json:"name" gorm:"column:filename;" binding:"required"`
}

// DocModel ... 文档物理模型
type DocModel struct {
	ID           uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string `json:"name" gorm:"column:filename;" binding:"required"`
	Content      string `json:"content" gorm:"column:content;" binding:"required"`
	Re           bool   `json:"re" gorm:"column:re;" binding:"required"`
	Top          bool   `json:"top" gorm:"column:top;" binding:"required"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	DeleteTime   string `json:"deleteTime" gorm:"column:delete_time;" binding:"required"`
	TeamID       uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	CreatorID    uint32 `json:"creatorId" gorm:"column:creator_id;" binding:"required"`
	EditorID     uint32 `json:"editorId" gorm:"column:editor_id;" binding:"required"`
	ProjectID    uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	LastEditTime string `json:"lastEditTime" gorm:"column:last_edit_time;" binding:"required"`
	FatherId     uint32 `json:"father_id" gorm:"column:father_id;" binding:"required"`
}

// TableName ... 物理表名
func (u *DocModel) TableName() string {
	return "docs"
}

// Create doc
func (u *DocModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// Update doc
func (u *DocModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetDoc ... 获取文档
func GetDoc(id uint32) (*DocModel, error) {
	s := &DocModel{}
	d := m.DB.Self.Where("id = ? AND re = 0", id).First(&s)
	return s, d.Error
}

// GetDocInfo ... 获取文档信息
func GetDocInfo(id uint32) (*DocInfo, error) {
	info := &DocInfo{}
	d := m.DB.Self.Table("docs").Select("id,name").Where("id = ? AND re = 0", id).Scan(&info)
	return info, d.Error
}

// GetDocInfoByIds ... 获取文档信息列表
func GetDocInfoByIds(ids []uint32) ([]*DocInfo, error) {
	s := make([]*DocInfo, 0)
	d := m.DB.Self.Table("docs").Where("id IN (?) AND re = 0", ids).Find(&s)
	return s, d.Error
}

// GetDocDetail ... 获取文档详情
func GetDocDetail(id uint32) (*DocDetail, error) {
	s := &DocDetail{}
	// multiple left join
	d := m.DB.Self.Table("docs").Where("docs.id = ? AND re = 0", id).Select("docs.*, c.name as creator, e.name as editor").Joins("left join users c on c.id = docs.creator_id").Joins("left join users e on e.id = docs.editor_id").First(&s)
	return s, d.Error
}

func CreateDoc(db *gorm.DB, doc *DocModel, childrenPositionIndex uint32) (uint32, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(doc).Error; err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	// 获取 isFatherProject
	isFatherProject := false
	fatherId := doc.FatherId
	if doc.FatherId == 0 {
		isFatherProject = true
		fatherId = doc.ProjectID
	}

	if err := AddDocChildren(tx, isFatherProject, fatherId, childrenPositionIndex, doc); err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	return doc.ID, tx.Commit().Error
}

// DeleteDoc ... 插入回收站 同步 redis
// 先查表找到 childrenPositionIndex
func DeleteDoc(db *gorm.DB, trashbin *TrashbinModel, fatherId uint32, isFatherProject bool) error {
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
	if err := tx.Create(trashbin).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 同步 redis
	// 不需要找子文件夹
	if err := m.SAddToRedis(constvar.Trashbin,
		fmt.Sprintf("%d-%d", trashbin.FileId, constvar.DocCode)); err != nil {
		tx.Rollback()
		return err
	}

	if err := DeleteDocChildren(tx, isFatherProject, fatherId, trashbin.FileId, constvar.NotFolderCode); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

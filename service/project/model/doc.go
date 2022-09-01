package model

import (
	"github.com/jinzhu/gorm"
	m "muxi-workbench/model"
)

// DocModel ... 文档物理模型
type DocModel struct {
	ID           uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string `json:"name" gorm:"column:filename;" binding:"required"`
	Content      string `json:"content" gorm:"column:content;" binding:"required"`
	Re           bool   `json:"re" gorm:"column:re;" binding:"required"`
	Top          bool   `json:"top" gorm:"column:top;" binding:"required"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	DeleteTime   string `json:"deleteTime" gorm:"column:delete_time;" binding:"required"`
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
func GetDocInfo(id uint32) (*FileInfo, error) {
	info := &FileInfo{}
	d := m.DB.Self.Table("docs").Select("id,name").Where("id = ? AND re = 0", id).Scan(&info)
	return info, d.Error
}

// GetDocInfoByIds ... 获取文档信息列表
func GetDocInfoByIds(ids []uint32) ([]*FileInfo, error) {
	s := make([]*FileInfo, 0)
	d := m.DB.Self.Table("docs").Where("id IN (?) AND re = 0", ids).Find(&s)
	return s, d.Error
}

// GetDocDetail ... 获取文档详情
func GetDocDetail(id uint32) (*FileDetail, error) {
	s := &FileDetail{}
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

	if err := AddChildren(tx, isFatherProject, fatherId, childrenPositionIndex, doc); err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	return doc.ID, tx.Commit().Error
}

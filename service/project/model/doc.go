package model

import (
	"errors"
	"fmt"
	m "muxi-workbench/model"

	"github.com/jinzhu/gorm"
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
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

// GetDocInfo ... 获取文档信息
func GetDocInfo(id uint32) (*DocInfo, error) {
	s := &DocModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	info := &DocInfo{}
	info.ID = s.ID
	info.Name = s.Name
	return info, d.Error
}

// GetDocInfoByIds ... 获取文档信息列表
func GetDocInfoByIds(ids []uint32) ([]*DocInfo, error) {
	s := make([]*DocInfo, 0)
	d := m.DB.Self.Table("docs").Where("id IN (?)", ids).Find(&s)
	return s, d.Error
}

// GetDocDetail ... 获取文档详情
func GetDocDetail(id uint32) (*DocDetail, error) {
	s := &DocDetail{}
	// multiple left join
	d := m.DB.Self.Table("docs").Where("docs.id = ?", id).Select("docs.*, c.name as creator, e.name as editor").Joins("left join users c on c.id = docs.creator_id").Joins("left join users e on e.id = docs.editor_id").First(&s)
	return s, d.Error
}

func CreateDoc(db *gorm.DB, doc *DocModel, fatherId, childrenPositionIndex uint32, fatherType bool) (uint32, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := doc.Create(); err != nil {
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

		// 根据 childrenPositionIndex 判断插入位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.DocChildren) {
			item.DocChildren = fmt.Sprintf("%s%d-%d,%s", item.DocChildren[:index], doc.ID, 0, item.DocChildren[index:])
		} else if index-1 == len(item.DocChildren) {
			item.DocChildren = fmt.Sprintf("%s,%d-%d", item.DocChildren, doc.ID, 0)
		} else {
			tx.Rollback()
			return uint32(0), errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			tx.Rollback()
			return uint32(0), err
		}
	} else {
		item, err := GetFolderForDocModel(fatherId)
		if err != nil {
			tx.Rollback()
			return uint32(0), err
		}

		// 根据 childrenPositionIndex 判断插入位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.Children) {
			item.Children = fmt.Sprintf("%s%d-%d,%s", item.Children[:index], doc.ID, 0, item.Children[index:])
		} else if index-1 == len(item.Children) {
			item.Children = fmt.Sprintf("%s,%d-%d", item.Children, doc.ID, 0)
		} else {
			tx.Rollback()
			return uint32(0), errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			tx.Rollback()
			return uint32(0), err
		}
	}

	return doc.ID, tx.Commit().Error
}

func DeleteDoc(db *gorm.DB, doc *DocModel, fatherId, childrenPositionIndex uint32, fatherType bool) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := doc.Update(); err != nil {
		tx.Rollback()
		return err
	}

	if fatherType {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 根据 childrenPositionIndex 判断删除位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.DocChildren) {
			item.DocChildren = item.DocChildren[:index] + item.DocChildren[index+1:]
		} else if index-1 == len(item.DocChildren) {
			item.DocChildren = item.DocChildren[:index]
		} else {
			tx.Rollback()
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		item, err := GetFolderForDocModel(fatherId)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 根据 childrenPositionIndex 判断删除位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.Children) {
			item.Children = item.Children[:index] + item.Children[index+1:]
		} else if index-1 == len(item.Children) {
			item.Children = item.Children[:index]
			tx.Rollback()
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

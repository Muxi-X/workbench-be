package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"

	"github.com/jinzhu/gorm"
)

type CommentModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Kind     uint32 `json:"kind" gorm:"column:kind;" binding:"required"` // 0 是一级，1 是二级
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Creator  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	TargetID uint32 `json:"target_id" gorm:"column:target_id;" binding:"required"`
	Re       bool   `json:"re" gorm:"column:re;" binding:"required"`
}

type CommentListItem struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Kind     uint32 `json:"kind" gorm:"column:kind;" binding:"required"` // 0 是一级，1 是二级
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Creator  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	Avatar   string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	UserName string `json:"username" gorm:"column:name;" binding:"required"`
}

func (c *CommentModel) TableName() string {
	return "comments_status"
}

func (c *CommentModel) Create(tx *gorm.DB) error {
	return tx.Create(c).Error
}

func (c *CommentModel) Update(content string) error {
	return m.DB.Self.Model(c).Update("content", content).Error
}

// Delete ... 软删除评论
func (c *CommentModel) Delete(tx *gorm.DB, uid uint32) error {
	return tx.Table("comments_status").Where("id = ? AND creator = ?", c.ID, uid).Update("re", 1).Error
}

func (c *CommentModel) GetModelById(id uint32) error {
	return m.DB.Self.Where("id = ? AND re = 0", id).First(c).Error
}

func ListComments(statusID, offset, limit, lastID uint32) ([]*CommentListItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	commentsList := make([]*CommentListItem, 0)

	query := m.DB.Self.Table("comments_status").Select("comments_status.*, users.name, users.avatar").Where("target_id = ? AND kind = 0 AND re = 0", statusID).Joins("left join users on users.id = comments_status.creator").Offset(offset).Limit(limit).Order("comments_status.id asc")

	if lastID != 0 {
		query = query.Where("comments_status.id < ?", lastID)
	}

	var count uint64
	if err := query.Scan(&commentsList).Count(&count).Error; err != nil {
		return commentsList, count, err
	}

	for i := 0; uint64(i) < count; i++ {
		comment := commentsList[i]
		commentsLevel2 := make([]*CommentListItem, 0)
		m.DB.Self.Table("comments_status").Select("comments_status.*, users.name, users.avatar").Where("target_id = ? AND kind = 1 AND re = 0", comment.ID).Joins("left join users on users.id = comments_status.creator").Order("comments_status.id asc").Scan(&commentsLevel2)
		commentsList = append(commentsList, commentsLevel2...)
	}

	return commentsList, count, nil
}

// CreateStatusComment ... 创建 comment 和修改 status 的评论总数
func CreateStatusComment(db *gorm.DB, u *CommentModel, m *StatusModel) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := u.Create(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(m).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DeleteStatusComment ... 删除 comment 和修改 status 的评论总数
func DeleteStatusComment(db *gorm.DB, id uint32, uid uint32) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var comment CommentModel
	comment.GetModelById(id)

	if err := comment.Delete(tx, uid); err != nil {
		tx.Rollback()
		return err
	}

	status, err := GetStatus(comment.TargetID)
	if err != nil {
		tx.Rollback()
		return err
	}
	status.Comment -= 1
	if err := tx.Save(status).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

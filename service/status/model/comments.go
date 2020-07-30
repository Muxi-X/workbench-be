package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type CommentsModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Kind     uint32 `json:"kind" gorm:"column:kind;" binding:"required"` // 0 是进度，1 是文档或者文件
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Creator  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	DocID    uint32 `json:"docId" gorm:"column:doc_id;"`
	FileID   uint32 `json:"fileId" gorm:"column:file_id;"`
	StatusID uint32 `json:"statusId" gorm:"column:statu_id;"`
}

type CommentListItem struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Creator  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	Avatar   string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	UserName string `json:"username" gorm:"column:name;" binding:"required"`
}

func (c *CommentsModel) TableName() string {
	return "comments"
}

// Create comments
func (u *CommentsModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// ListComments list all comments
func ListComments(statusID, offset, limit, lastID uint32) ([]*CommentListItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	commentsList := make([]*CommentListItem, 0)

	query := m.DB.Self.Table("comments").Select("comments.*, users.name, users.avatar").Where("comments.statu_id = ?", statusID).Joins("left join users on users.id = comments.creator").Offset(offset).Limit(limit).Order("comments.id desc")

	if lastID != 0 {
		query = query.Where("comments.id < ?", lastID)
	}

	var count uint64

	if err := query.Scan(&commentsList).Count(&count).Error; err != nil {
		return commentsList, count, err
	}

	return commentsList, count, nil
}

func DeleteComment(id uint32) error {
	comment := &CommentsModel{}
	comment.ID = id
	return m.DB.Self.Delete(&comment).Error
}

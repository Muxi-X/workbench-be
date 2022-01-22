package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type DocCommentModel struct {
	CommentModel
}

func (d *DocCommentModel) TableName() string {
	return "comments_docs"
}

func (*DocCommentModel) Create(comment CommentModel) error {
	d := DocCommentModel{CommentModel: comment}
	return m.DB.Self.Create(&d).Error
}

func (d *DocCommentModel) Update(content string) error {
	return m.DB.Self.Model(d).Update("content", content).Error
}

// Delete ... 软删除评论
func (d *DocCommentModel) Delete(uid uint32) error {
	return m.DB.Self.Table("comments_docs").Where("id = ? AND creator = ?", d.ID, uid).Update("re", 1).Error
}

func (d *DocCommentModel) GetModelById(id uint32) error {
	return m.DB.Self.Where("id = ? AND re = 0", id).First(d).Error
}

func (d *DocCommentModel) Verify(uId uint32) bool {
	return d.Creator == uId
}

func (d *DocCommentModel) List(docID, offset, limit, lastID uint32) ([]*CommentListItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	commentsList := make([]*CommentListItem, 0)

	query := m.DB.Self.Table("comments_docs").Select("comments_docs.*, users.name, users.avatar").Where("target_id = ? AND kind = 0 AND re = 0", docID).Joins("LEFT JOIN users ON users.id = comments_docs.creator").Offset(offset).Limit(limit).Order("comments_docs.id ASC")

	if lastID != 0 {
		query = query.Where("comments_docs.id < ?", lastID)
	}

	var count uint64
	if err := query.Scan(&commentsList).Count(&count).Error; err != nil {
		return commentsList, count, err
	}

	for i := 0; uint64(i) < count; i++ {
		comment := commentsList[i]
		commentsLevel2 := make([]*CommentListItem, 0)
		m.DB.Self.Table("comments_docs").Select("comments_docs.*, users.name, users.avatar").Where("target_id = ? AND kind = 1 AND re = 0", comment.ID).Joins("LEFT JOIN users ON users.id = comments_docs.creator").Order("comments_docs.id ASC").Scan(&commentsLevel2)
		commentsList = append(commentsList, commentsLevel2...)
	}

	return commentsList, count, nil
}

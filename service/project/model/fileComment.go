package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type FileCommentModel struct {
	CommentModel
}

func (f *FileCommentModel) TableName() string {
	return "comments_files"
}

func (*FileCommentModel) Create(comment CommentModel) error {
	f := FileCommentModel{CommentModel: comment}
	return m.DB.Self.Create(&f).Error
}

func (f *FileCommentModel) Update(content string) error {
	return m.DB.Self.Model(f).Update("content", content).Error
}

// Delete ... 软删除评论
func (f *FileCommentModel) Delete(uid uint32) error {
	return m.DB.Self.Table("comments_files").Where("id = ? AND creator = ?", f.ID, uid).Update("re", 1).Error
}

func (f *FileCommentModel) GetModelById(id uint32) error {
	return m.DB.Self.Where("id = ? AND re = 0", id).First(f).Error
}

func (f *FileCommentModel) Verify(uId uint32) bool {
	return f.Creator == uId
}

func (f *FileCommentModel) List(fileID, offset, limit, lastID uint32) ([]*CommentListItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	commentsList := make([]*CommentListItem, 0)

	query := m.DB.Self.Table("comments_files").Select("comments_files.*, users.name, users.avatar").Where("target_id = ? AND kind = 0 AND re = 0", fileID).Joins("left join users on users.id = comments_files.creator").Offset(offset).Limit(limit).Order("comments_files.id asc")

	if lastID != 0 {
		query = query.Where("comments_files.id < ?", lastID)
	}

	var count uint64
	if err := query.Scan(&commentsList).Count(&count).Error; err != nil {
		return commentsList, count, err
	}

	for i := 0; uint64(i) < count; i++ {
		comment := commentsList[i]
		commentsLevel2 := make([]*CommentListItem, 0)
		m.DB.Self.Table("comments_files").Select("comments_files.*, users.name, users.avatar").Where("target_id = ? AND kind = 1 AND re = 0", comment.ID).Joins("left join users on users.id = comments_files.creator").Order("comments_files.id asc").Scan(&commentsLevel2)
		commentsList = append(commentsList, commentsLevel2...)
	}

	return commentsList, count, nil
}

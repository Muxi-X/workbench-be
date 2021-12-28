package model

import (
	// "muxi-workbench-project/model"
	m "muxi-workbench/model"
)

type StatusCommentModel struct {
	// model.CommentModel
}

func (s *StatusCommentModel) TableName() string {
	return "comments_status"
}

// func (*StatusCommentModel) Create(comment model.CommentModel) error {
// 	s := StatusCommentModel{CommentModel: comment}
// 	return m.DB.Self.Create(&s).Error
// }

func (s *StatusCommentModel) Update(content string) error {
	return m.DB.Self.Update("content", content).Error
}

// Delete ... 软删除评论
func (s *StatusCommentModel) Delete(id, uid uint32) error {
	return m.DB.Self.Model(s).Where("id = ? AND creator = ?", id, uid).Update("re", "1").Error
}

func (s *StatusCommentModel) GetModelById(id uint32) error {
	return m.DB.Self.Where("id = ?", id).First(s).Error
}

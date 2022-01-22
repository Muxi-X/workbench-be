package handler

import (
	"PROJECT_SCRIPT/model"
)

func CommentMigrate() {
	comments := model.GetAllComments()

	for _, comment := range comments {
		if comment.StatusID != 0 {
			err := model.CreateStatusComment(model.NewCommentModel{
				Kind:     0,
				Content:  comment.Content,
				Time:     comment.Time,
				Creator:  comment.Creator,
				TargetID: comment.StatusID,
			})
			if err != nil {
				panic(err)
			}
		} else if comment.DocID != 0 {
			err := model.CreateDocComment(model.NewCommentModel{
				Kind:     0,
				Content:  comment.Content,
				Time:     comment.Time,
				Creator:  comment.Creator,
				TargetID: comment.DocID,
			})
			if err != nil {
				panic(err)
			}
		} else if comment.FileID != 0 {
			err := model.CreateFileComment(model.NewCommentModel{
				Kind:     0,
				Content:  comment.Content,
				Time:     comment.Time,
				Creator:  comment.Creator,
				TargetID: comment.FileID,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

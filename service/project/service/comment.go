package service

import (
	"fmt"
	"muxi-workbench-project/model"
	"muxi-workbench/pkg/constvar"
)

// getCommentType 得到需要comment的类型
func getCommentType(typeId uint8) (model.Commenter, error) {
	switch typeId {
	case constvar.DocCommentLevel1, constvar.DocCommentLevel2:
		return &model.DocCommentModel{}, nil
	case constvar.FileCommentCodeLevel1, constvar.FileCommentCodeLevel2:
		return &model.FileCommentModel{}, nil
	}
	err := fmt.Errorf("wrong type_id(%d)", typeId)
	return nil, err
}

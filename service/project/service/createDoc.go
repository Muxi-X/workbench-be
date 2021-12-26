package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateDoc ... 创建文档
// 事务自动更新文件树
func (s *Service) CreateDoc(ctx context.Context, req *pb.CreateDocRequest, res *pb.ProjectIDResponse) error {
	t := time.Now()

	doc := model.DocModel{
		CreatorID:    req.UserId,
		Name:         req.Title,
		Content:      req.Content,
		Re:           false,
		Top:          false,
		CreateTime:   t.Format("2006-01-02 15:04:05"),
		ProjectID:    req.ProjectId,
		EditorID:     req.UserId,
		LastEditTime: t.Format("2006-01-02 15:04:05"),
		FatherId:     req.FatherId,
	}

	// 事务
	id, err := model.CreateDoc(&doc, req.ChildrenPositionIndex)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = id

	return nil
}

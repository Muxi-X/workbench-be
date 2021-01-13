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
func (s *Service) CreateDoc(ctx context.Context, req *pb.CreateDocRequest, res *pb.ProjectNameAndIDResponse) error {
	t := time.Now()

	doc := model.DocModel{
		Name:       req.Title,
		Content:    req.Content,
		Re:         false,
		Top:        false,
		TeamID:     0, // 查询一下用户信息
		CreateTime: t.Format("2006-01-02 15:04:05"),
		ProjectID:  req.ProjectId,
		EditorID:   req.UserId,
	}

	if err := doc.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	name, err := model.GetProjectName(req.ProjectId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = doc.ID
	res.Name = name

	return nil
}

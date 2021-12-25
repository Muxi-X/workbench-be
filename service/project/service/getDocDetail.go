package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetDocDetail ... 获取文档详情
func (s *Service) GetDocDetail(ctx context.Context, req *pb.GetFileDetailRequest, res *pb.DocDetail) error {
	// 判断自己 id 和父 id 是否被删
	// 不用判断 father 是否是 project，因为 project 被删访问不到。
	isDeleted, err := model.AdjustSelfAndFatherIfExist(req.Id, req.FatherId, constvar.DocCode, constvar.DocFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted {
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}
	doc, err := model.GetDocDetail(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	if doc.ProjectID != req.ProjectId {
		return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
	}

	project, err := model.GetProject(doc.ProjectID)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	res.Title = doc.Name
	res.Id = doc.ID
	res.Content = doc.Content
	res.Creator = doc.Creator
	res.LastEditor = doc.Editor
	res.CreateTime = doc.CreateTime
	res.LastEditTime = doc.LastEditTime
	res.ProjectName = project.Name
	return nil
}

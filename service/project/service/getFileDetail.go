package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetFileDetail ... 获取文件详情
func (s *Service) GetFileDetail(ctx context.Context, req *pb.GetFileDetailRequest, res *pb.FileDetail) error {
	// 判断自己 id 和父 id 是否被删
	isDeleted, err := model.AdjustSelfAndFatherIfExist(req.Id, req.FatherId, uint8(req.TypeId), uint8(req.TypeId)+2)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted { // 存在 redis 返回 1, 说明被删了
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	var file *model.FileDetail
	if uint8(req.TypeId) == constvar.DocCode {
		file, err = model.GetDocDetail(req.Id)
	} else if uint8(req.TypeId) == constvar.FileCode {
		file, err = model.GetFileDetail(req.Id)
	} else {
		return e.BadRequestErr(errno.ErrBind, "wrong type_id")
	}

	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if file.ProjectID != req.ProjectId {
		return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
	}

	projectName, err := model.GetProjectName(file.ProjectID)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = file.ID
	res.Url = file.URL
	res.Creator = file.Creator
	res.CreateTime = file.CreateTime
	res.Name = file.Name
	res.Content = file.Content
	res.LastEditTime = file.LastEditTime
	res.Editor = file.Editor
	res.ProjectName = projectName

	return nil
}

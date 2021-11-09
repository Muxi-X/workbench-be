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
	isDeleted, err := model.AdjustSelfAndFatherIfExist(req.Id, req.FatherId, constvar.FileCode, constvar.FileFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted { // 存在 redis 返回 1, 说明被删了
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	// ok
	file, err := model.GetFileDetail(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 此处为了获取project 改了projectDetail 可能出问题
	project, err := model.GetProject(file.ProjectID)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	res.Id = file.ID
	res.Url = file.URL
	res.Creator = file.Creator
	res.CreateTime = file.CreateTime
	res.Name = file.Name
	res.ProjectName = project.Name
	return nil
}

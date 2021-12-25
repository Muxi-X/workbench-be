package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetFolderInfoList ... 获取文档文件夹信息列表
func (s *Service) GetFolderInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetFolderListResponse) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	scope, err := model.AdjustFolderListIfExist(req.List, uint8(req.TypeId))
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if scope == nil {
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	// 获取文档夹的名字信息
	var f func([]uint32) ([]*model.FolderInfo, error)
	if uint8(req.TypeId) == constvar.DocFolderCode {
		f = model.GetFolderForDocInfoByIds
	} else if uint8(req.TypeId) == constvar.FileFolderCode {
		f = model.GetFolderForFileInfoByIds
	}

	list, err := f(scope)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.FolderInfo, 0)

	for _, item := range list {
		if item.ProjectID != req.ProjectId {
			return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
		}
		resList = append(resList, &pb.FolderInfo{
			Id:   item.ID,
			Name: item.Name,
		})
	}

	res.List = resList

	return nil
}

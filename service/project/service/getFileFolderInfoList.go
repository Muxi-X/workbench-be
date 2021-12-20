package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetFileFolderInfoList ... 获取文件文件夹信息列表
func (s *Service) GetFileFolderInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetFileFolderListResponse) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	scope, err := model.AdjustFolderListIfExist(req.List, constvar.FileFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if scope == nil {
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	// ok
	list, err := model.GetFolderForFileInfoByIds(scope)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.FileFolderDetail, 0)

	for _, item := range list {
		if item.ProjectID != req.ProjectId {
			return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
		}
		resList = append(resList, &pb.FileFolderDetail{
			Id:   item.ID,
			Name: item.Name,
		})
	}

	res.List = resList

	return nil
}

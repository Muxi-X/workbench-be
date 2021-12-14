package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetFileInfoList ... 获取文件信息列表
func (s *Service) GetFileInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetFileInfoListResponse) error {
	// 新增过滤 id
	scope, err := model.AdjustFileListIfExist(req.List, req.FatherId, constvar.FileCode, constvar.FileFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if scope == nil {
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	// ok
	list, err := model.GetFileInfoByIds(scope)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.FileInfo, 0)

	for _, item := range list {
		resList = append(resList, &pb.FileInfo{
			Id:    item.ID,
			Title: item.Name,
		})
	}

	res.List = resList

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetDocFolderInfoList ... 获取文档文件夹信息列表
func (s *Service) GetDocFolderInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetDocFolderListResponse) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	scope, err := model.AdjustFolderListIfExist(req.List, constvar.DocFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if scope == nil {
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	// 获取文档夹的名字信息
	list, err := model.GetFolderForDocInfoByIds(scope)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.DocFolderDetail, 0)

	for index := 0; index < len(scope); index++ {
		item := list[index]
		resList = append(resList, &pb.DocFolderDetail{
			Id:   item.ID,
			Name: item.Name,
		})
	}

	res.List = resList

	return nil
}

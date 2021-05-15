package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetDocInfoList ... 获取文档信息列表
// TODO: 加上 father_id 限定条件，给的 idlist 必须是对应 father 的 children
func (s *Service) GetDocInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetDocInfoListResponse) error {
	// 新增过滤 id
	scope, err := model.AdjustFileListIfExist(req.List, req.FatherId, constvar.DocCode, constvar.DocFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if scope == nil {
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	// 获取文档的名字信息
	list, err := model.GetDocInfoByIds(scope)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.DocInfo, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.DocInfo{
			Id:    item.ID,
			Title: item.Name,
		})
	}

	res.List = resList

	return nil
}

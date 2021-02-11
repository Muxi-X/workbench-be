package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetDocInfoList ... 获取文档信息列表
func (s *Service) GetDocInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetDocInfoListResponse) error {

	// 获取文档的名字信息
	list, err := model.GetDocInfoByIds(req.List)
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

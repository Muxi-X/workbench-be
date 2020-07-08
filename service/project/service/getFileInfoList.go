package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetFileInfoList ... 获取文件信息列表
func (s *Service) GetFileInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetFileInfoListResponse) error {

	list, err := model.GetFileInfoByIds(req.List)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.FileInfo, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.FileInfo{
			Id:    item.ID,
			Title: item.Name,
		})
	}

	res.List = resList

	return nil
}

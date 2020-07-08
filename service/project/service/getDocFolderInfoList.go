package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetDocFolderInfoList ... 获取文档文件夹信息列表
func (s *Service) GetDocFolderInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetDocFolderListResponse) error {

	list, err := model.GetFolderForDocInfoByIds(req.List)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.DocFolderDetail, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.DocFolderDetail{
			Id:   item.ID,
			Name: item.Name,
		})
	}

	res.List = resList

	return nil
}

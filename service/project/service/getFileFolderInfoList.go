package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetFileFolderInfoList ... 获取文件文件夹信息列表
func (s *Service) GetFileFolderInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetFileFolderListResponse) error {

	list, err := model.GetFolderForFileInfoByIds(req.List)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.FileFolderDetail, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.FileFolderDetail{
			Id:   item.ID,
			Name: item.Name,
		})
	}

	res.List = resList

	return nil
}

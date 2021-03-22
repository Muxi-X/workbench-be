package service

import (
	"context"
	"fmt"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetFileFolderInfoList ... 获取文件文件夹信息列表
func (s *Service) GetFileFolderInfoList(ctx context.Context, req *pb.GetInfoByIdsRequest, res *pb.GetFileFolderListResponse) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	var scope []uint32
	for _, v := range req.List {
		target := fmt.Sprintf("%d-%d", v, constvar.FileFolderCode)
		isDeleted, err := m.SIsmembersFromRedis(constvar.Trashbin, target)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
		if isDeleted {
			scope = append(scope, v)
		}
	}

	// ok
	list, err := model.GetFolderForFileInfoByIds(scope)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.FileFolderDetail, 0)

	for index := 0; index < len(scope); index++ {
		item := list[index]
		resList = append(resList, &pb.FileFolderDetail{
			Id:   item.ID,
			Name: item.Name,
		})
	}

	res.List = resList

	return nil
}

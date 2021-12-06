package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
	"strconv"
	"strings"
)

// GetFileChildren ... 获取任意文件夹目录下的文件树
func (s *Service) GetFileChildren(ctx context.Context, req *pb.GetRequest, res *pb.ChildrenList) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	isDeleted, err := model.AdjustSelfIfExist(req.Id, constvar.FileFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted { // 存在 redis 返回 1, 说明被删
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	item, err := model.GetFileChildrenById(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	var list []*pb.Children
	if item.Children != "" {
		raw := strings.Split(item.Children, ",")
		for _, v := range raw {
			r := strings.Split(v, "-")
			id, _ := strconv.Atoi(r[0])
			if r[1] == "0" {
				file, err := model.GetFileDetail(uint32(id))
				if err != nil {
					return e.ServerErr(errno.ErrDatabase, err.Error())
				}
				file.Creator, err = GetInfoFromUserService(file.CreatorID)
				if err != nil {
					return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
				}
				list = append(list, &pb.Children{
					Type:        false,
					Name:        file.Name,
					CreatorName: file.Creator,
					CreatTime:   file.CreateTime,
					// TODO Path:        doc.FatherId,根据fatherId一路找上去
				})
			} else {
				folder, err := model.GetFolderForFileModel(uint32(id))
				if err != nil {
					return e.ServerErr(errno.ErrDatabase, err.Error())
				}
				creatorName, err := GetInfoFromUserService(folder.CreatorID)
				if err != nil {
					return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
				}
				list = append(list, &pb.Children{
					Type:        true,
					Name:        folder.Name,
					CreatorName: creatorName,
					CreatTime:   folder.CreateTime,
					// TODO Path:        doc.FatherId,根据fatherId一路找上去
				})
			}
		}
	}
	res.List = list

	return nil
}

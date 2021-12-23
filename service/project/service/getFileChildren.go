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
	if item.ProjectID != req.ProjectId {
		return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
	}
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	var list []*pb.Children
	if item.Children != "" {
		path := GetPath(req.Id, constvar.FileFolderCode)

		raw := strings.Split(item.Children, ",")
		for _, v := range raw {
			r := strings.Split(v, "-")
			id, _ := strconv.Atoi(r[0])
			if r[1] == "0" {
				file, err := model.GetFileDetail(uint32(id))
				if err != nil {
					return e.ServerErr(errno.ErrDatabase, err.Error())
				}
				list = append(list, &pb.Children{
					Id:          file.ID,
					Type:        false,
					Name:        file.RealName,
					CreatorName: file.Creator,
					CreatTime:   file.CreateTime,
					Path:        path,
				})
			} else {
				folder, err := model.GetFolderForFileDetail(uint32(id))
				if err != nil {
					return e.ServerErr(errno.ErrDatabase, err.Error())
				}
				list = append(list, &pb.Children{
					Id:          folder.ID,
					Type:        true,
					Name:        folder.Name,
					CreatorName: folder.Creator,
					CreatTime:   folder.CreateTime,
					Path:        path,
				})
			}
		}
	}
	res.List = list

	return nil
}

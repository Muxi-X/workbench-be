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

// GetFolderChildren ... 获取任意文件夹目录下的文件树
func (s *Service) GetFolderChildren(ctx context.Context, req *pb.GetRequest, res *pb.ChildrenList) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	isDeleted, err := model.AdjustSelfIfExist(req.Id, uint8(req.TypeId))
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted { // 存在 redis 返回 1, 说明被删
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	var getDetail func(id uint32) (*model.FileDetail, error)
	var getFolderDetail func(id uint32) (*model.FolderDetail, error)

	var item *model.FolderChildren
	if uint8(req.TypeId) == constvar.DocFolderCode {
		item, err = model.GetDocChildrenById(req.Id)
		getDetail = model.GetDocDetail
		getFolderDetail = model.GetFolderForDocDetail
	} else if uint8(req.TypeId) == constvar.FileFolderCode {
		item, err = model.GetFileChildrenById(req.Id)
		getDetail = model.GetFileDetail
		getFolderDetail = model.GetFolderForFileDetail
	}
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if item.ProjectID != req.ProjectId {
		return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
	}

	var list []*pb.Children
	if item.Children != "" {
		path := GetPath(req.Id, uint8(req.TypeId))

		raw := strings.Split(item.Children, ",")
		for _, v := range raw {
			r := strings.Split(v, "-")
			id, _ := strconv.Atoi(r[0])
			if r[1] == "0" {
				file, err := getDetail(uint32(id))
				if err != nil {
					return e.ServerErr(errno.ErrDatabase, err.Error())
				}
				list = append(list, &pb.Children{
					Id:          file.ID,
					Type:        false,
					Name:        file.Name,
					CreatorName: file.Creator,
					CreatTime:   file.CreateTime,
					Path:        path,
				})
			} else {
				folder, err := getFolderDetail(uint32(id))
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

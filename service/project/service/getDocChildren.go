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

	"github.com/jinzhu/gorm"
)

// GetDocChildren ... 获取任意文档夹目录下的文档树
func (s *Service) GetDocChildren(ctx context.Context, req *pb.GetRequest, res *pb.ChildrenList) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	isDeleted, err := model.AdjustSelfIfExist(req.Id, constvar.DocFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted {
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	item, err := model.GetDocChildrenById(req.Id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return e.NotFoundErr(errno.ErrNotFound, err.Error())
		}
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	var list []*pb.Children
	{
		raw := strings.Split(item.Children, ",")
		for _, v := range raw {
			r := strings.Split(v, "-")
			id, _ := strconv.Atoi(r[0])
			if r[1] == "0" {
				doc, err := model.GetDocDetail(uint32(id))
				if err != nil {
					return e.ServerErr(errno.ErrDatabase, err.Error())
				}
				doc.Creator, err = GetInfoFromUserService(doc.CreatorID)
				if err != nil {
					return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
				}
				list = append(list, &pb.Children{
					Type:        false,
					Name:        doc.Name,
					CreatorName: doc.Creator,
					CreatTime:   doc.CreateTime,
					// TODO Path:        doc.FatherId,根据fatherId一路找上去
				})
			} else {
				folder, err := model.GetFolderForDocModel(uint32(id))
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

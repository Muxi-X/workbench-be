package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// DeleteDocFolder ... 删除文档
// 寻找子文件同步 redis 修改文件树 插入回收站
func (s *Service) DeleteDocFolder(ctx context.Context, req *pb.DeleteRequest, res *pb.Response) error {
	item, err := model.GetFolderForDocModel(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 权限判定
	if item.CreatorID != req.UserId {
		if req.Role <= constvar.Normal {
			return e.BadRequestErr(errno.ErrPermissionDenied, "")
		}
	}

	trashbin := &model.TrashbinModel{
		FileId:   req.Id,
		FileType: constvar.DocFolderCode,
		Name:     item.Name,
	}

	// 事务
	err = model.DeleteDocFolder(m.DB.Self, trashbin, req.FatherId, req.ChildrenPositionIndex, req.FatherType)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

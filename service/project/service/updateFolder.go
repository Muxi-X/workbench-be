package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// UpdateFolder ... 修改folder，改名
func (s *Service) UpdateFolder(ctx context.Context, req *pb.UpdateFolderRequest, res *pb.Response) error {
	var item *model.FolderModel
	var err error
	if uint8(req.TypeId) == constvar.DocFolderCode {
		item, err = model.GetFolderForDocModel(req.FolderId)
	} else if uint8(req.TypeId) == constvar.FileFolderCode {
		item, err = model.GetFolderForFileModel(req.FolderId)
	} else {
		return e.BadRequestErr(errno.ErrBind, "wrong type_id")
	}
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	if item.ProjectID != req.ProjectId {
		return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
	}

	item.Name = req.Name

	if err = item.Update(uint8(req.TypeId)); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

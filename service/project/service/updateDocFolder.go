package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateDocFolder ... 修改文档夹，改名
func (s *Service) UpdateDocFolder(ctx context.Context, req *pb.UpdateFolderRequest, res *pb.Response) error {
	item, err := model.GetFolderForDocModel(req.FolderId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	if item.ProjectID != req.ProjectId {
		return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
	}

	item.Name = req.Name

	docFolder := model.FolderForDocModel{FolderModel: *item}
	if err = docFolder.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

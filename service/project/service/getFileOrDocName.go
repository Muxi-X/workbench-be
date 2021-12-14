package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetFileOrDocName ... 获取文件或文档的名字
func (s *Service) GetFileOrDocName(ctx context.Context, req *pb.GetFileOrDocNameRequest, res *pb.GetFileOrDocNameResponse) error {
	switch req.Type {
	case uint32(constvar.FileCode):
		file, err := model.GetFileDetail(req.Id)
		if err != nil {
			return e.NotFoundErr(errno.ErrDatabase, err.Error())
		}
		res.Name = file.RealName
	case uint32(constvar.DocCode):
		doc, err := model.GetDocDetail(req.Id)
		if err != nil {
			return e.NotFoundErr(errno.ErrDatabase, err.Error())
		}
		res.Name = doc.Name
	case uint32(constvar.FileFolderCode):
		fileFolder, err := model.GetFolderForFileModel(req.Id)
		if err != nil {
			return e.NotFoundErr(errno.ErrDatabase, err.Error())
		}
		res.Name = fileFolder.Name
	case uint32(constvar.DocFolderCode):
		docFolder, err := model.GetFolderForDocModel(req.Id)
		if err != nil {
			return e.NotFoundErr(errno.ErrDatabase, err.Error())
		}
		res.Name = docFolder.Name
	default:
		return e.BadRequestErr(errno.ErrGetDataFromRPC, "type == 1 -> file, 2 -> doc, 3 -> file folder, 4 -> doc folder")
	}

	return nil
}

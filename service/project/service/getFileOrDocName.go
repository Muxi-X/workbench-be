package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetFileOrDocNames ... 获取文件或文档的名字
func (s *Service) GetFileOrDocNames(ctx context.Context, req *pb.GetFileOrDocNamesRequest, res *pb.GetFileOrDocNamesResponse) error {
	// var names []string
	if req.Type == 1 { // file
		file, err := model.GetFileDetail(req.Id)
		if err != nil {
			return e.NotFoundErr(errno.ErrDatabase, err.Error())
		}
		res.Name = file.Name

	} else if req.Type == 2 { // doc
		doc, err := model.GetDocDetail(req.Id)
		if err != nil {
			return e.NotFoundErr(errno.ErrDatabase, err.Error())
		}
		res.Name = doc.Name
	} else {
		return e.BadRequestErr(errno.ErrGetDataFromRPC, "type == 1 -> file, 2 -> doc")
	}

	return nil
}

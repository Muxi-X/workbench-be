package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetFileDetail ... 获取文件详情
func (s *Service) GetFileDetail(ctx context.Context, req *pb.GetRequest, res *pb.FileDetail) error {

	file, err := model.GetFileDetail(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = file.ID
	res.Url = file.URL
	res.Creator = file.Creator
	res.CreateTime = file.CreateTime

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetDocDetail ... 获取文档详情
func (s *Service) GetDocDetail(ctx context.Context, req *pb.GetRequest, res *pb.DocDetail) error {

	doc, err := model.GetDocDetail(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Title = doc.Name
	res.Id = doc.ID
	res.Content = doc.Content
	res.Creator = doc.Creator
	res.LastEditor = doc.Editor
	res.CreateTime = doc.CreateTime
	res.LastEditTime = doc.LastEditTime // TODO LastEditTime 数据缺失

	return nil
}

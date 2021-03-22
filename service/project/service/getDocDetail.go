package service

import (
	"context"
	"fmt"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetDocDetail ... 获取文档详情
func (s *Service) GetDocDetail(ctx context.Context, req *pb.GetFileDetailRequest, res *pb.DocDetail) error {
	// 判断自己 id 和父 id 是否被删
	var target []string
	self := fmt.Sprintf("%d-%d", req.Id, constvar.DocCode)
	father := fmt.Sprintf("%d-%d", req.FatherId, constvar.DocFolderCode)
	target = append(target, self, father)
	isDeleted, err := m.SIsmembersFromRedis(constvar.Trashbin, target)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted {
		return e.BadRequestErr(errno.ErrDatabase, "This file has been deleted.")
	}

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
	res.LastEditTime = doc.LastEditTime
	return nil
}

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

// GetFileDetail ... 获取文件详情
func (s *Service) GetFileDetail(ctx context.Context, req *pb.GetFileDetailRequest, res *pb.FileDetail) error {
	// 判断自己 id 和父 id 是否被删
	var target []string
	self := fmt.Sprintf("%d-%d", req.Id, constvar.FileCode)
	father := fmt.Sprintf("%d-%d", req.FatherId, constvar.FileFolderCode)
	target = append(target, self, father)
	isDeleted, err := m.SIsmembersFromRedis(constvar.Trashbin, target)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted {
		return e.BadRequestErr(errno.ErrDatabase, "This file has been deleted.")
	}

	// ok
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

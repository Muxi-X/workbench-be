package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateFile ... 创建文件
func (s *Service) CreateFile(ctx context.Context, req *pb.CreateFileRequest, res *pb.ProjectIDResponse) error {

	t := time.Now()

	file := model.FileModel{
		Name:       req.HashName,
		RealName:   req.Name,
		Re:         false,
		Top:        false,
		TeamID:     req.TeamId, // 查询一下用户信息，user 服务 rpc
		CreateTime: t.Format("2006-01-02 15:04:05"),
		ProjectID:  req.ProjectId,
		URL:        req.Url,
	}

	id, err := model.CreateFile(m.DB.Self, &file, req.FatherId, req.FatherType)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = id

	return nil
}

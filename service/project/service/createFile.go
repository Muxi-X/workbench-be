package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateFile ... 创建文件
func (s *Service) CreateFile(ctx context.Context, req *pb.CreateFileRequest, res *pb.Response) error {

	t := time.Now()

	file := model.FileModel{
		Name:       req.HashName,
		RealName:   req.Name,
		Re:         false,
		Top:        false,
		TeamID:     0, // 查询一下用户信息，user 服务 rpc
		CreateTime: t.Format("2006-01-02 15:04:05"),
		ProjectID:  req.ProjectId,
		URL:        req.Url,
	}

	if err := file.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

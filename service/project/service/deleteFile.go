package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteFile ... 删除文件
func (s *Service) DeleteFile(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	// TODO：软删除，DB 要添加 deleted_at 字段
	if err := model.DeleteFile(req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

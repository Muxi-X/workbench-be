package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteTrashbin ... 从回收站删除文件
// 修改回收站 re 即可
// 删除回收站记录、修改原表 re 字段、同步删除 redis 给协程
func (s *Service) DeleteTrashbin(ctx context.Context, req *pb.DeleteTrashbinRequest, res *pb.Response) error {
	if err := model.DeleteTrashbin(req.Id, uint8(req.Type)); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

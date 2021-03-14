package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetDocChildren ... 获取任意文档夹目录下的文档树
func (s *Service) GetDocChildren(ctx context.Context, req *pb.GetRequest, res *pb.Children) error {

	item, err := model.GetDocChildrenById(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Children = item.Children

	return nil
}

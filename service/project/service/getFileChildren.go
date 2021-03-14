package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetFileChildren ... 获取任意文件夹目录下的文件树
func (s *Service) GetFileChildren(ctx context.Context, req *pb.GetRequest, res *pb.Children) error {

	item, err := model.GetFileChildrenById(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Children = item.Children

	return nil
}

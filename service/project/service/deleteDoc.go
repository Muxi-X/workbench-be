package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteDoc ... 删除文档
// 所有的 delete 都需要前端调用 updareTree 的 api
// 也就是文件删一次，文件树里删一次
func (s *Service) DeleteDoc(ctx context.Context, req *pb.GetRequest, res *pb.ProjectIDResponse) error {
	// 先查找再删除
	item, err := model.GetDoc(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	item.Re = true
	res.Id = item.ProjectID

	// TODO：软删除，DB 要添加 deleted_at 字段
	if err := item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteDocComment ... 删除文档评论
func (s *Service) DeleteDocComment(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	// TODO：软删除，DB 要添加 deleted_at 字段（待评论拆表之后处理）
	if err := model.DeleteComment(req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

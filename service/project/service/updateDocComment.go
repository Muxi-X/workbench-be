package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

func (s *Service) UpdateDocComment(ctx context.Context, req *pb.UpdateDocCommentRequest, res *pb.Response) error {
	item, err := model.GetCommentModelById(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	item.Content = req.Content

	if err = item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// CreateApplication …… 生成申请
func (ts *TeamService) CreateApplication(ctx context.Context, req *pb.ApplicationRequest, res *pb.Response) error {
	apply := &model.ApplyModel{
		UserID: req.UserId,
	}

	if err := apply.Check(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	if err := apply.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

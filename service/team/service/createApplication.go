package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	"github.com/Muxi-X/workbench-be/service/team/errno"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
)

func (ts *TeamService) CreateApplication (ctx context.Context, req *pb.ApplicationRequest, res *pb.Response) error {
	apply := &model.ApplyModel{
		UserID: req.UserId,
	}
	if err := apply.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

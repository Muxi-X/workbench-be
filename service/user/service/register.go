package service

import (
	"context"
	errno "muxi-workbench-user/errno"
	e "muxi-workbench/pkg/err"
	model "muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
)

// Register ... 注册
func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest, res *pb.Response) error {

	user := &model.UserModel{
		Name: req.Name,
		Email: req.Email,
	}

	if err := user.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

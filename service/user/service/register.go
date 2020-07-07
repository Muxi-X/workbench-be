package service

import (
	"context"

	errno "muxi-workbench-user/errno"
	model "muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
)

// Register ... 注册
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest, res *pb.Response) error {

	user := &model.UserModel{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := user.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

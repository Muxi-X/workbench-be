package service

import (
	"context"

	"muxi-workbench-user/errno"
	"muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateInfo ... 更新用户信息
func (s *UserService) UpdateInfo(ctx context.Context, req *pb.UpdateInfoRequest, res *pb.Response) error {
	user, err := model.GetUser(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	if user == nil {
		return e.ServerErr(errno.ErrUserExisted, err.Error())
	}

	user.Name = req.Info.Nick
	user.RealName = req.Info.Name
	user.Avatar = req.Info.AvatarUrl

	if err := user.Save(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-user/errno"
	"muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
)

// GetProfile ... 获取用户个人信息
func (s *UserService) GetProfile(ctx context.Context, req *pb.GetRequest, res *pb.UserProfile) error {
	user, err := model.GetUser(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	if user == nil {
		return e.ServerErr(errno.ErrUserNotExisted, "")

	}
	res.Id = user.ID
	res.Name = user.Name
	res.RealName = user.RealName
	res.Avatar = user.Avatar
	res.Email = user.Email
	res.Tel = user.Tel
	res.Role = user.Role
	res.Team = user.TeamID
	res.Group = user.GroupID

	return nil
}

package service

import (
	"context"

	errno "muxi-workbench-user/errno"
	model "muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
)

// GetInfo ... 获取用户信息
func (s *UserService) GetInfo(ctx context.Context, req *pb.GetInfoRequest, res *pb.UserInfoResponse) error {

	list, err := model.GetUserByIds(req.Ids)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	userInfos := make([]*pb.UserInfo, 0)

	for _, user := range list {
		userInfos = append(userInfos, &pb.UserInfo{
			Id:        user.ID,
			Name:      user.Name,
			RealName:  user.RealName,
			AvatarUrl: user.Avatar,
			Email:     user.Email,
		})
	}

	res.List = userInfos

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-user/errno"
	model "muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
)

// GetInfo ... 获取用户信息
func (s *Service) GetInfo(ctx context.Context, req *pb.GetInfoRequest, res *pb.UserInfoResponse) error {

	list, err := model.GetUserByIds(req.Ids)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	userInfos := make([]*pb.UserInfo, 0)

	for index := 0; index < len(list); index++ {
		user := list[index]
		userInfos = append(userInfos, &pb.UserInfo{
			Id:        user.ID,
			Nick:      user.Name,
			Name:      user.RealName,
			AvatarUrl: user.Avatar,
		})
	}

	res.List = userInfos

	return nil
}

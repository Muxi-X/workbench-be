package service

import (
	"context"

	errno "muxi-workbench-user/errno"
	"muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
)

// List ... 获取用户列表
func (s *UserService) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {

	filter := &model.UserModel{TeamID: req.Team}
	if req.Group != 0 {
		filter.GroupID = req.Group
	}

	list, err := model.ListUser(req.Offset, req.Limit, req.LastId, filter)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.User, 0)

	for _, item := range list {
		resList = append(resList, &pb.User{
			Id:     item.ID,
			Nick:   item.Name,
			Name:   item.RealName,
			Email:  item.Email,
			Avatar: item.Avatar,
			Role:   item.Role,
			Team:   item.TeamID,
			Group:  item.GroupID,
		})
	}

	res.Count = uint32(len(list))
	res.List = resList

	return nil
}

package service

import (
	"context"
	"muxi-workbench-team/errno"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

type MemberModel struct {
	UserID    uint32
	TeamID    uint32
	GroupID   uint32
	GroupName string
	Role      uint32
	Email     string
	Avatar    string
	Name      string
}

// GetMemberList …… 组别内成员列表
func (ts *TeamService) GetMemberList(ctx context.Context, req *pb.MemberListRequest, res *pb.MemberListResponse) error {
	list, count, err := GetMemberInfo(req.GroupId, req.Limit, req.Offset, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.Member, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.Member{
			Id:        item.UserID,
			Name:      item.Name,
			TeamId:    item.TeamID,
			GroupId:   item.GroupID,
			GroupName: item.GroupName,
			Email:     item.Email,
			Role:      item.Role,
			Avatar:    item.Avatar,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

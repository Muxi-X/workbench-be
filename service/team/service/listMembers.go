package service

import (
	"context"
	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

//list all members of a group …… 组别内成员列表
func (ts *TeamService) GetMemberList(ctx context.Context, req *pb.MemberListRequest, res *pb.MemberListResponse) error {
	list, count, err := model.GetMemberInfo(req.GroupId, req.Limit, req.Offset, req.Pagination)
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
			Role:      item.Role,
			Email:     item.Email,
			Avatar:    item.Avatar,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

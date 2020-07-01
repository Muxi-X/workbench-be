package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
)

//list … 组别成员列表
func (ts *TeamService) GetMemberList (ctx context.Context,req *pb.MemberListRequest,res *pb.MemberListResponse) error {
	list, count, err := model.ListMembersOfAGroup(req.GroupId, req.Limit, req.Offset, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.Member, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.Member{
			Id:                   item.ID,
			Name:                 item.Name,
			TeamId:               item.TeamID,
			GroupId:              item.GroupID,
			GroupName:            item.GroupName,
			Role:                 item.Role,
			Email:                item.Email,
			Avatar:               item.Avatar,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

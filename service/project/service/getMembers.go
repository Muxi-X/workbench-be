package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetMembers ... 获取项目的成员列表
func (s *Service) GetMembers(ctx context.Context, req *pb.GetMemberListRequest, res *pb.MembersListResponse) error {

	list, _, err := model.GetProjectMemberList(req.ProjectId, req.Offset, req.Limit, req.Lastid, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.MembersListItem, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.MembersListItem{
			Id:        item.ID,
			Name:      item.Name,
			Avatar:    item.Avatar,
			GroupName: "groupName",
			Role:      item.Role,
		})
	}

	res.List = resList

	return nil
}

package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// GetGroupName …… 根据groupID获取groupName
func (ts *TeamService) GetGroupName(ctx context.Context, req *pb.GroupRequest, res *pb.GroupNameResponse) error {

	group, err := model.GetGroup(req.GroupId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	res.GroupName = group.Name
	return nil
}

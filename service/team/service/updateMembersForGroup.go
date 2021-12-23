package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateMembersForGroup …… 更新组别内成员信息
func (ts *TeamService) UpdateMembersForGroup(ctx context.Context, req *pb.UpdateMembersRequest, res *pb.Response) error {
	g, err := model.GetGroup(req.GroupId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	g.Count = uint32(len(req.UserList))
	if err := g.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// update members info
	if err := UpdateUsersGroupIDOrTeamID(req.UserList, req.GroupId, model.GROUP); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// update count in groups
	group, err := model.GetGroup(req.GroupId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	group.Count = uint32(len(req.UserList))
	return group.Update()
}

package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteGroup … 删除组别
func (ts *TeamService) DeleteGroup(ctx context.Context, req *pb.GroupRequest, res *pb.Response) error {
	// 获取usersId
	usersId, err := GetUsersIdByGroupId(req.GroupId)
	if err != nil {
		return e.ServerErr(errno.ErrClient, err.Error())
	}

	if err := UpdateUsersGroupIDOrTeamID(usersId, model.NOGROUP, model.GROUP); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 删除组别
	err = model.DeleteGroup(req.GroupId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

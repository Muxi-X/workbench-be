package service

import (
	"context"
	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

//Create … 建立组别
func (ts *TeamService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest, res *pb.Response) error {
	//判断权限
	if req.Role != model.SUPERADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}
	//构建组别信息,新建组别
	t := time.Now()
	group := &model.GroupModel{
		Name:    req.GroupName,
		Order:   0,
		Counter: uint32(len(req.UserList)),
		Leader:  0,
		Time:    t.Format("2006-01-02 15:04:05"),
	}
	if err := group.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	//获取新建好的组别id, 并且设置好组别内成员的group_id
	groupid, err := model.GetGroupId(req.GroupName)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if err := model.UpdateUsersGroupidOrTeamid(req.UserList, groupid, model.GROUP); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

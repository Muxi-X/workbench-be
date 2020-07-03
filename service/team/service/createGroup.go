package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
	"time"
)

//Create … 建立组别
func (ts *TeamService) Create(ctx context.Context, req *pb.CreateGroupRequest, res *pb.Response) error {
	if req.Role != model.SUPERADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}

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

    groupid, err := model.GetGroupId(req.GroupName)
    if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}



/*
	//second update user's info
	list, err := um.GetUserByIds(req.UserList)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	for _, member := range list {
		member.GroupID = group.ID
		//if err := member.Update(); err != nil {
		//     return e.ServerErr(errno.ErrDatabase, err.Error())
		//}
	}
*/

	return nil
}

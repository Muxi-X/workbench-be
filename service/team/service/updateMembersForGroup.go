package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
)

func (ts *TeamService) UpdateMembersForGroup (ctx context.Context, req *pb.UpdateMembersRequest, res *pb.Response) error {
	//judge the role
	if req.Role != model.SUPERADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}
	//传入的为该组别全体的成员,统计切片长度,即为总人数..之后更新数据
	g, err := model.GetGroup(req.GroupId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	g.Counter = uint32(len(req.UserList))
	if err := g.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase,err.Error())
	}

	//update members info
	if err := UpdateUsersGroupidOrTeamid(req.UserList, req.GroupId, model.GROUP); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}


	return nil
}

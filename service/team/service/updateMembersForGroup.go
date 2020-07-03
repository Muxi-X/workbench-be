package service

import (
	"context"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	um "github.com/Muxi-X/workbench-be/service/user/model"
)

func (ts *TeamService) UpdateMembersForGroup (ctx context.Context, req *pb.UpdateMembersRequest, res *pb.Response) error {
	//judge the role
	if req.Role != model.SUPERADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}


	g, err := model.GetGroup(req.GroupId)
	//用零暂时代替表达式.以后要做更改
	g.Counter = 0
	g.Update()


	//update members info
	list, err := um.GetUserByIds(req.UserList)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	for _, member := range list {
		member.GroupID = req.GroupId
		//if err := member.Update(); err != nil {
		//     return e.ServerErr(errno.ErrDatabase, err.Error())
		//}
	}


	return nil
}

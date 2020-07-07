package service

import (
	"context"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"

	e "github.com/Muxi-X/workbench-be/pkg/err"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
)

//Delete … 删除组别
func (ts *TeamService) Delete (ctx context.Context, req *pb.DeleteGroupRequest, res *pb.Response) error {
	//判断权限
	if req.Role != model.SUPERADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}
	//删除组别
	err := model.DeleteGroup(req.GroupId)
	if err != nil{
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
    //获取usersid
	usersid, err := GetUsersId(req.GroupId)
	if err := UpdateUsersGroupidOrTeamid(usersid, model.NOGROUP, model.GROUP); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
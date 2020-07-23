package service

import (
	"context"
	"log"
	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

//Delete … 删除组别
func (ts *TeamService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest, res *pb.Response) error {
	//判断权限
	if req.Role != model.SUPERADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}

	//获取usersid
	usersid, err := model.GetUsersIdByGroupid(req.GroupId)
	if err := model.UpdateUsersGroupidOrTeamid(usersid, model.NOGROUP, model.GROUP); err != nil {
		log.Println(err)
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	//删除组别
	err = model.DeleteGroup(req.GroupId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}


	return nil
}

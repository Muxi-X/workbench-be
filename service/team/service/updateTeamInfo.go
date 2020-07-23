package service

import (
	"context"

	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)


//Update …… 更新组别信息
func (ts *TeamService) UpdateTeamInfo(ctx context.Context, req *pb.UpdateTeamInfoRequest, res *pb.Response) error {
	//判断权限
	if req.Role != model.SUPERADMIN && req.Role != model.ADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}
	//获取group结构体,用以更新对应的数据
	team, err := model.GetTeam(req.TeamId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	team.Name = req.NewName
	if err := team.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

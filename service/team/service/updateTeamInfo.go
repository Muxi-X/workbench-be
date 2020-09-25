package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateTeamInfo …… 更新团队信息
func (ts *TeamService) UpdateTeamInfo(ctx context.Context, req *pb.UpdateTeamInfoRequest, res *pb.Response) error {
	// 获取group结构体,用以更新对应的数据
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

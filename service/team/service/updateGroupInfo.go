package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateGroupInfo …… 更新组别信息
func (ts *TeamService) UpdateGroupInfo(ctx context.Context, req *pb.UpdateGroupInfoRequest, res *pb.Response) error {
	// 判断权限
	if req.Role != model.SUPERADMIN && req.Role != model.ADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}

	// 获取group结构体,用以更新对应的数据
	group, err := model.GetGroup(req.GroupId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	group.Name = req.NewName
	if err := group.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

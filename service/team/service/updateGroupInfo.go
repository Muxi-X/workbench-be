package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
)

func (ts *TeamService) UpdateGroupInfo (ctx context.Context, req *pb.UpdateGroupInfoRequest, res *pb.Response) error {
	if req.Role != model.SUPERADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}

	group, err := model.GetGroup(req.GroupId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase,err.Error())
	}

    group.Name =req.NewName

	if err := group.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase,err.Error())
	}

	return nil
}

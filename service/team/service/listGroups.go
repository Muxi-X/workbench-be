package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
)

func (ts *TeamService) GetGroupList(ctx context.Context,req *pb.GroupListRequest,res *pb.GroupListResponse) error {
	list, count, err := model.ListGroup(req.Offset, req.Limit, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	resList := make([]*pb.GroupItem, 0)

}
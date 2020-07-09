package service

import (
	"context"
	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

//List …… 列举申请
func (ts *TeamService) GetApplications(ctx context.Context, req *pb.ApplicationListRequest, res *pb.ApplicationListResponse) error {
	applys, count, err := model.ListApplys(req.Offset, req.Limit, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	list, _, err := model.GetUsersByApplys(applys, count)
	if err != nil {
		return e.ServerErr(errno.ErrClient, err.Error())
	}

	resList := make([]*pb.ApplyUserItem, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.ApplyUserItem{
			Id:    item.ID,
			Name:  item.Name,
			Email: item.Eamil,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil

}

package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
)

func (ts *TeamService) GetApplications(ctx context.Context,req *pb.ApplicationListRequest,res *pb.ApplicationListResponse) error {
	list, count, err := model.ListApplictions(req.Offset, req.Limit, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.ApplyUserItem, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.ApplyUserItem{
			Id:                   item.ID,
			Name:                 item.Name,
			Email:                item.Eamil,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

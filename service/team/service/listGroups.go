package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
)

//list … 组别列表
func (ts *TeamService) GetGroupList(ctx context.Context,req *pb.GroupListRequest,res *pb.GroupListResponse) error {
	list, count, err := model.ListGroup(req.Offset, req.Limit, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.GroupItem, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.GroupItem{
			Id:                   item.ID,
			Name:                 item.Name,
			UserCount:            item.Counter,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}
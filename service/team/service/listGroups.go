package service

import (
	"context"
	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

//list … 组别列表
func (ts *TeamService) GetGroupList(ctx context.Context, req *pb.GroupListRequest, res *pb.GroupListResponse) error {
	list, count, err := model.ListGroup(req.Offset, req.Limit, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.GroupItem, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.GroupItem{
			Id:        item.ID,
			Name:      item.Name,
			UserCount: item.Counter,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
)

// List ... 动态列表
func (s *StatusService) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	list, count, err := model.ListStatus(req.Group, req.Team, req.Offset, req.Limit, req.Lastid, &model.StatusModel{
		UserID: req.Uid,
	})
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	statusLikeList, err := model.GetStatusLikeRecordForUser(req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.Status, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.Status{
			Id:       item.ID,
			Content:  item.Content,
			Title:    item.Title,
			Time:     item.Time,
			Avatar:   item.Avatar,
			UserName: item.UserName,
		})
		for j := 0; j < len(statusLikeList); j++ {
			if statusLikeList[j].StatusID == item.ID {
				resList[index].IfLike = 1
				break // 删除此时节点
			}
		}
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

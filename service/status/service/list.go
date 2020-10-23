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
	list, count, err := model.ListStatus(req.Group, req.Team, req.Offset, req.Limit, req.LastId, &model.StatusModel{
		UserID: req.Uid,
	})
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	var scope []int
	for i := 0; i < len(list); i++ {
		scope = append(scope, int(list[i].ID))
	}

	statusLikeList, err := model.GetStatusLikeRecordForUser(req.UserId, scope)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.Status, 0)

	// 合并结果
	lenLike := len(statusLikeList)
	lenList := len(list)
	i := 0 // list 的指针
	j := 0 // like 的指针
	for i != lenList && j != lenLike {
		if statusLikeList[j].StatusID > list[i].ID { // 如果 likelist 当前 id 比 statuslist id 大，说明不在范围内，直接跳过。
			j++ // 只有 j 索引往后移动
			continue
		}
		if statusLikeList[j].StatusID == list[i].ID { // 如果 likelist 当前 id 等于 statuslist id ，是该用户点赞的记录，liked 设置 1。
			item := list[i]
			resList = append(resList, &pb.Status{
				Id:       item.ID,
				Content:  item.Content,
				Title:    item.Title,
				Time:     item.Time,
				Avatar:   item.Avatar,
				UserName: item.UserName,
				Liked:    true,
			})
			i++
			j++
			continue
		}
		if statusLikeList[j].StatusID < list[i].ID { // 如果 likelist 当前 id 小于 statuslist id，该记录不是目标记录，liked设置成 0。
			item := list[i]
			resList = append(resList, &pb.Status{
				Id:       item.ID,
				Content:  item.Content,
				Title:    item.Title,
				Time:     item.Time,
				Avatar:   item.Avatar,
				UserName: item.UserName,
				Liked:    false,
			})
			i++ // 索引 i 往后走，j 等待目标
			continue
		}
	}

	// 若 statuslist 没走完，需要把剩下的status插入结果
	if i < lenList {
		for i != lenList {
			item := list[i]
			resList = append(resList, &pb.Status{
				Id:       item.ID,
				Content:  item.Content,
				Title:    item.Title,
				Time:     item.Time,
				Avatar:   item.Avatar,
				UserName: item.UserName,
				Liked:    false,
			})
			i++
		}
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

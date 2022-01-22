package service

import (
	"context"
	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
)

// Search ... 搜索进度
func (s *StatusService) Search(ctx context.Context, req *pb.SearchRequest, res *pb.SearchResponse) error {
	var list []*model.SearchResult
	// 筛选条件
	filter := &model.FilterParams{
		UserName: req.UserName,
		GroupId:  req.GroupId,
		Key:      "%" + req.Keyword + "%",
	}

	contentList, count, err := model.Search(filter, req.Offset, req.Limit, req.LastId, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	list = append(list, contentList...)
	res.Count = count

	for _, item := range list {
		res.List = append(res.List, &pb.SearchResult{
			Id:       item.Id,
			Title:    item.Title,
			UserName: item.UserName,
			Content:  item.Content,
			Time:     item.Time,
		})
	}

	return nil
}

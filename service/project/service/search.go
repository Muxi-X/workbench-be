package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// Search ... 搜索文档和文件
// search_type: 为0时搜索doc title and content，为1时搜索file title
func (s *Service) Search(ctx context.Context, req *pb.SearchRequest, res *pb.SearchResponse) error {
	var projectIDs []uint32
	if req.ProjectId == 0 {
		projects, err := model.GetUserToProjectByUser(req.UserId)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}

		for _, project := range projects {
			projectIDs = append(projectIDs, project.ProjectID)
		}
	} else {
		projectIDs = append(projectIDs, req.ProjectId)
	}

	var list []*model.SearchResult

	if req.Type == 0 { // 在文档title and content中查询关键字
		titleList, count, err := model.SearchDoc(projectIDs, req.Keyword, req.Offset, req.Limit, req.LastId, req.Pagination)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}

		list = append(list, titleList...)
		res.Count = count

	} else if req.Type == 1 { // 在文件title中查询关键字
		contentList, count, err := model.SearchFile(projectIDs, req.Keyword, req.Offset, req.Limit, req.LastId, req.Pagination)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}

		list = append(list, contentList...)
		res.Count = count
	} else {
		return e.BadRequestErr(errno.ErrBind, "wrong type_id")
	}

	for _, item := range list {
		res.List = append(res.List, &pb.SearchResult{
			Id:          item.Id,
			Type:        uint32(item.Type),
			Title:       item.Title,
			UserName:    item.UserName,
			Content:     item.Content,
			ProjectName: item.ProjectName,
			Time:        item.Time,
		})
	}

	return nil
}

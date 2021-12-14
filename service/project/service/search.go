package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// Search ... 搜索文档和文件
// search_type: 为1时搜索doc_title and file_title，为2时搜索doc_content
func (s *Service) Search(ctx context.Context, req *pb.SearchRequest, res *pb.SearchResponse) error {
	projects, err := model.GetUserToProjectByUser(req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	projectIDs := make([]uint32, len(projects))
	for i, project := range projects {
		projectIDs[i] = project.ProjectID
	}
	var list []*model.SearchResult

	if req.Type == 1 { // 在文档文件title中查询关键字
		titleList, count, err := model.SearchTitle(projectIDs, req.Keyword, req.Offset, req.Limit, req.LastId, req.Pagination)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}

		list = append(list, titleList...)
		res.Count = count

		// } else if req.Type == 2 { // 在文档content中查询关键字
		// 	contentList, count, err := model.SearchContent(project.ProjectID, req.Keyword, req.Offset, req.Limit, req.LastId, req.Pagination)
		// 	if err != nil {
		// 		return e.ServerErr(errno.ErrDatabase, err.Error())
		// 	}
		//
		// 	list = append(list, contentList...)
		// 	res.Count = count
	}

	for _, item := range list {
		res.List = append(res.List, &pb.SearchResult{
			Id: item.Id,
			// Type:        item.Type,
			Title:       item.Title,
			UserName:    item.UserName,
			Content:     item.Content,
			ProjectName: item.ProjectName,
			Time:        item.Time,
		})
	}

	return nil
}

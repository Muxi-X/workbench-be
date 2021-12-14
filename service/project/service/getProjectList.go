package service

import (
	"context"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"

	errno "muxi-workbench-project/errno"
	e "muxi-workbench/pkg/err"
)

// GetProjectList ... 获取项目信息
func (s *Service) GetProjectList(ctx context.Context, req *pb.GetProjectListRequest, res *pb.ProjectListResponse) error {

	list, _, err := model.ListProject(req.UserId, req.Offset, req.Limit, req.Lastid, req.Pagination)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.ProjectListItem, 0)

	for _, item := range list {
		resList = append(resList, &pb.ProjectListItem{
			Id:   item.ID,
			Name: item.Name,
			Logo: "",
		})
	}

	res.List = resList

	return nil
}

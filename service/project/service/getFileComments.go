package service

import (
	"context"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"

	errno "muxi-workbench-project/errno"
	e "muxi-workbench/pkg/err"
)

// GetFileComments ... 获取文件评论列表
func (s *Service) GetFileComments(ctx context.Context, req *pb.GetCommentListRequest, res *pb.GetCommentListResponse) error {

	list, _, err := model.ListComments("docId", req.Id, req.Offset, req.Limit, req.Lastid)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.Comment, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.Comment{
			Id:         item.ID,
			Content:    item.Content,
			UserName:   item.UserName,
			Avatar:     item.Avatar,
			CreateTime: item.Time,
			UserId:     item.Creator,
		})
	}

	res.List = resList

	return nil
}

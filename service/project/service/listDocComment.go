package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// ListDocComment ... 获取文档评论列表
func (s *Service) ListDocComment(ctx context.Context, req *pb.ListDocCommentRequest, res *pb.CommentListResponse) error {
	list, count, err := model.ListDocComments(req.DocId, req.Offset, req.Limit, req.LastId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.Comment, 0)

	for index := 0; index < len(list); index++ {
		item := list[index]
		resList = append(resList, &pb.Comment{
			Id:       item.ID,
			Content:  item.Content,
			Time:     item.Time,
			Avatar:   item.Avatar,
			UserName: item.UserName,
			UserId:   item.Creator,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
)

// ListComment ... 评论列表
func (s *StatusService) ListComment(ctx context.Context, req *pb.CommentListRequest, res *pb.CommentListResponse) error {

	list, count, err := model.ListComments(req.StatusId, req.Offset, req.Limit, req.LastId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.Comment, 0)

	for _, item := range list {
		resList = append(resList, &pb.Comment{
			Id:       item.ID,
			Content:  item.Content,
			Time:     item.Time,
			Avatar:   item.Avatar,
			UserName: item.UserName,
			UserId:   item.Creator,
			Kind:     item.Kind,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetCommentList ... 获取文档/文件评论列表
func (s *Service) GetCommentList(ctx context.Context, req *pb.GetCommentRequest, res *pb.CommentListResponse) error {
	comment, err := getCommentType(uint8(req.TypeId))
	if err != nil {
		return e.BadRequestErr(errno.ErrBind, err.Error())
	}

	list, count, err := comment.List(req.TargetId, req.Offset, req.Limit, req.LastId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.Comment, 0)

	for _, item := range list {
		resList = append(resList, &pb.Comment{
			Id:       item.ID,
			Kind:     item.Kind,
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

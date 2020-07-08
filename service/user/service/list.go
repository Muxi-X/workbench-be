package service

import (
	"context"
	pb "muxi-workbench-user/proto"
)

// List ... 获取用户列表
func (s *Service) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {

	// status, err := model.GetStatus(req.Id)
	// if err != nil {
	// 	return e.ServerErr(errno.ErrDatabase, err.Error())
	// }

	// res.Status = &pb.Status{
	// 	Id:      status.ID,
	// 	Title:   status.Title,
	// 	Content: status.Content,
	// 	UserId:  status.UserID,
	// 	Time:    status.Time,
	// }

	return nil
}

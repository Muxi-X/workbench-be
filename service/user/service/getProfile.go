package service

import (
	"context"
	pb "muxi-workbench-user/proto"
)

// GetProfile ... 获取用户信息
func (s *Service) GetProfile(ctx context.Context, req *pb.GetRequest, res *pb.UserProfile) error {

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

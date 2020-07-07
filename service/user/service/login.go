package service

import (
	"context"
	pb "muxi-workbench-user/proto"
)

// Login ... 登录
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest, res *pb.Response) error {

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

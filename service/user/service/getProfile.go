package service

import (
	"context"
	pb "muxi-workbench-user/proto"
)

// GetProfile ... 获取用户信息
func (s *UserService) GetProfile(ctx context.Context, req *pb.GetRequest, res *pb.UserProfile) error {
	return nil
}

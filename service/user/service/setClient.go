package service

import (
	"context"

	"muxi-workbench-user/pkg/auth"
	pb "muxi-workbench-user/proto"
)

// SetClient set the client id and secret.
func (s *UserService) SetClient(ctx context.Context, req *pb.ClientRequest, res *pb.Response) error {
	auth.OauthManager.SetClient(req.Id, req.Secret)
	return nil
}

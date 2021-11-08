package service

import (
	"context"
	"muxi-workbench-attention/model"
	pb "muxi-workbench-attention/proto"
)

// Delete ... 删除attention
func (s *AttentionService) Delete(ctx context.Context, req *pb.PushRequest, res *pb.Response) error {

	attention := &model.AttentionModel{
		UserId: req.UserId,
		DocId:  req.DocId,
	}

	if err := attention.Delete(); err != nil {
		return err
	}

	return nil
}

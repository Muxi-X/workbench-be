package service

import (
	"context"
	"time"

	"muxi-workbench-attention/model"
	pb "muxi-workbench-attention/proto"
)

// Create ... 新增attention
func (s *AttentionService) Create(ctx context.Context, req *pb.PushRequest, res *pb.Response) error {
	// TODO: 判断docId存在 与 禁止重复创建
	attention := &model.AttentionModel{
		UserId:  req.UserId,
		DocId:   req.DocId,
		TimeDay: time.Now().Format("2006/01/02"),
		TimeHm:  time.Now().Format("15:04"),
	}

	if err := attention.Create(); err != nil {
		return err
	}

	return nil
}

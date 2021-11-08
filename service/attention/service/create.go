package service

import (
	"context"
	"github.com/pkg/errors"
	"time"

	"muxi-workbench-attention/model"
	pb "muxi-workbench-attention/proto"
)

// Create ... 新增attention
func (s *AttentionService) Create(ctx context.Context, req *pb.PushRequest, res *pb.Response) error {
	// TODO: 判断docId存在
	attention := &model.AttentionModel{
		UserId:  req.UserId,
		DocId:   req.DocId,
		TimeDay: time.Now().Format("2006/01/02"),
		TimeHm:  time.Now().Format("15:04"),
	}
	if attention.GetByUserAndDoc(); attention.Id != 0 {
		return errors.New("this attention already exists")
	}
	if err := attention.Create(); err != nil {
		return err
	}

	return nil
}

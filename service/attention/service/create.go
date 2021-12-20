package service

import (
	"context"
	"github.com/pkg/errors"
	"muxi-workbench/pkg/constvar"
	"time"

	"muxi-workbench-attention/model"
	pb "muxi-workbench-attention/proto"
)

// Create ... 新增attention
func (s *AttentionService) Create(ctx context.Context, req *pb.PushRequest, res *pb.Response) error {
	// 判断fileId存在
	if req.FileKind == uint32(constvar.DocCode) {
		if _, err := model.GetDocDetail(req.FileId); err != nil {
			return err
		}
	} else if req.FileKind == uint32(constvar.FileCode) {
		if _, err := model.GetFileDetail(req.FileId); err != nil {
			return err
		}
	} else {
		errors.New("file_kind wrong : 1 -> doc, 2 -> file")
	}
	attention := &model.AttentionModel{
		UserId:   req.UserId,
		FileId:   req.FileId,
		TimeDay:  time.Now().Format("2006/01/02"),
		TimeHm:   time.Now().Format("15:04"),
		FileKind: req.FileKind,
	}
	if attention.GetByUserAndFile(); attention.Id != 0 {
		return errors.New("this attention already exists")
	}
	if err := attention.Create(); err != nil {
		return err
	}

	return nil
}

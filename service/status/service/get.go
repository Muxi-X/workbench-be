package service

import (
	"context"
	"errors"

	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"

	"github.com/jinzhu/gorm"
)

// Get ... 获取动态
func (s *StatusService) Get(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {

	status, err := model.GetStatus(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 获取 liked
	liked := true
	err = model.GetStatusLikeRecord(req.Uid, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			liked = false
		} else {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	}

	res.Status = &pb.Status{
		Id:      status.ID,
		Title:   status.Title,
		Content: status.Content,
		UserId:  status.UserID,
		Time:    status.Time,
		Like:    status.Like,
		Liked:   liked,
	}

	return nil
}

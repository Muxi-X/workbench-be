package service

import (
	"context"
	"errors"

	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	m "muxi-workbench/model"
	e "muxi-workbench/pkg/err"

	"github.com/jinzhu/gorm"
)

// Like ... 点赞动态
func (s *StatusService) Like(ctx context.Context, req *pb.LikeRequest, res *pb.Response) error {

	var notFound int

	record, err := model.GetStatusLikeRecord(req.UserId, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			notFound = 1
		} else {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	}

	status, err := model.GetStatus(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 加事物
	if notFound == 1 {
		status.Like = status.Like + 1
		record.UserID = req.UserId
		record.StatusID = req.Id

		err = model.AddStatusLike(m.DB.Self, record, status)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}

	} else {
		status.Like = status.Like - 1

		err = model.CancelStatusLike(m.DB.Self, record, status)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}

	}

	return nil
}

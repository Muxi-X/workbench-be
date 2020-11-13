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

	liked := true

	err := model.GetStatusLikeRecord(req.UserId, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			liked = false
		} else {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	}

	if liked == true && req.Liked == false {
		return e.ServerErr(errno.ErrDuplicateStatusLike, "")
	}

	if liked == false && req.Liked == true {
		return e.ServerErr(errno.ErrNoStatusLikeRecord, "")
	}

	status, err := model.GetStatus(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 加事务
	if !req.Liked {
		status.Like = status.Like + 1

		err = model.AddStatusLike(m.DB.Self, req.UserId, status)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}

	} else {
		status.Like = status.Like - 1

		err = model.CancelStatusLike(m.DB.Self, int(req.UserId), status)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}

	}

	return nil
}

package service

import (
	"context"
	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
)

// Like ... 点赞动态
func (s *StatusService) Like(ctx context.Context, req *pb.LikeRequest, res *pb.Response) error {

	record, err := model.GetStatusLikeRecord(req.UserId, req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	status, err := model.GetStatus(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	if record.ID == 0 {
		status.Like = status.Like + 1
		record.UserID = req.UserId
		record.StatusID = req.Id
		err = record.Create()
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	} else {
		status.Like = status.Like - 1
		err = model.DeleteStatusLikeRecord(record.UserID, record.StatusID, record.ID)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	}

	if err := status.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

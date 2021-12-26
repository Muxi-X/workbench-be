package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateFolder ... 建立folder
func (s *Service) CreateFolder(ctx context.Context, req *pb.CreateFolderRequest, res *pb.ProjectIDResponse) error {
	t := time.Now()

	folder := &model.FolderModel{
		Name:       req.Name,
		Re:         false,
		CreateTime: t.Format("2006-01-02 15:04:05"),
		CreatorID:  req.CreatorId,
		ProjectID:  req.ProjectId,
		FatherId:   req.FatherId,
	}

	id, err := model.CreateFolder(*folder, req.ChildrenPositionIndex, uint8(req.TypeId))
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = id

	return nil
}

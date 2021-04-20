package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateFileFolder ... 建立文件夹
func (s *Service) CreateFileFolder(ctx context.Context, req *pb.CreateFolderRequest, res *pb.ProjectIDResponse) error {
	t := time.Now()

	folder := &model.FolderForFileModel{
		Name:       req.Name,
		Re:         false,
		CreateTime: t.Format("2006-01-02 15:04:05"),
		CreatorID:  req.CreatorId,
		ProjectID:  req.ProjectId,
		FatherId:   req.FatherId,
	}

	id, err := model.CreateFileFolder(m.DB.Self, folder, req.ChildrenPositionIndex)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = id

	return nil
}

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

// CreateDocFolder ... 建立文档夹
func (s *Service) CreateDocFolder(ctx context.Context, req *pb.CreateFolderRequest, res *pb.ProjectIDResponse) error {
	t := time.Now()

	folder := &model.FolderForDocModel{
		Name:       req.Name,
		Re:         false,
		CreateTime: t.Format("2006-01-02 15:04:05"),
		CreatorID:  req.CreatorId,
		ProjectID:  req.ProjectId,
	}

	id, err := model.CreateDocFolder(m.DB.Self, folder, req.FatherId, req.FatherType)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = id

	return nil
}

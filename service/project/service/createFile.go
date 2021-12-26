package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	e "muxi-workbench/pkg/err"
	"strings"
	"time"
)

// CreateFile ... 创建文件
func (s *Service) CreateFile(ctx context.Context, req *pb.CreateFileRequest, res *pb.ProjectIDResponse) error {

	t := time.Now()
	index := strings.LastIndex(req.Url, "/")
	hashName := req.Url[index+1:]

	file := model.FileModel{
		CreatorID:  req.UserId,
		Name:       hashName,
		RealName:   req.Name,
		Re:         false,
		Top:        false,
		CreateTime: t.Format("2006-01-02 15:04:05"),
		ProjectID:  req.ProjectId,
		URL:        req.Url,
		FatherId:   req.FatherId,
	}

	id, err := model.CreateFile(m.DB.Self, &file, req.ChildrenPositionIndex)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = id

	return nil
}

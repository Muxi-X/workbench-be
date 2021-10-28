package service

import (
	"context"
	"encoding/json"
	"time"

	"muxi-workbench-attention/errno"
	"muxi-workbench-attention/model"
	pb "muxi-workbench-attention/proto"
	logger "muxi-workbench/log"
	e "muxi-workbench/pkg/err"
)

// Create ... 新增attention
func (s *AttentionService) Create(ctx context.Context, req *pb.PushRequest, res *pb.Response) error {
	// get username by userId from user-service
	userName, err := GetInfoFromUserService(req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
	}

	doc, err := GetInfoFromProjectService(req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
	}

	attention := &model.AttentionModel{
		UserId:         req.UserId,
		Username:       userName,
		DocId:          req.DocId,
		DocName:        doc.Name,
		DocCreatorId:   0,
		DocCreatorName: "",
		DocProjectName: "",
		DocProjectId:   doc.ProjectId,
		TimeDay:        time.Now().Format("2006/01/02"),
		TimeHm:         time.Now().Format("15:04"),
	}

	msg, err := json.Marshal(attention)
	if err != nil {
		return e.ServerErr(errno.ErrJsonMarshal, err.Error())
	}

	if err := model.PublishMsg(msg); err != nil {
		logger.Error("Publish data error")
		return e.ServerErr(errno.ErrPublishMsg, err.Error())
	}
	logger.Info("Publish data OK")
	return nil
}

package service

import (
	"context"

	"muxi-workbench-attention/errno"
	"muxi-workbench-attention/model"
	pb "muxi-workbench-attention/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// List ... attention列表
func (s *AttentionService) List(ctx context.Context, req *pb.ListRequest, res *pb.AttentionListResponse) error {
	// get username by userId from user-service
	userName, err := GetInfoFromUserService(req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
	}

	// 筛选条件
	var filter = &model.FilterParams{
		UserId: req.UserId,
	}

	attentions, err := model.List(req.LastId, req.Limit, filter)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	for _, a := range attentions {
		var file = &model.File{}
		if a.Kind == uint32(constvar.DocCode) {
			file, err = GetDocFromProjectService(a.File.Id)
		} else if a.Kind == uint32(constvar.FileCode) {
			file, err = GetFileFromProjectService(a.File.Id)
		}

		if err != nil {
			return err
		} else {
			a.File = *file
		}
	}

	// 数据格式化
	list, err := FormatListData(attentions, userName)
	if err != nil {
		return e.ServerErr(errno.ErrFormatList, err.Error())
	}

	res.List = list
	res.Count = uint32(len(list))

	return nil
}

func FormatListData(list []*model.AttentionDetail, username string) ([]*pb.AttentionItem, error) {
	var result []*pb.AttentionItem

	for _, attention := range list {
		data := &pb.AttentionItem{
			Date: attention.TimeDay,
			Time: attention.TimeHm,

			User: &pb.User{
				Name: username,
			},
			File: &pb.File{
				Name: attention.File.Name,
				FileCreator: &pb.User{
					Name: attention.File.CreatorName,
				},
				Kind:        attention.Kind,
				ProjectName: attention.File.ProjectName,
			},
		}
		result = append(result, data)
	}

	return result, nil
}

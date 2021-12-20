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
			file, err = model.GetDocDetail(a.File.Id)
		} else if a.Kind == uint32(constvar.FileCode) {
			file, err = model.GetFileDetail(a.File.Id)
		}
		file.Kind = a.Kind

		if err != nil {
			return err
		} else {
			a.File = *file
		}
	}

	// 数据格式化
	list, err := FormatListData(attentions)
	if err != nil {
		return e.ServerErr(errno.ErrFormatList, err.Error())
	}

	res.List = list
	res.Count = uint32(len(list))

	return nil
}

func FormatListData(list []*model.AttentionDetail) ([]*pb.AttentionItem, error) {
	var result []*pb.AttentionItem

	for _, attention := range list {
		data := &pb.AttentionItem{
			Date: attention.TimeDay,
			Time: attention.TimeHm,
			User: &pb.User{
				Name: attention.UserName,
			},
			File: &pb.File{
				Id:        attention.File.Id,
				ProjectId: attention.ProjectId,
				Name:      attention.File.Name,
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

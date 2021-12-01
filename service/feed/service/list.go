package service

import (
	"context"

	"muxi-workbench-feed/errno"
	"muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	e "muxi-workbench/pkg/err"
)

const (
	NOBODY     = 0 // 无权限用户
	NORMAL     = 1 // 普通用户
	ADMIN      = 3 // 管理员
	SUPERADMIN = 7 // 超管
)

// List ... feed列表
func (s *FeedService) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	var projectIds []uint32
	var err error

	// 普通用户，只能返回有权限访问的 projects
	if req.Role == NORMAL {
		projectIds, err = GetFilterFromProjectService(req.UserId)
		if err != nil {
			return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
		}
	}

	// 筛选条件
	var filter = &model.FilterParams{
		UserId:     req.Filter.UserId,
		GroupId:    req.Filter.GroupId,
		ProjectIds: projectIds,
	}

	feeds, err := model.GetFeedList(req.LastId, req.Limit, filter)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 数据格式化
	list, err := FormatListData(feeds)
	if err != nil {
		return e.ServerErr(errno.ErrFormatList, err.Error())
	}

	res.List = list
	res.Count = uint32(len(list))

	return nil
}

func FormatListData(list []*model.FeedModel) ([]*pb.FeedItem, error) {
	var result []*pb.FeedItem
	var date string
	var sourceId uint32

	for index, feed := range list {

		data := &pb.FeedItem{
			Id:          feed.Id,
			Action:      feed.Action,
			ShowDivider: false,
			Date:        feed.TimeDay,
			Time:        feed.TimeHm,
			User: &pb.User{
				Name:      feed.UserName,
				Id:        feed.UserId,
				AvatarUrl: feed.UserAvatar,
			},
			Source: &pb.Source{
				Kind:        feed.SourceKindId,
				Id:          feed.SourceObjectId,
				Name:        feed.SourceObjectName,
				ProjectId:   feed.SourceProjectId,
				ProjectName: feed.SourceProjectName,
			},
		}

		// showDivider --> 分割线
		// 需要分割的情况
		// 1.第一条数据 2.不同日期 3.不同项目
		if index == 0 {
			date = data.Date
			sourceId = data.Source.Id
			data.ShowDivider = true
		} else if date != data.Date {
			date = data.Date
			data.ShowDivider = true
		} else if sourceId != data.Source.Id {
			sourceId = data.Source.Id
			data.ShowDivider = true
		}

		result = append(result, data)
	}

	return result, nil
}

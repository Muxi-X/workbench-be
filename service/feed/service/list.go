package service

import (
	"context"
	"sync"

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

// 全部feed列表
func (s *FeedService) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	// 用户无权限
	if req.Role == NOBODY {
		return nil
	}

	// 获取feed数据
	list, err := model.GetFeedList(req.LastId, req.Size)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 格式化feed数据
	dataList, err := FormatListData(list)
	if err != nil {
		return e.ServerErr(errno.ErrFormatList, err.Error())
	}

	// 获取数据总数
	rows, err := model.GetRowsSum()
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 最大页数
	remain := rows % req.Size
	pageMax := rows / req.Size
	if remain != 0 {
		pageMax += 1
	}

	// 是否有下一页
	var hasNext = true
	if req.Page >= pageMax {
		hasNext = false
	}

	res.DataList = dataList
	res.HasNext = hasNext
	res.PageMax = pageMax
	res.PageNum = req.Page
	res.RowsNum = uint32(len(list))

	// TO DO:
	// 为管理员，则返回所有数据
	// 为普通用户，则查询用户所在的project ids,从所有的data中删去不在的数据，再返回
	// 如果直接删去不需要的数据，则返回的数据数不与请求的page size一致，最好还是sql查询时就进行该过程
	// 涉及到project服务

	if req.Role == ADMIN || req.Role == SUPERADMIN {
		return nil
	} else if req.Role == NORMAL {
		// ...
	}

	return nil
}

// 个人feed列表
func (s *FeedService) PersonalList(ctx context.Context, req *pb.PersonalListRequest, res *pb.ListResponse) error {
	// 获取feed数据
	list, err := model.GetPersonalFeedList(req.UserId, req.LastId, req.Size)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 格式化feed数据
	dataList, err := FormatListData(list)
	if err != nil {
		return e.ServerErr(errno.ErrFormatList, err.Error())
	}

	// 获取数据总数
	rows, err := model.GetPersonalRowsSum(req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 最大页数
	remain := rows % req.Size
	pageMax := rows / req.Size
	if remain != 0 {
		pageMax += 1
	}

	// 是否有下一页
	var hasNext = true
	if req.Page >= pageMax {
		hasNext = false
	}

	res.DataList = dataList
	res.HasNext = hasNext
	res.PageMax = pageMax
	res.PageNum = req.Page
	res.RowsNum = uint32(len(list))

	return nil
}

type DataMap struct {
	Lock *sync.Mutex
	Data map[uint32]*pb.SingleData
}

// 格式化，转为rpc相应的格式
func FormatListData(list []*model.FeedModel) ([]*pb.SingleData, error) {
	var dataMap = DataMap{
		Lock: new(sync.Mutex),
		Data: make(map[uint32]*pb.SingleData, len(list)),
	}

	wg := &sync.WaitGroup{}

	for _, l := range list {
		wg.Add(1)

		go func(feed model.FeedModel) {
			defer wg.Done()

			user := &pb.User{
				Name:      feed.Username,
				Id:        feed.UserId,
				AvatarUrl: feed.UserAvatar,
			}

			source := &pb.Source{
				KindId:      feed.SourceKindId,
				ObjectId:    feed.SourceObjectId,
				ObjectName:  feed.SourceObjectName,
				ProjectId:   feed.SourceProjectId,
				ProjectName: feed.SourceProjectName,
			}

			data := &pb.SingleData{
				Action:  feed.Action,
				FeedId:  feed.Id,
				IfSplit: false,
				TimeDay: feed.TimeDay,
				TimeHm:  feed.TimeHm,
				User:    user,
				Source:  source,
			}

			dataMap.Lock.Lock()
			defer dataMap.Lock.Unlock()
			dataMap.Data[feed.Id] = data

		}(*l)
	}

	wg.Wait()

	var result []*pb.SingleData
	var timeDay string
	var kindId uint32

	for index, data := range dataMap.Data {

		// ifSplit --> 分割线
		// 需要分割的情况
		// 1.第一条数据 2.不同日期 3.不同项目（source.KindId）
		if index == 0 {
			timeDay = data.TimeDay
			kindId = data.Source.KindId
			data.IfSplit = true
		} else if timeDay != data.TimeDay {
			timeDay = data.TimeDay
			data.IfSplit = true
		} else if kindId != data.Source.KindId {
			kindId = data.Source.KindId
			data.IfSplit = true
		}

		result = append(result, data)
	}

	return result, nil
}

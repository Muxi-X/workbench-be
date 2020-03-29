package service

import (
	"context"
	"sync"

	"muxi-workbench-feed/errno"
	"muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	e "muxi-workbench/pkg/err"
)

// 全部feed列表
func (s *FeedService) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	list, err := model.GetFeedList(req.LastId, req.Size) // page从1开始
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	dataList, err := FormatListData(list)
	if err != nil {
		return e.ServerErr(errno.ErrFormatList, err.Error())
	}

	rows, err := model.GetRowsSum()
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	remain := rows % req.Size
	pageMax := rows / req.Size
	if remain != 0 {
		pageMax += 1
	}

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

	return nil
}

// 个人feed列表
func (s *FeedService) PersonalList(ctx context.Context, req *pb.PersonalListRequest, res *pb.ListResponse) error {
	list, err := model.GetPersonalFeedList(req.UserId, req.LastId, req.Size)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	dataList, err := FormatListData(list)
	if err != nil {
		return e.ServerErr(errno.ErrFormatList, err.Error())
	}

	rows, err := model.GetPersonalRowsSum(req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	remain := rows % req.Size
	pageMax := rows / req.Size
	if remain != 0 {
		pageMax += 1
	}

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
	//errCh := new(chan error)
	//finish := new(chan bool)

	for _, l := range list {
		wg.Add(1)

		go func(feed model.FeedModel) {
			defer wg.Done()

			var ifSplit = false
			// TO DO:
			// how to split?

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
				IfSplit: ifSplit,
				TimeDay: feed.TimeDay,
				TimeHm:  feed.TimeHm,
				User:    user,
				Source:  source,
			}

			dataMap.Lock.Lock()
			dataMap.Data[feed.Id] = data
			dataMap.Lock.Unlock()

		}(*l)
	}

	wg.Wait()

	var result []*pb.SingleData
	for _, data := range dataMap.Data {
		result = append(result, data)
	}

	return result, nil
}

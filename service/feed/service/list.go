package service

import (
	"context"
	"strconv"

	"muxi-workbench-feed/errno"
	"muxi-workbench-feed/model"
	pb "muxi-workbench-feed/proto"
	e "muxi-workbench/pkg/err"
)

const PageSize = 40

// 全部feed列表
func (s *StatusService) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	page, err := strconv.ParseUint(req.Page, 10, 32)
	if err != nil {
		return e.ServerErr(errno.ErrBind, err.Error())
	}

	list, err := model.FeedList(uint32(page)*PageSize, PageSize)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	rows, err := model.GetRowsSum()
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	remain := rows % PageSize
	pageMax := rows / PageSize
	if remain != 0 {
		pageMax += 1
	}

	var hasNext = true
	if uint32(page) >= pageMax {
		hasNext = false
	}

	res = &pb.ListResponse{
		DataList: list,
		HasNext:  hasNext,
		PageMax:  pageMax,
		PageNum:  uint32(page),
		RowsNum:  uint32(len(list)),
	}

	// TO DO:
	// 为管理员，则返回所有数据
	// 为普通用户，则查询用户所在的project ids,从所有的datas中删去不在的数据，再返回

	return nil
}

// 个人feed列表
func (s *StatusService) PersonalList(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {

	return nil
}

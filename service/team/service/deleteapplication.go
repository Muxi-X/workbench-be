package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// DeleteApplication …… 删除申请请求
func (ts *TeamService) DeleteApplication(ctx context.Context, req *pb.DeleteApplicationRequest, res *pb.Response) error {
	if err := model.DeleteApply(req.ApplyList); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

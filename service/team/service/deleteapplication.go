package service

import (
	"context"
	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

//Delete …… 删除申请请求
func (ts *TeamService) DeleteApplication(ctx context.Context, req *pb.ApplicationRequest, res *pb.Response) error {
	if err := model.DeleteApply(req.UserId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

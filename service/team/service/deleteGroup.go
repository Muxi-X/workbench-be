package service

import (
	"context"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
	"github.com/Muxi-X/workbench-be/service/team/model"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
)

//Delete … 删除组别
func (ts *TeamService) Delete (ctx context.Context, req *pb.DeleteGroupRequest, res *pb.Response) error {
	err := model.DeleteGroup(req.GroupId)
	if err != nil{
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
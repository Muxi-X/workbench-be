package service

import (
	"context"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
	"github.com/Muxi-X/workbench-be/service/team/model"
	"time"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	errno "github.com/Muxi-X/workbench-be/service/team/errno"
)

//Create … 建立组别
func (ts *TeamService) Create(ctx context.Context, req *pb.CreateGroupRequest, res *pb.Response) error {
	t := time.Now()
	group := &model.GroupModel{
		Name:    req.GroupName,
		Order:   0,
		Counter: 0,
		Leader:  0,
		Time:    t.Format("2006-01-02 15:04:05"),
	}

	if err := group.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
}

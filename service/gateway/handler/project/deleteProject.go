package project

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// 需要 delete 和 feed push
func DeleteProject(c *gin.Context) {
	log.Info("Project delete function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 pid
	var pid int
	var err error

	pid, err = strconv.Atoi(c.Param("pid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取请求
	var req deleteRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 发送 delete 请求
	_, err2 := service.ProjectClient.DeleteProject(context.Background(), &pbp.GetRequest{
		Id: uint32(pid),
	})
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "删除",
		UserId: req.UserId,
		Source: &pbf.Source{
			Kind:        2,
			Id:          0, // 暂时从前端获取
			Name:        "",
			ProjectId:   uint32(pid),
			ProjectName: req.Projectname,
		},
	}

	// 发送 push 请求
	_, err3 := service.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		SendError(c, errno.InternalServerError, nil, err3.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}

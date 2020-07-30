package status

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
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
)

// 需要调用 feed push 和 status delete
func Delete(c *gin.Context) {
	log.Info("Status delete function call",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 sid
	var sid int
	var err error
	sid, err = strconv.Atoi(c.Param("sid"))
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

	// 调用 delete
	_, err2 := service.StatusClient.Delete(context.Background(), &pbs.GetRequest{
		Id: uint32(sid),
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
			Kind:        6,
			Id:          uint32(sid), // 暂时从前端获取
			Name:        req.Title,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err3 := service.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		SendError(c, errno.InternalServerError, nil, err3.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}
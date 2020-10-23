package status

import (
	"context"
	"strconv"

	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Delete ... 删除进度
func Delete(c *gin.Context) {
	log.Info("Status delete function call",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 sid
	var sid int
	var err error
	sid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求
	var req DeleteRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 调用 delete
	_, err = service.StatusClient.Delete(context.Background(), &pbs.GetRequest{
		Id: uint32(sid),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	id := c.MustGet("userID").(uint32)

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "删除",
		UserId: id,
		Source: &pbf.Source{
			Kind:        6,
			Id:          uint32(sid),
			Name:        req.Title,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err = service.FeedClient.Push(context.Background(), pushReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

package status

import (
	"context"
	"strconv"

	// pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// Update ... 编辑进度
// @Summary update status api
// @Description 编辑进度
// @Tags status
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "status_id"
// @Param object body UpdateRequest true "update_request"
// @Security ApiKeyAuth
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /status/detail/{id} [put]
func Update(c *gin.Context) {
	log.Info("Status update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 sid 和请求
	var sid int
	var err error

	sid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求体
	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	id := c.MustGet("userID").(uint32)

	// 构造 updateReq 并发送请求
	updateReq := &pbs.UpdateRequest{
		Id:      uint32(sid),
		Title:   req.Title,
		Content: req.Content,
		UserId:  id,
	}

	_, err = service.StatusClient.Update(context.Background(), updateReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	/*
		// 构造 push 请求
		pushReq := &pbf.PushRequest{
			Action: "编辑",
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
	*/

	SendResponse(c, errno.OK, nil)
}

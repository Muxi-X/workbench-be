package status

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Get ... 获取动态详情
// @Summary get status api
// @Description 获取进度实体
// @Tags status
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "status_id"
// @Success 200 {object} GetResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /status/detail/{id} [get]
func Get(c *gin.Context) {
	log.Info("Status get function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var sid int
	var err error

	sid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	uid := c.MustGet("userID").(uint32)

	// 构造 get status 请求并发送
	getReq := &pbs.GetRequest{
		Id:  uint32(sid),
		Uid: uid,
	}

	getResp, err := service.StatusClient.Get(context.Background(), getReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	resp := GetResponse{
		Sid:       uint32(sid),
		Title:     getResp.Status.Title,
		Content:   getResp.Status.Content,
		UserId:    getResp.Status.UserId,
		Time:      getResp.Status.Time,
		Liked:     getResp.Status.Liked,
		LikeCount: getResp.Status.Like,
	}

	SendResponse(c, nil, resp)
}

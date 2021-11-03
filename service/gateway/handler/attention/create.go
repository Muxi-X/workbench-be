package attention

import (
	"context"
	"strconv"

	pb "muxi-workbench-attention/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Create ... create new attention
// @Summary create attention api
// @Description 创建关注
// @Tags attention
// @Accept  application/json
// @Produce  application/json
// @Param id path int true "doc_id"
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /attention/{id} [post]
func Create(c *gin.Context) {
	log.Info("Attention create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 userId
	userId := c.MustGet("userID").(uint32)

	docId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	// 构造 create 请求
	createReq := &pb.PushRequest{
		DocId:  uint32(docId),
		UserId: userId,
	}

	// 向创建进度发起请求
	createResp, err := service.AttentionClient.Create(context.Background(), createReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, createResp)
}

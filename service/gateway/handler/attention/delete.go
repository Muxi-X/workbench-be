package attention

import (
	"context"
	pb "muxi-workbench-attention/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Delete ... delete attention
// @Summary delete attention api
// @Description 取消关注
// @Tags attention
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body FileRequest true "delete_attention_request"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /attention [delete]
func Delete(c *gin.Context) {
	log.Info("Attention create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获得请求
	var req FileRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userId
	userId := c.MustGet("userID").(uint32)

	// 构造 delete 请求
	deleteReq := &pb.PushRequest{
		FileId:   req.Id,
		UserId:   userId,
		FileKind: req.Kind,
	}

	// 向取消关注发起请求
	deleteResp, err := service.AttentionClient.Delete(context.Background(), deleteReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, deleteResp)
}

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

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// DeleteComment ... 删除进度的评论
// @Summary delete comment api
// @Description 通过 status_id 和 status_title 删除 status_comment
// @Tags status
// @Accept  application/json
// @Produce  application/json
// @Param id path int true "comment_id"
// @Param Authorization header string true "token 用户令牌"
// @Param object body DeleteCommentRequest  true "delete_comment_request"
// @Security ApiKeyAuth
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /status/comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	log.Info("Status delete function call",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 cid
	var cid int
	var err error
	cid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求体
	var req DeleteCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	id := c.MustGet("userID").(uint32)

	// 调用 Status 服务删除评论
	_, err = service.StatusClient.DeleteComment(context.Background(), &pbs.DeleteCommentRequest{
		UserId:    id,
		CommentId: uint32(cid),
		StatusId:  req.StatusId,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "取消评论",
		UserId: id,
		Source: &pbf.Source{
			Kind:        6,
			Id:          uint32(cid),
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

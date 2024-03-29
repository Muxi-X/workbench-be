package status

import (
	"context"
	pbf "muxi-workbench-feed/proto"
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

// CreateComment ... 评论 Status
// @Summary create comment api
// @Description 创建评论
// @Tags status
// @Accept  application/json
// @Produce  application/json
// @Param id path int true "status_id"
// @Param object body CreateCommentRequest true "create_comment_request"
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /status/comment/{id} [post]
func CreateComment(c *gin.Context) {
	log.Info("Status createComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 targetId
	targetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求体
	var req CreateCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	userId := c.MustGet("userID").(uint32)

	_, err = service.StatusClient.CreateComment(context.Background(), &pbs.CreateCommentRequest{
		UserId:   userId,
		TargetId: uint32(targetId),
		Content:  req.Content,
		Kind:     req.Kind,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 要通过 statusId 获取 status 的 title
	getReq := &pbs.GetRequest{
		Id: uint32(3), // TODO wrong id
	}

	// TODO: 需要获取创建的 comment id
	getResp, err := service.StatusClient.Get(context.Background(), getReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "评论",
		UserId: userId,
		Source: &pbf.Source{
			Kind:      6,
			Id:        uint32(targetId), // 暂时从前端获取
			Name:      getResp.Status.Title,
			ProjectId: 0,
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

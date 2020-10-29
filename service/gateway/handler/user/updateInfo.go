package user

import (
	"context"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pb "muxi-workbench-user/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateInfo ... 修改用户个人信息
// @Summary update info api
// @Description 修改用户个人信息
// @Tags user
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body UpdateInfoRequest  true "update_info_request"
// @Security ApiKeyAuth
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /user [put]
func UpdateInfo(c *gin.Context) {
	log.Info("User getInfo function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req UpdateInfoRequest
	if err := c.BindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userId := c.MustGet("userID").(uint32)

	// 构造请求给 getInfo
	updateInfoReq := &pb.UpdateInfoRequest{
		Id: userId,
		Info: &pb.UserInfo{
			Nick:      req.Nick,
			Name:      req.Name,
			AvatarUrl: req.AvatarURL,
			Email:     req.Email,
		},
	}

	// 发送请求
	_, err := service.UserClient.UpdateInfo(context.Background(), updateInfoReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

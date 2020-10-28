package user

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pb "muxi-workbench-user/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetProfile ... 获取 userProfile
// @Summary get user_profile api
// @Description 通过 userId 获取完整 user 信息
// @Tags user
// @Accept  application/json
// @Produce  application/json
// @Param id path int true "user_id"
// @Param Authorization header string true "token 用户令牌"
// @Param object body GetProfileRequest  true "get_profile_request"
// @Security ApiKeyAuth
// @Success 200 {object} UserProfile
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /user/profile/{id} [get]
func GetProfile(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var id int
	var err error

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	getProfileReq := &pb.GetRequest{Id: uint32(id)}

	// 发送请求
	getProfileResp, err := service.UserClient.GetProfile(context.Background(), getProfileReq)
	if err != nil {
		// TO DO: 判断错误是否是用户不存在
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	resp := UserProfile{
		Id:     getProfileResp.Id,
		Nick:   getProfileResp.Nick,
		Name:   getProfileResp.Name,
		Avatar: getProfileResp.Avatar,
		Email:  getProfileResp.Email,
		Tel:    getProfileResp.Tel,
		Role:   getProfileResp.Role,
		Team:   getProfileResp.Team,
		Group:  getProfileResp.Group,
	}

	SendResponse(c, nil, resp)
}

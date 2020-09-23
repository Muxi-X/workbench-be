package user

import (
	"context"
	// "strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	pb "muxi-workbench-user/proto"
	// "muxi-workbench-gateway/pkg/token"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
)

// 暂时不知道 router
// GetProfile 通过 userid 获取完整 user 信息
func GetProfile(c *gin.Context) {
	log.Info("User getInfo function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 Id (暂时
	var req getProfileRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造请求给 getProfile
	getProfileReq := &pb.GetRequest{
		Id: req.Id,
	}

	// 发送请求
	getProfileResp, err2 := service.UserClient.GetProfile(context.Background(), getProfileReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	// 构造返回 response
	resp := userProfile{
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

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

// GetProfile 通过 userid 获取完整 user 信息
func GetProfile(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	id := c.MustGet("userID").(uint32)

	getProfileReq := &pb.GetRequest{Id: id}

	// 发送请求
	getProfileResp, err := service.UserClient.GetProfile(context.Background(), getProfileReq)
	if err != nil {
		// TO DO: 判断错误是否是用户不存在
		SendError(c, errno.InternalServerError, nil, err.Error())
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

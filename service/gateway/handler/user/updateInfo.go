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
// userid 应该是通过 token 获取
// UpdateInfo 通过 userid 和 userinfo 上传信息
func UpdateInfo(c *gin.Context) {
	log.Info("User getInfo function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 Id (暂时
	var req updateInfoRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 getInfo
	updateInfoReq := &pb.UpdateInfoRequest{
		Id: req.Id,
		Info: &pb.UserInfo{
			Id:        req.Info.Id,
			Nick:      req.Info.Nick,
			Name:      req.Info.Name,
			AvatarUrl: req.Info.AvatarURL,
			Email:     req.Info.Email,
		},
	}

	// 发送请求
	_, err2 := service.UserClient.UpdateInfo(context.Background(), updateInfoReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

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

// 暂时不知道 router
// GetInfo 通过 userid 数组获取对应的 userInfoList
func GetInfo(c *gin.Context) {
	log.Info("User getInfo function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 Ids
	var req getInfoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if len(req.Ids) == 0 {
		SendResponse(c, nil, &getInfoResponse{})
		return
	}

	// 构造请求给 getInfo
	var getInfoReq = &pb.GetInfoRequest{}
	for _, id := range req.Ids {
		getInfoReq.Ids = append(getInfoReq.Ids, id)
	}

	// 发送请求
	getInfoResp, err := service.UserClient.GetInfo(context.Background(), getInfoReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	var resp getInfoResponse
	for _, item := range getInfoResp.List {
		resp.List = append(resp.List, userInfo{
			Id:        item.Id,
			Nick:      item.Nick,
			Name:      item.Name,
			AvatarURL: item.AvatarUrl,
			Email:     item.Email,
		})
	}

	SendResponse(c, nil, resp)
}
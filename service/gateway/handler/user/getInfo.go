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
// GetInfo 通过 userid 数组获取对应的 userInfoList
func GetInfo(c *gin.Context) {
	log.Info("User getInfo function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 Ids
	var req getInfoRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 getInfo
	var getInfoReq *pb.GetInfoRequest
	for i := 0; i < len(req.Ids); i++ {
		getInfoReq.Ids = append(getInfoReq.Ids, req.Ids[i])
	}

	// 发送请求
	getInfoResp, err2 := service.UserClient.GetInfo(context.Background(), getInfoReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error(), GetLine())
		return
	}

	// 构造返回 response
	var resp getInfoResponse
	for i := 0; i < len(getInfoResp.List); i++ {
		resp.List = append(resp.List, userInfo{
			Id:        getInfoResp.List[i].Id,
			Nick:      getInfoResp.List[i].Nick,
			Name:      getInfoResp.List[i].Name,
			AvatarURL: getInfoResp.List[i].AvatarUrl,
			Email:     getInfoResp.List[i].Email,
		})
	}

	SendResponse(c, nil, resp)
}

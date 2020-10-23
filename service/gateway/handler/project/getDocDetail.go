package project

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// 只调用一次 getdocdetail
// 不需要从 token 获取 userid
func GetDocDetail(c *gin.Context) {
	log.Info("project getDocDetail function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 did
	var did int
	var err error

	did, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	getDocDetailResp, err := service.ProjectClient.GetDocDetail(context.Background(), &pbp.GetRequest{
		Id: uint32(did),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	//构造返回结果
	resp := getDocDetailResponse{
		Id:           getDocDetailResp.Id,
		Title:        getDocDetailResp.Title,
		Content:      getDocDetailResp.Content,
		Creator:      getDocDetailResp.Creator,
		Createtime:   getDocDetailResp.CreateTime,
		Lasteditor:   getDocDetailResp.LastEditor,
		Lastedittime: getDocDetailResp.LastEditTime,
	}

	SendResponse(c, nil, resp)
}

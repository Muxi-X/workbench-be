package project

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetTrashbin ... 获取回收站文件
// type： 0-project 1-doc 2-file 3-doc folder 4-file folder
func GetTrashbin(c *gin.Context) {
	log.Info("project getTrashbin function call.", zap.String("X-Request-Id", util.GetReqID(c)))

	var limit, page int
	var err error

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	getTrashbinResp, err := service.ProjectClient.GetTrashbin(context.Background(), &pbp.GetTrashbinRequest{
		Offset: uint32(limit * page),
		Limit:  uint32(limit),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 解析结果
	var resp GetTrashbinResponse
	var count uint32
	for _, v := range getTrashbinResp.List {
		resp.List = append(resp.List, &Trashbin{
			Id:   v.Id,
			Type: v.Type,
			Name: v.Name,
		})
		count++
	}
	resp.Count = count

	SendResponse(c, nil, &resp)
}

package project

import (
	"context"

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
// 通过参数 type 获取不同类型的资源
// type： 0-project 1-doc 2-file 3-doc folder 4-file folder
func GetTrashbin(c *gin.Context) {
	log.Info("project getTrashbin function call.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	trashbinType := c.Param("type")

	if trashbinType > "5" || trashbinType < "0" {
		SendBadRequest(c, errno.ErrTrashbinType, nil, "get param type fail.", GetLine())
	}

	// 发送请求
	getTrashbinResp, err := service.ProjectClient.GetTrashbin(context.Background(), &pbp.GetTrashbinRequest{
		Type: trashbinType,
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

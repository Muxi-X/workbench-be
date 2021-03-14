package project

import (
	"context"
	"strconv"
	"strings"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetFileTree ... 获取某个文件夹下的文件树
func GetFileTree(c *gin.Context) {
	log.Info("project getFileTree function call.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 folderID
	folderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	getFileTreeResp, err := service.ProjectClient.GetFileChildren(context.Background(), &pbp.GetRequest{
		Id: uint32(folderID),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 解析结果
	var list []*FileChildrenItem
	raw := strings.Split(getFileTreeResp.Children, ",")
	for _, v := range raw {
		r := strings.Split(v, "-")
		if r[1] == "0" {
			list = append(list, &FileChildrenItem{
				Id:   r[0],
				Type: false,
			})
		} else {
			list = append(list, &FileChildrenItem{
				Id:   r[0],
				Type: true,
			})
		}
	}

	SendResponse(c, nil, &GetFileChildrenResponse{
		Count:        uint32(len(list)),
		FileChildren: list,
	})
}
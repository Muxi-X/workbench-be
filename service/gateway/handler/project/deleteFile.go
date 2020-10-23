package project

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// 调用 deletefile 和 feed push
// 需要从 token 获取 userid
func DeleteFile(c *gin.Context) {
	log.Info("File delete function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 fid
	var fid int
	var err error

	fid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取 req
	var req deleteFileRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	id := c.MustGet("userID").(uint32)

	// 请求
	_, err = service.ProjectClient.DeleteFile(context.Background(), &pbp.GetRequest{
		Id: uint32(fid),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造 push 请求
	// 待确认，file 的传法
	pushReq := &pbf.PushRequest{
		Action: "删除",
		UserId: id,
		Source: &pbf.Source{
			Kind:        4,
			Id:          uint32(fid), // 暂时从前端获取
			Name:        req.Filename,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err = service.FeedClient.Push(context.Background(), pushReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

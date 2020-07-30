package project

import (
	"context"

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

// 调用 createfile 和 feedpush
func CreateFile(c *gin.Context) {
	log.Info("File create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req createFileRequest
	if err := c.BindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造请求
	createFileReq := &pbp.CreateFileRequest{
		ProjectId: req.Pid,
		Name:      req.Filename,
		HashName:  req.Hashname,
		Url:       req.Url,
		UserId:    req.UserId,
	}
	_, err2 := service.ProjectClient.CreateFile(context.Background(), createFileReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "创建",
		UserId: req.UserId,
		Source: &pbf.Source{
			Kind:        4,
			Id:          req.Fid, // 暂时从前端获取
			Name:        req.Filename,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err3 := service.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		SendError(c, errno.InternalServerError, nil, err3.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}
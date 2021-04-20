package project

import (
	"context"

	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateFile creates a new file
// 调用 createfile 和 feed push
// 更新：同步 gateway 和 project
func CreateFile(c *gin.Context) {
	log.Info("project createFile function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req CreateFileRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	userID := c.MustGet("userID").(uint32)
	teamID := c.MustGet("teamID").(uint32)

	// 构造请求
	createFileReq := &pbp.CreateFileRequest{
		ProjectId:             req.ProjectID,
		Name:                  req.FileName,
		HashName:              req.HashName,
		Url:                   req.Url,
		UserId:                userID,
		TeamId:                teamID,
		FatherId:              req.FatherId,
		ChildrenPositionIndex: req.ChildrenPositionIndex,
	}
	resp, err := service.ProjectClient.CreateFile(context.Background(), createFileReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	/* --- 新增 feed --- */

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "创建",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        4,
			Id:          resp.Id, // 应该从接口获取
			Name:        req.FileName,
			ProjectId:   req.ProjectID,
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

package project

import (
	"context"
	"strconv"

	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// UpdateMembers updates the members in the project
// 调用一次 update 和一次 feed push
func UpdateMembers(c *gin.Context) {
	log.Info("Project member update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求
	var req UpdateMemberRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userID
	userID := c.MustGet("userID").(uint32)

	// 构造请求
	// 这里 list 应该是 uint32 表示 uid
	updateMemReq := &pbp.UpdateMemberRequest{
		Id: uint32(projectID),
	}

	for i := 0; i < len(req.UserList); i++ {
		updateMemReq.List = append(updateMemReq.List, req.UserList[i])
	}

	// 发送请求
	_, err = service.ProjectClient.UpdateMembers(context.Background(), updateMemReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	/* --- 新增 feed 动态 --- */

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "编辑",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        2,
			Id:          uint32(projectID),
			Name:        req.ProjectName,
			ProjectId:   uint32(projectID),
			ProjectName: req.ProjectName,
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

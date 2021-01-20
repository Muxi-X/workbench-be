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

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateDoc updates a doc' info
func UpdateDoc(c *gin.Context) {
	log.Info("project updateDoc function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 docID
	docID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求
	var req UpdateDocRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	userID := c.MustGet("userID").(uint32)

	updateReq := &pbp.UpdateDocRequest{
		Id:      uint32(docID),
		Title:   req.Title,
		Content: req.Content,
	}

	resp, err := service.ProjectClient.UpdateDoc(context.Background(), updateReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	/* --- 新增 feed --- */

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "编辑",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        3,
			Id:          uint32(docID),
			Name:        req.Title,
			ProjectId:   resp.Id,
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

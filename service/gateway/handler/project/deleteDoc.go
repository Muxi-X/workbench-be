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

// DeleteDoc deletes a doc
func DeleteDoc(c *gin.Context) {
	log.Info("Doc delete function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 docID
	docID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求
	var req DeleteDocRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userID
	userID := c.MustGet("userID").(uint32)
	role := c.MustGet("role").(uint32)

	resp, err := service.ProjectClient.DeleteDoc(context.Background(), &pbp.DeleteRequest{
		Id:         uint32(docID),
		FatherId:   req.FatherId,
		FatherType: req.FatherType,
		UserId:     userID,
		Role:       role,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	/* --- 新增 feed --- */

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "删除",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        3,
			Id:          uint32(docID),
			Name:        req.DocName,
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

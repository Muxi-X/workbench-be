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

// DeleteFile deletes a file
// @Summary delete a file api
// @Description 删除文件
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body DeleteFileRequest true "delete_file_request"
// @Param id path int true "file_id"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/file/{id} [delete]
func DeleteFile(c *gin.Context) {
	log.Info("File delete function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 fileID
	fileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取 req
	var req DeleteFileRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userId
	userID := c.MustGet("userID").(uint32)
	role := c.MustGet("role").(uint32)

	// 请求
	_, err = service.ProjectClient.DeleteFile(context.Background(), &pbp.DeleteRequest{
		Id:     uint32(fileID),
		UserId: userID,
		Role:   role,
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
			Kind:        4,
			Id:          uint32(fileID), // 暂时从前端获取
			Name:        req.FileName,
			ProjectId:   req.ProjectId,
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

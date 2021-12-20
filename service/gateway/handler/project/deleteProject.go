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

// DeleteProject deletes a project
// 要求要超管权限
// 需要 delete project 和 feed push
// @Summary deletes a project api
// @Description 删除项目
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body DeleteFolderRequest true "delete_folder_request"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /project [delete]
func DeleteProject(c *gin.Context) {
	log.Info("project deleteProject function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	projectID := c.MustGet("projectID").(uint32)

	// 获取 userid
	userID := c.MustGet("userID").(uint32)

	// 发送 delete 请求
	_, err := service.ProjectClient.DeleteProject(context.Background(), &pbp.GetRequest{
		Id: projectID,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	/* --- 新增 feed 动态 --- */

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "删除",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        2,
			Id:          projectID,
			Name:        "",
			ProjectId:   uint32(0),
			ProjectName: "",
		},
	}

	// 发送 push 请求
	_, err = service.FeedClient.Push(context.Background(), pushReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

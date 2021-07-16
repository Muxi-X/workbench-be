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

// GetProjectIdsForUser gets project ids for user
// @Summary get project ids for user api
// @Description 获取用户的所有项目
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} GetProjectIdsForUserResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /project/ids [get]
func GetProjectIdsForUser(c *gin.Context) {
	log.Info("Project getProjectIdsForUser function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 user_id
	userID := c.MustGet("userID").(uint32)

	getUserIdsForUserResponse, err := service.ProjectClient.GetProjectIdsForUser(context.Background(), &pbp.GetRequest{
		Id: userID,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, &GetProjectIdsForUserResponse{
		Ids: getUserIdsForUserResponse.List,
	})
}

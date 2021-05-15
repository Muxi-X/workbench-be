package middleware

import (
	"context"
	"fmt"
	"strconv"

	"muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// ProjectMiddleware ... 检查用户是否有 projectid 权限
func ProjectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID, err := strconv.Atoi(c.Param("project_id"))
		if err != nil {
			handler.SendBadRequest(c, errno.ErrNoProjectId, nil, err.Error(), handler.GetLine())
			c.Abort()
			return
		}

		userID := c.MustGet("userID").(uint32)
		fmt.Println(userID)

		// 检查
		resp, err := service.ProjectClient.CheckProjectForUser(context.Background(), &pbp.CheckProjectRequest{
			UserId:    userID,
			ProjectId: uint32(projectID),
		})
		if err != nil {
			handler.SendError(c, errno.InternalServerError, nil, err.Error(), handler.GetLine())
			c.Abort()
			return
		}

		if resp.IfValid {
			c.Set("projectID", projectID)
		} else {
			handler.SendBadRequest(c, errno.ErrProjectPermissionDenied, nil, err.Error(), handler.GetLine())
			c.Abort()
		}

		c.Next()
	}
}
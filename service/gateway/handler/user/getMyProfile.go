package user

import (
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetMyProfile ... 获取 myProfile
// @Summary get my_profile api
// @Description 获取 my 完整 user 信息
// @Tags user
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} UserProfile
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /user/myprofile [get]
func GetMyProfile(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	userId := c.MustGet("userID").(uint32)

	user, err := GetUserProfile(userId)

	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, user)
}

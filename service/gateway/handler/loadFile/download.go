package loadFile

import (
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Download
// @Tags load_file
// @Summary 下载文件
// @Description 通过存储的文件url，返回可用于直接下载的url
// @Param Authorization header string true "token 用户令牌"
// @Param file formData file true "二进制文件"
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} UrlModel
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /load_file/download [post]
func Download(c *gin.Context) {
	log.Info("loadFile Download function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求体
	var req UrlModel
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	url := service.Download(req.Url)

	SendResponse(c, nil, UrlModel{Url: url})
}

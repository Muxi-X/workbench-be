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
// @Summary download file api
// @Description 通过存储的文件url，返回可用于直接下载的url
// @Param Authorization header string true "token 用户令牌"
// @Param url body UrlModel true "file_url"
// @Accept  application/json
// @Produce  application/json
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

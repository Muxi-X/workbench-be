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

type UrlModel struct {
	Url string `json:"url"`
} // @name UrlModel

// Upload
// @Tags loadFile
// @Summary upload file api
// @Description 上传文件 图片，返回url
// @Param file formData file true "二进制文件"
// @Accept multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} UrlModel
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /load_file/upload [post]
func Upload(c *gin.Context) {
	log.Info("loadFile Upload function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		SendError(c, errno.ErrGetFile, nil, err.Error(), GetLine())
		return
	}

	dataLen := header.Size
	id, _ := c.Get("userID")

	url, err := service.UploadFile(header.Filename, id.(uint32), file, dataLen)
	if err != nil {
		SendError(c, errno.ErrUploadFile, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, &UrlModel{
		Url: url,
	})
}

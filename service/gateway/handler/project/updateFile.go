package project

import (
	"muxi-workbench-gateway/handler"

	"github.com/gin-gonic/gin"
)

// UpdateFile ... 修改文件
func UpdateFile(c *gin.Context) {

	handler.SendResponse(c, nil, nil)
}

package project

import (
	. "muxi-workbench-gateway/handler"

	"github.com/gin-gonic/gin"
)

// CreateProject creates new project
func CreateProject(c *gin.Context) {

	SendResponse(c, nil, nil)
}

package middleware

import (
	"muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		var ctx *token.Context
		var err error
		if ctx, err = token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, err.Error())
			c.Abort()
			return
		}

		c.Set("context", ctx)

		c.Next()
	}
}

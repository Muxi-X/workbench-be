package router

import (
	"net/http"

	"muxi-workbench-gateway/handler/sd"
	"muxi-workbench-gateway/router/middleware"

    "muxi-workbench-gateway/handler"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// api for authentication functionalities
	// g.POST("/login", user.Login)

	// The user handlers, requiring authentication
	// u := g.Group("/v1/user")
	// u.Use(middleware.AuthMiddleware())
	// {
	// 	u.POST("", user.Create)
	// 	u.DELETE("/:id", user.Delete)
	// 	u.PUT("/:id", user.Update)
	// 	u.GET("", user.List)
	// 	u.GET("/:username", user.Get)
	// }

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

    //下面是我写的路由
    //feed
    feed:=g.Group("/feed")
    {
        feed.GET("/list",handler.FeedList)
        feed.POST("/push",handler.FeedPush)
    }

    /*
    //status
    status:=g.Group("/status")
    {
        status.GET("/source",handler.StatusGet)
        status.GET("/:page",handler.StatusList)
        status.POST("/:sid",handler.StatusCreate)
        status.POST("/:sid",handler.StatusUpdate)
        status.POST("/:sid",handler.StatusDelete)
        status.PUT("/:sid/like",handler.StatusLike)
        status.POST("/:sid/comment",handler.StatusCreateComment)
        status.GET("/comments/:page",handler.StatusListComment)
    }

    //project
    g.POST("/project/new")

    project:=g.Group("/project/:pid")
    {
        project.GET("/",)
        project.DELETE("/",)
        project.POST("/",)
    }

    proMem:=g.Group("/project/:pid/member")
    {
        project.GET("/",)
        project.PUT("/",)
    }

    proDoc:=g.Group("/project/:pid/doc/:fid")
    {
        project.POST("/comments",)
        project.GET("/comments",)
        project.DELETE("/comment/:cid",)
        project.GET("/comment/:cid",)
    }

    proFile:=g.Group("/project/:pid/file/:fid")
    {
        project.GET("/comments",)
        project.POST("/comments",)
        project.GET("/comment/:cid",)
        project.DELETE("/comment/:cid",)
    }

    folder:=g.Group("/folder")
    {
        folder.GET("/filetree/:pid",)
        folder.PUT("/filetree/:pid",)
        folder.GET("/doctree/:pid",)
        folder.PUT("/doctree/:pid",)
    }*/

	return g
}

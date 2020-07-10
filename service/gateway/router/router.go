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
	feed := g.Group("/feed")
	{
		feed.GET("/", handler.FeedList)
		feed.POST("/", handler.FeedPush)
	}

	//status
	status := g.Group("/status")
	{
		status.GET("/:sid", handler.StatusGet)
		status.GET("/", handler.StatusList)
		status.POST("/", handler.StatusCreate)
		status.PUT("/:sid", handler.StatusUpdate)
		status.DELETE("/:sid", handler.StatusDelete)
		status.PUT("/:sid/like", handler.StatusLike)
		status.POST("/:sid/comments", handler.StatusCreateComment)
		status.GET("/:sid/comments/:page", handler.StatusListComment)
	}

	//project
	project := g.Group("/project")
	{
		project.GET("/", handler.GetProjectList)
		project.GET("/:pid", handler.GetProjectInfo)
		project.PUT("/:pid", handler.UpdateProjectInfo)
		project.DELETE("/:pid", handler.DeleteProject)
        project.GET("/:pid/member", handler.GetMembers)
        project.PUT("/:pid/member", handler.UpdateMembers)
        project.GET("/:pid/profile/:uid", handler.GetProjectIdsForUser)
	}

    folder:=g.Group("/folder")
    {
        folder.GET("/filetree/:pid", handler.GetFileTree)
		folder.GET("/doctree/:pid", handler.GetDocTree)
		folder.PUT("/filetree/:pid", handler.UpdateFileTree)
		folder.PUT("/doctree/:pid", handler.UpdateDocTree)
		folder.POST("/file", handler.CreateFile)
		folder.DELETE("/file/:id", handler.DeleteFile)
		folder.GET("/file/:id", handler.GetFileDetail)
	    folder.GET("/file", handler.GetFileInfoList)
		folder.GET("/list/:page", handler.GetFileFolderInfoList)
    }

    file:=g.Group("/file")
    {
        file.POST("/doc", handler.CreateDoc)
		file.PUT("/doc/:id", handler.UpdateDoc)
		file.DELETE("/doc/:id", handler.DeleteDoc)
		file.GET("/doc/:id", handler.GetDocDetail)
		file.GET("/doc", handler.GetDocInfoList)
		file.GET("/list/:page", handler.GetDocFolderInfoList)
    }

	return g
}

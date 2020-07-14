package router

import (
	"net/http"

	"muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/handler/feed"
	"muxi-workbench-gateway/handler/sd"
	"muxi-workbench-gateway/router/middleware"

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

	// 下面是我写的路由
	// feed
	feed := g.Group("api/v1/feed")
	{
		feed.GET("/list", feed.List)
		feed.GET("/list/:uid", feed.ListUser)
		feed.GET("/list/:uid/:gid", feed.ListGroup)
	}

	// status
	status := g.Group("api/v1/status")
	{
		status.GET("/:sid", status.Get)
		status.POST("/", status.Create)
		status.PUT("/:sid", status.Update)
		status.DELETE("/:sid", status.Delete)
		status.GET("/", status.List) // 暂时这样写
		status.GET("/:sid/filter/:uid", status.ListUser)
		status.GET("/:sid/filter/:uid/:gid", status.ListGroup)
		status.PUT("/:sid/like", status.Like)
		status.POST("/:sid/comments", status.CreateComment)
		// status.DELETE("/:sid/comment", status.DeleteComment)
	}

	// project
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

	folder := g.Group("/folder")
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

	file := g.Group("/file")
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

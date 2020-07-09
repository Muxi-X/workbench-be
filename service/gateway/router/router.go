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


    //status
    status:=g.Group("/status")
    {
        status.GET("/get",handler.StatusGet)
        status.GET("/list",handler.StatusList)
        status.POST("/create",handler.StatusCreate)
        status.POST("/update",handler.StatusUpdate)
        status.POST("/delete",handler.StatusDelete)
        status.PUT("/like",handler.StatusLike)
        status.POST("/comment",handler.StatusCreateComment)
        status.GET("/comments",handler.StatusListComment)
    }

    //project
    project:=g.Group("/project"){
        project.GET("/list",handler.GetPeojectList)
        project.GET("/info",handler.GetProjectInfo)
        project.POST("/info/update",handler.UpdateProjectInfo)
        project.DELETE("/delete",handler.DeleteProject)
        project.GET("/file/tree",handler.GetFileTree)
        project.GET("/doc/tree",handler.GetDocTree)
        project.POST("/file/tree",handler.UpdateFileTree)
        project.post("/doc/tree",handler.UpdateDocTree)
        project.GET("/members",handler.GetMembers)
        project.POST("/members",handler.UpdateMembers)
        project.GET("/ids",handler.GetProjectIdsForUser)
        project.POST("/file",handler.CreateFile)
        project.DELETE("file",handler.DeleteFile)
        project.GET("/file",handler.GetFileDetail)
        project.GET("/file/list",handler.GetFileInfoList)
        project.GET("/file/folder"handler.GetFileFolderInfoList)
        project.POST("/doc/new",handler.CreateDoc)
        project.POST("/doc",handler.UpdateDoc)
        project.DELETE("/doc",handler.DeleteDoc)
        project.GET("/doc",handler.GetDocDetail)
        project.GET("/doc/list",handler.GetDocInfoList)
        project.GET("/doc/folder",handler.GetDocFolderInfoList)
    }

	return g
}

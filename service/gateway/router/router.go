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

        // 多了一个筛选 group 的 api
        status.GET("/:sid/filter/:uid/:gid", status.ListGroup)
		status.PUT("/:sid/like", status.Like)
		status.POST("/:sid/comments", status.CreateComment)

        // 少了一个删除评论 api
		// status.DELETE("/:sid/comment", status.DeleteComment)
	}

	// project
	project := g.Group("/project")
	{
        // 创建一个 project  缺少 api
		// project.POST("/",project.CreateProject)

        // 获取一个 project 的信息，简介，之类的
        project.GET("/:pid",project.GetProjectInfo)//

        // 删除一个 project
        project.DELETE("/:pid",project.DeleteProject)//

        // 修改 project 的信息，简介之类的
        project.PUT("/:pid",project.UpdateProjectInfo)//

        // 获取一个 project 的成员
        project.GET("/:pid/member",project.GetMember)//

        // 编辑一个 project 的成员
        project.PUT("/:pid/member",project.UpdateMember)// 请求参数string有问题

        // 获取 project 的 list ,  swagger 里面没有
        project.GET("/",project.GetProjectList)//

        // 有关 project file doc 的评论的 api 全部没有

        // 好像是获取一个 user 的全部 project 的 id , 可能是用于别的 api 里面
		project.GET("/:pid/profile/:uid", handler.GetProjectIdsForUser)
	}

	folder := g.Group("/folder")
	{
        // 获取文件树
		folder.GET("/filetree/:pid", project.GetFileTree)//

		// 获取文档树
        folder.GET("/doctree/:pid", project.GetDocTree)//

        // 编辑文件树
        folder.PUT("/filetree/:pid", project.UpdateFileTree)//

        // 编辑文档树
        folder.PUT("/doctree/:pid", project.UpdateDocTree)//

        // 待修改
        folder.POST("/file", handle.CreateFile)
		folder.DELETE("/file/:id", handler.DeleteFile)
		folder.GET("/file/:id", handler.GetFileDetail)
		folder.GET("/file", handler.GetFileInfoList)
		folder.GET("/list/:page", handler.GetFileFolderInfoList)
	}

	file := g.Group("/file")
	{
        // 没有创建/编辑/删除 file/doc 文件夹的 api
        file.POST("/file",project.CreateFile)//
        file.DELETE("/file/:id",project.DeleteFile)//
        file.PUT("/file/:id",project.UpdateFile)//没有
        file.POST("/doc",project.CreateDoc)//
        file.GET("/doc/:id",project.GetDocDetail)//
        file.DELETE("/doc/:id",project.DeleteDoc)//
        file.PUT("/doc/:id",project.UpdateDoc)//


        // 待修改
		file.POST("/doc", handler.CreateDoc)
		file.PUT("/doc/:id", handler.UpdateDoc)
		file.DELETE("/doc/:id", handler.DeleteDoc)
		file.GET("/doc/:id", handler.GetDocDetail)
		file.GET("/doc", handler.GetDocInfoList)
		file.GET("/list/:page", handler.GetDocFolderInfoList)
	}

	return g
}

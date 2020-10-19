package router

import (
	"net/http"

	"muxi-workbench-gateway/handler/feed"
	"muxi-workbench-gateway/handler/project"
	"muxi-workbench-gateway/handler/sd"
	"muxi-workbench-gateway/handler/status"
	"muxi-workbench-gateway/handler/user"
	"muxi-workbench-gateway/router/middleware"
	"muxi-workbench/pkg/constvar"

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

	// 权限要求，普通用户/管理员/超管
	normalRequired := middleware.AuthMiddleware(constvar.AuthLevelNormal)
	// adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)
	// superAdminRequired := middleware.AuthMiddleware(constvar.AuthLevelSuperAdmin)

	// auth 模块
	authRouter := g.Group("api/v1/auth")
	{
		authRouter.POST("/login", user.Login)
		authRouter.POST("/signup", user.Register)
	}

	// user 模块
	userRouter := g.Group("api/v1/user")
	userRouter.Use(normalRequired)
	{
		// userRouter.GET("/infos", user.GetInfo)
		userRouter.GET("/profile/:id", user.GetProfile)
		userRouter.GET("/list", user.List)
		userRouter.PUT("", user.UpdateInfo)
	}

	// feed
	feedRouter := g.Group("api/v1/feed")
	feedRouter.Use(normalRequired)
	{
		feedRouter.GET("/list", feed.List)
		feedRouter.GET("/list/user/:id", feed.ListUser)
		feedRouter.GET("/list/group/:id", feed.ListGroup)
	}

	// status
	statusRouter := g.Group("api/v1/status")
	statusRouter.Use(normalRequired)
	{
		statusRouter.GET("/detail/:id", status.Get)
		statusRouter.POST("", status.Create)
		statusRouter.PUT("/detail/:id", status.Update)
		statusRouter.DELETE("/detail/:id", status.Delete)
		statusRouter.GET("", status.List)
		statusRouter.PUT("/like/:id", status.Like)
		statusRouter.POST("/comment/:id", status.CreateComment)
		statusRouter.DELETE("/comment/:id", status.DeleteComment)
		statusRouter.GET("/detail/:id/comments", status.CommentList)
	}

	// project
	projectRouter := g.Group("api/v1/project")
	projectRouter.Use(normalRequired)
	{
		// 创建一个 project  缺少 api
		// projectRouter.POST("/",project.CreateProject)

		// 获取一个 project 的信息，简介，之类的
		projectRouter.GET("/:id", project.GetProjectInfo)

		// 删除一个 project
		projectRouter.DELETE("/:id", project.DeleteProject)

		// 修改 project 的信息，简介之类的
		projectRouter.PUT("/:id", project.UpdateProjectInfo)

		// 获取一个 project 的成员
		projectRouter.GET("/:id/member", project.GetMembers)

		// 编辑一个 project 的成员
		projectRouter.PUT("/:id/member", project.UpdateMembers) // 请求参数string有问题

		// 获取 project 的 list ,  swagger 里面没有
		projectRouter.GET("", project.GetProjectList)

		// 有关 project file doc 的评论的 api 全部没有

		// 好像是获取一个 user 的全部 project 的 id , 可能是用于别的 api 里面
		// projectRouter.GET("/:pid/profile/:id", project.GetProjectIdsForUser) // uid

		// 待修改
		/*
		   projectRouter.GET("/:pid/re")
		   projectRouter.PUT("/:pid/re")
		   projectRouter.DELETE("/:pid/re")
		*/

		// comment
		/*
		   projectRouter.POST("/:pid/doc/:id/comments") // fid
		   projectRouter.GET("/:pid/doc/:id/comments")
		   projectRouter.DELETE("/:pid/doc/:id/comment/:cid")
		   projectRouter.GET("/:pid/doc/:id/comment/:cid")
		   projectRouter.GET("/:pid/file/:id/comments")
		   projectRouter.POST("/:pid/file:id/comments")
		   projectRouter.GET("/:pid/file/:id/comment/:cid")
		   projectRouter.DELETE("/:pid/file/:id/comment/:cid")
		*/
	}

	folderRouter := g.Group("api/v1/folder")
	folderRouter.Use(normalRequired)
	{
		// 获取文件树
		folderRouter.GET("/filetree/:id", project.GetFileTree)

		// 获取文档树
		folderRouter.GET("/doctree/:id", project.GetDocTree)

		// 编辑文件树
		folderRouter.PUT("/filetree/:id", project.UpdateFileTree)

		// 编辑文档树
		folderRouter.PUT("/doctree/:id", project.UpdateDocTree)

		// 待修改
		/*
		   folderRouter.GET("/file/:id", handler.GetFileDetail)
		   folderRouter.GET("/file", handler.GetFileInfoList)
		   folderRouter.GET("/list", handler.GetFileFolderInfoList)
		   folderRouter.POST("/file")              // create file folder
		   folderRouter.PUT("/file/:id")           // change folder name
		   folderRouter.DELETE("/file/:id")        // delete file folder
		   folderRouter.POST("/doc")               // create doc folder
		   folderRouter.PUT("/doc/:id")            // change doc folder name
		   folderRouter.DELETE("/doc/:id")         // delete doc folder
		   folderRouter.POST("/file/:id/children") // file children
		   folderRouter.POST("/doc/:id/children")  // doc children
		*/
	}

	fileRouter := g.Group("api/v1/file")
	fileRouter.Use(normalRequired)
	{
		// 没有创建/编辑/删除 file/doc 文件夹的 api
		fileRouter.POST("/file", project.CreateFile)       //
		fileRouter.DELETE("/file/:id", project.DeleteFile) //
		// fileRouter.PUT("/file/:id", project.UpdateFile)    //没有
		fileRouter.POST("/doc", project.CreateDoc)       //
		fileRouter.GET("/doc/:id", project.GetDocDetail) //
		fileRouter.DELETE("/doc/:id", project.DeleteDoc) //
		fileRouter.PUT("/doc/:id", project.UpdateDoc)    //

		// 待修改
		/*
		   fileRouter.POST("/doc", handler.CreateDoc)
		   fileRouter.PUT("/doc/:id", handler.UpdateDoc)
		   fileRouter.DELETE("/doc/:id", handler.DeleteDoc)
		   fileRouter.GET("/doc/:id", handler.GetDocDetail)
		   fileRouter.GET("/doc", handler.GetDocInfoList)
		   fileRouter.GET("/list", handler.GetDocFolderInfoList)
		*/
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}

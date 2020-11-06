package router

import (
	"net/http"

	"muxi-workbench-gateway/handler/feed"
	"muxi-workbench-gateway/handler/project"
	"muxi-workbench-gateway/handler/sd"
	"muxi-workbench-gateway/handler/status"
	"muxi-workbench-gateway/handler/team"
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
	adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)
	superAdminRequired := middleware.AuthMiddleware(constvar.AuthLevelSuperAdmin)

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
	{
		projectRouter.POST("", adminRequired, project.CreateProject)
		projectRouter.GET("", normalRequired, project.GetProjectList)           // 获取 project 的 list
		projectRouter.GET("/:id", normalRequired, project.GetProjectInfo)       // 获取一个 project 的信息
		projectRouter.DELETE("/:id", superAdminRequired, project.DeleteProject) // 删除一个 project
		projectRouter.PUT("/:id", adminRequired, project.UpdateProjectInfo)     // 修改 project 的信息
		projectRouter.GET("/:id/member", normalRequired, project.GetMembers)    // 获取一个 project 的成员
		projectRouter.PUT("/:id/member", adminRequired, project.UpdateMembers)  // 编辑一个 project 的成员

		// 有关 project file doc 的评论的 api 全部没有
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
	// folderRouter.Use(normalRequired)
	{
		folderRouter.GET("/filetree/:id", project.GetFileTree)    // 获取文件树
		folderRouter.GET("/doctree/:id", project.GetDocTree)      // 获取文档树
		folderRouter.PUT("/filetree/:id", project.UpdateFileTree) // 编辑文件树
		folderRouter.PUT("/doctree/:id", project.UpdateDocTree)   // 编辑文档树

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

	// 文件&文档
	fileRouter := g.Group("api/v1/file")
	fileRouter.Use(normalRequired)
	{
		// 文件
		fileRouter.POST("file", project.CreateFile)
		fileRouter.DELETE("/file/:id", project.DeleteFile)
		fileRouter.PUT("/file/:id", project.UpdateFile)
		// 文档
		fileRouter.POST("/doc", project.CreateDoc)
		fileRouter.GET("/doc/:id", project.GetDocDetail)
		fileRouter.DELETE("/doc/:id", project.DeleteDoc)
		fileRouter.PUT("/doc/:id", project.UpdateDoc)
	}

	teamRouter := g.Group("api/v1/team")
	{
		// team
		teamRouter.POST("/member", adminRequired, team.Join)
		teamRouter.DELETE("/member", adminRequired, team.Remove)
		teamRouter.POST("", superAdminRequired, team.CreateTeam)
		teamRouter.PUT("", superAdminRequired, team.UpdateTeamInfo)

		// invitation
		teamRouter.GET("/invitation", normalRequired, team.CreateInvitation)
		teamRouter.GET("/invitation/:hash", normalRequired, team.ParseInvitation)

		// group
		teamRouter.GET("/group/list", normalRequired, team.GetGroupList)
		teamRouter.GET("/group/members/list/:id", normalRequired, team.GetMemberList)
		teamRouter.PUT("/group/members", adminRequired, team.UpdateMembersForGroup)
		teamRouter.POST("/group", superAdminRequired, team.CreateGroup)
		teamRouter.DELETE("/group/:id", superAdminRequired, team.DeleteGroup)
		teamRouter.PUT("/group/:id", superAdminRequired, team.UpdateGroupInfo)

		// application
		teamRouter.POST("/application", normalRequired, team.CreateApplication)
		teamRouter.GET("/application/list", adminRequired, team.GetApplications)
		teamRouter.DELETE("/application", adminRequired, team.DeleteApplication)
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

package router

import (
	"muxi-workbench-gateway/handler/loadFile"
	"muxi-workbench-gateway/handler/sd"
	"net/http"

	_ "muxi-workbench-gateway/docs"
	"muxi-workbench-gateway/handler/attention"
	"muxi-workbench-gateway/handler/feed"
	"muxi-workbench-gateway/handler/project"
	"muxi-workbench-gateway/handler/status"
	"muxi-workbench-gateway/handler/team"
	"muxi-workbench-gateway/handler/user"
	"muxi-workbench-gateway/router/middleware"
	"muxi-workbench/pkg/constvar"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

	// swagger API doc
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 权限要求，普通用户/管理员/超管
	normalRequired := middleware.AuthMiddleware(constvar.AuthLevelNormal)
	adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)
	superAdminRequired := middleware.AuthMiddleware(constvar.AuthLevelSuperAdmin)

	// project 权限
	projectCheck := middleware.ProjectMiddleware()

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
		userRouter.GET("/myprofile/", user.GetMyProfile)
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
		projectRouter.GET("/list", normalRequired, project.GetProjectList)                // 获取 project 的 list
		projectRouter.GET("", normalRequired, projectCheck, project.GetProjectInfo)       // 获取一个 project 的信息
		projectRouter.DELETE("", superAdminRequired, projectCheck, project.DeleteProject) // 删除一个 project
		projectRouter.PUT("", adminRequired, projectCheck, project.UpdateProjectInfo)     // 修改 project 的信息
		projectRouter.GET("/member", normalRequired, projectCheck, project.GetMembers)    // 获取一个 project 的成员
		projectRouter.PUT("/member", adminRequired, projectCheck, project.UpdateMembers)  // 编辑一个 project 的成员
		projectRouter.GET("/ids", normalRequired, project.GetProjectIdsForUser)
		projectRouter.POST("/search", normalRequired, project.Search)
	}

	folderRouter := g.Group("api/v1/folder")
	folderRouter.Use(normalRequired)
	folderRouter.Use(projectCheck)
	{
		// 文档文件夹下的文件树
		folderRouter.GET("/file_children/:id", project.GetFileChildren) // 获取文件树
		folderRouter.GET("/doc_children/:id", project.GetDocChildren)   // 获取文档树
		// 移动文件
		folderRouter.PUT("/children/:old_father_id/:id", project.UpdateFilePosition)
		// 文档夹 crud
		folderRouter.GET("/docfolder", project.GetDocFolderInfoList)
		folderRouter.POST("/docfolder", project.CreateDocFolder)
		folderRouter.PUT("/docfolder/:id", project.UpdateDocFolder)
		folderRouter.DELETE("/docfolder/:id", project.DeleteDocFolder)

		// 文件夹 crud
		folderRouter.GET("/filefolder", project.GetFileFolderInfoList)
		folderRouter.POST("/filefolder", project.CreateFileFolder)
		folderRouter.PUT("/filefolder/:id", project.UpdateFileFolder)
		folderRouter.DELETE("/filefolder/:id", project.DeleteFileFolder)
	}

	// 文件&文档
	fileRouter := g.Group("api/v1/file")
	fileRouter.Use(normalRequired)
	fileRouter.Use(projectCheck)
	{
		// 文件
		fileRouter.POST("/file", project.CreateFile)
		fileRouter.DELETE("/file/:id", project.DeleteFile)
		fileRouter.PUT("/file/:id", project.UpdateFile)
		fileRouter.GET("/file/:id/children/:file_id", project.GetFileDetail)
		fileRouter.GET("/files", project.GetFileInfoList)
		// 文档
		fileRouter.POST("/doc", project.CreateDoc)
		fileRouter.GET("/doc/:id/children/:file_id", project.GetDocDetail)
		fileRouter.DELETE("/doc/:id", project.DeleteDoc)
		fileRouter.PUT("/doc/:id", project.UpdateDoc)
		fileRouter.POST("/doc/:id/comment", project.CreateDocComment)
		fileRouter.GET("/doc/:id/comments", project.ListDocComment)
		fileRouter.PUT("/doc/:id/comment/:comment_id", project.UpdateDocComment)
		fileRouter.DELETE("/doc/:id/comment/:comment_id", project.DeleteDocComment)
	}

	// 回收站 read delete recover
	trashbinRouter := g.Group("api/v1/trashbin")
	trashbinRouter.Use(normalRequired)
	trashbinRouter.Use(projectCheck)
	{
		trashbinRouter.GET("", project.GetTrashbin)
		trashbinRouter.PUT("/:id", project.UpdateTrashbin)
		trashbinRouter.DELETE("/:id", project.DeleteTrashbin)
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
		teamRouter.GET("/group", normalRequired, team.GetGroupList)
		teamRouter.GET("/group/members/:id", normalRequired, team.GetMemberList)
		teamRouter.PUT("/group/members", adminRequired, team.UpdateMembersForGroup)
		teamRouter.POST("/group", superAdminRequired, team.CreateGroup)
		teamRouter.DELETE("/group/detail/:id", superAdminRequired, team.DeleteGroup)
		teamRouter.PUT("/group/detail/:id", superAdminRequired, team.UpdateGroupInfo)

		// application
		teamRouter.POST("/application", normalRequired, team.CreateApplication)
		teamRouter.GET("/application/list", adminRequired, team.GetApplications)
		teamRouter.DELETE("/application", adminRequired, team.DeleteApplication)
	}

	attentionRouter := g.Group("api/v1/attention")
	attentionRouter.Use(normalRequired)
	{
		attentionRouter.GET("/list/:id", attention.List)
		attentionRouter.DELETE("", attention.Delete)
		attentionRouter.POST("", attention.Create)
	}

	uploadRouter := g.Group("api/v1/load_file")
	uploadRouter.Use(normalRequired)
	{
		uploadRouter.POST("/upload", loadFile.Upload)
		uploadRouter.POST("/download", loadFile.Download)
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

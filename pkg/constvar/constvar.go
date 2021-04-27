package constvar

const (
	DefaultLimit = 50

	// 角色权限
	Nobody     = 0 // 无权限用户
	Normal     = 1 // 普通用户
	Admin      = 3 // 管理员
	SuperAdmin = 7 // 超管

	// 权限限制等级
	AuthLevelNobody     = 0 // 无权限用户级别
	AuthLevelNormal     = 1 // 普通用户级别
	AuthLevelAdmin      = 2 // 管理员级别
	AuthLevelSuperAdmin = 4 // 超管级别

	// 项目编码 字符串
	ProjectCode    = uint8(0)
	DocCode        = uint8(1)
	FileCode       = uint8(2)
	DocFolderCode  = uint8(3)
	FileFolderCode = uint8(4)

	// 判断是否是文件夹的 code
	NotFolderCode = uint8(0)
	IsFolderCode  = uint8(1)

	// trashbin redis key
	Trashbin = "trashbin"
)

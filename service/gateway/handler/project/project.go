package project

type GetProjectInfoResponse struct {
	ProjectID   uint32 `json:"project_id"`
	ProjectName string `json:"project_name"`
	Intro       string `json:"intro"`
	UserCount   uint32 `json:"user_count"`
}

type UpdateRequest struct {
	ProjectName string `json:"project_name"`
	Intro       string `json:"intro"`
}

type ProjectListItem struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type MemberListItem struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	GroupName string `json:"group_name"`
	Role      uint32 `json:"role"`
}

type GetMemberResponse struct {
	Count uint32            `json:"count"`
	List  []*MemberListItem `json:"list"`
}

type UpdateMemberRequest struct {
	UserList []uint32 `json:"user_list"` // users' ids
}

type GetProjectListResponse struct {
	Count uint32             `json:"count"`
	List  []*ProjectListItem `json:"list"`
}

type GetFileTreeResponse struct {
	FileTree string `json:"file_tree"`
}

type UpdateFileTreeRequest struct {
	ProjectName string `json:"project_name"`
	FileTree    string `json:"file_tree"`
}

type GetDocTreeResponse struct {
	DocTree string `json:"doc_tree"`
}

type UpdateDocTreeRequest struct {
	ProjectName string `json:"project_name"`
	DocTree     string `json:"doc_tree"`
}

type CreateFileRequest struct {
	ProjectID uint32 `json:"project_id"`
	FileID    uint32 `json:"file_id"`
	FileName  string `json:"file_name"`
	HashName  string `json:"hash_name"`
	Url       string `json:"url"`
}

type DeleteFileRequest struct {
	FileName string `json:"file_name"`
}

type CreateDocRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	ProjectID uint32 `json:"project_id"`
	DocName   string `json:"doc_name"`
}

type GetDocDetailResponse struct {
	Id           uint32 `json:"doc_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Creator      string `json:"creator"`
	CreateTime   string `json:"create_time"`
	LastEditor   string `json:"last_editor"`
	LastEditTime string `json:"last_edit_time"`
}

type DeleteDocRequest struct {
	DocName string `json:"doc_name"`
}

type UpdateDocRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 下面是待抽离的 api
/*
func GetProjectIdsForUser(c *gin.Context) {
	var req pb.GetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetProjectIdsForUser(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetFileDetail(c *gin.Context) {
	var req pb.GetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetFileDetail(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetFileInfoList(c *gin.Context) {
	var req pb.GetInfoByIdsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetFileInfoList(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetFileFolderInfoList(c *gin.Context) {
	var req pb.GetInfoByIdsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetFileFolderInfoList(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetDocInfoList(c *gin.Context) {
	var req pb.GetInfoByIdsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetDocInfoList(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetDocFolderInfoList(c *gin.Context) {
	var req pb.GetInfoByIdsRequest
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetDocFolderInfoList(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}
*/

package project

type getProInfoResponse struct {
	Projectid   uint32 `json:"projectid"`
	Projectname string `json:"projectname"`
	Intro       string `json:"intro"`
	Usercount   uint32 `json:"usercount"`
}

type deleteRequest struct {
	UserId      uint32 `json:"userid"`
	Projectname string `json:"projectname"`
}

type updateRequest struct {
	UserId      uint32 `json:"userid"`
	Projectname string `json:projectname"`
	Intro       string `json:"intro"`
	Usercount   uint32 `json:"usercount"`
}

type projectListItem struct {
	Id   uint32 `json:"uid"`
	Name string `json:"username"`
	Logo string `json:logo"`
}

type memberListItem struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	GroupName string `json:"groupname"`
	Role      uint32 `json:"role"`
}

type getMemberResponse struct {
	Count      uint32           `json:"count"`
	Memberlist []memberListItem `json:"memberlist"`
}

// UserList 是 uint32
type updateMemberRequest struct {
	UserId      uint32   `json:"userid"`
	ProjectName string   `json:"projectname"`
	Userlist    []string `json:"userlist"`
}

type getProjectListRequest struct {
	UserId uint32 `json:"userid"`
}

type getProjectListResponse struct {
	Projectlist []projectListItem `json:"projectlist"`
}

type getFileTreeResponse struct {
	Filetree string `json:"filetree"`
}

type updateFileTreeRequest struct {
	UserId      uint32 `json:"userid"`
	Projectname string `json:"projectname"`
	Filetree    string `json:"filetree"`
}

type getDocTreeResponse struct {
	Doctree string `json:"doctree"`
}

type updateDocTreeRequest struct {
	UserId      uint32 `json:"userid"`
	Projectname string `json:"projectname"`
	Doctree     string `json:"doctree"`
}

type createFileRequest struct {
	UserId   uint32 `json:"userid"`
	Pid      uint32 `json:"projectid"`
	Filename string `json:"filename"`
	Hashname string `json:"hashname"`
	Url      string `json:"url"`
	Fid      uint32 `json:"fid"`
}

type deleteFileRequest struct {
	UserId   uint32 `json:"userid"`
	Filename string `json:"filename"`
}

type createDocRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Pid     uint32 `json:"projectid"`
	Docname string `json:"docname"`
	UserId  uint32 `json:"userid"`
}

type getDocDetailResponse struct {
	Id           uint32 `json:"docid"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Creator      string `json:"creator"`
	Createtime   string `json:"createtime"`
	Lasteditor   string `json:"lasteditor"`
	Lastedittime string `json:"lastedittime"`
}

type deleteDocRequest struct {
	UserId  uint32 `json:"userid"`
	Docname string `json:"docname"`
}

type updateDocRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  uint32 `json:"userId"`
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

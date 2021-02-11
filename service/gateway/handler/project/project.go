package project

type DeleteDocCommentRequest struct {
	ProjectId uint32 `json:"project_id"`
}

type EditTrashbinRequest struct {
	Type string `json:"type"`
}

type Trashbin struct {
	Id   uint32 `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type GetTrashbinResponse struct {
	Count uint32      `json:"count"`
	List  []*Trashbin `json:"list"`
}

type UpdateFileRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CreateProjectRequest struct {
	Name   string `json:"name"`
	Intro  string `json:"intro"`
	TeamId uint32 `json:"team_id"`
}

type UpdateFolderRequest struct {
	Name string `json:"name"`
}

type CreateFolderRequest struct {
	FatherId   uint32 `json:"father_id"`
	FatherType bool   `json:"father_type"`
	Name       string `json:"name"`
	ProjectId  uint32 `json:"project_id"`
}

type CreateDocCommentRequest struct {
	Content   string `json:"content"`
	ProjectId uint32 `json:"project_id"`
}

type UpdateDocCommentRequest struct {
	Content   string `json:"content"`
	ProjectId uint32 `json:"project_id"`
}

type Comment struct {
	Cid      uint32 `json:"cid"`
	Uid      uint32 `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Time     string `json:"time"`
	Content  string `json:"content"`
}

type CommentListResponse struct {
	Count       uint32    `json:"count"`
	CommentList []Comment `json:"commentlist"`
}

type GetProjectTreeResponse struct {
	DocTree  []*FileTreeItem `json:"doc_tree"`
	FileTree []*FileTreeItem `json:"file_tree"`
}

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

// FileTreeItem ... 文件树 包括和文件 文档
type FileTreeItem struct {
	Id   string `json:"id"`
	Type bool   `json:"type"` // 判断是不是 folder 0->file 1->folder
}

// GetFileTreeResponse ... 文件文档共用
type GetFileTreeResponse struct {
	Count    uint32          `json:"count"`
	FileTree []*FileTreeItem `json:"file_tree"`
}

// UpdateFileTreeRequest ... 文件文档共用
type UpdateFileTreeRequest struct {
	FileTree []*FileTreeItem `json:"file_tree"`
}

// UpdateProjectTreeRequest ... 用于更改项目的子树，相较于前者多了一个 bool 字段判断修改 docChildren 还是 FileChildren
type UpdateProjectTreeRequest struct {
	FileTree []*FileTreeItem
	Type     bool // 判断是 doc 还是 file
}

type CreateFileRequest struct {
	ProjectID uint32 `json:"project_id"`
	FileID    uint32 `json:"file_id"`
	FileName  string `json:"file_name"`
	HashName  string `json:"hash_name"`
	Url       string `json:"url"`
}

type GetFileDetailResponse struct {
	Id         uint32 `json:"file_id"`
	Url        string `json:"url"`
	Creator    string `json:"creator"`
	CreateTime string `json:"create_time"`
}

type DeleteFileRequest struct {
	FileName  string `json:"file_name"`
	ProjectId uint32 `json:"project_id"`
}

type CreateDocRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	ProjectID  uint32 `json:"project_id"`
	DocName    string `json:"doc_name"`
	FatherID   uint32 `json:"father_id"`   // 父节点 id
	FatherType bool   `json:"father_type"` // 0->project 1->folder
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

// GetFileInfoListRequest ... 获取文件信息请求，包括文件 文档 文件夹
type GetFileInfoListRequest struct {
	Ids []uint32 `json:"ids"`
}

type FileInfoItem struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

// GetFileInfoListResponse ... 获取文件信息，包括文件 文档 文件夹
type GetFileInfoListResponse struct {
	Count uint32          `json:"count"`
	List  []*FileInfoItem `json:"list"`
}

type GetProjectIdsForUserResponse struct {
	Ids []uint32 `json:"ids"`
}

package project

import "strings"

type UpdateFilePositionRequest struct {
	Id                    uint32
	FatherId              uint32
	FatherType            uint32
	Type                  uint8
	ChildrenPositionIndex uint32
} //@name UpdateFilePositionRequest

// DeleteFolderRequest ... 删除文件夹请求，文档文件共用
type DeleteFolderRequest struct {
	Id uint32 `json:"id"`
} //@name DeleteFolderRequest

// DeleteDocCommentRequest ... 删除文档评论请求
type DeleteDocCommentRequest struct {
	ProjectId uint32 `json:"project_id"`
} //@name DeleteDocCommentRequest

// DeleteTrashbinRequest ... 删除回收站请求
type DeleteTrashbinRequest struct {
	Type uint32 `json:"type"`
} //@name DeleteTrashbinRequest

// RemoveTrashbinRequest ... 恢复回收站文件请求
type RemoveTrashbinRequest struct {
	Type                  uint32 `json:"type"`
	FatherId              uint32 `json:"fatherId"`
	ChildrenPositionIndex uint32 `json:"children_position_index"`
	IsFatherProject       bool   `json:"is_father_project"`
} //@name RemoveTrashbinRequest

// Trashbin
type Trashbin struct {
	Id         uint32 `json:"id"`
	Type       uint32 `json:"type"`
	Name       string `json:"name"`
	DeleteTime string `json:"delete_time"`
	CreateTime string `json:"create_time"`
} //@name Trashbin

// GetTrashbinResponse ... 获取回收站资源响应
type GetTrashbinResponse struct {
	Count uint32      `json:"count"`
	List  []*Trashbin `json:"list"`
} //@name GetTrashbinResponse

// UpdateFileRequest ... 修改文件请求
type UpdateFileRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
} //@name UpdateFileRequest

type CreateProjectRequest struct {
	Name     string   `json:"name"`
	Intro    string   `json:"intro"`
	UserList []uint32 `json:"user_list"`
} //@name CreateProjectRequest

// UpdateFolderRequest ... 修改文档夹名字请求
type UpdateFolderRequest struct {
	Name string `json:"name"`
} //@name UpdateFolderRequest

// CreateFolderRequest ... 新建文件夹请求(文档夹和文件夹共用)
type CreateFolderRequest struct {
	FatherId              uint32 `json:"father_id"`
	Name                  string `json:"name"`
	ChildrenPositionIndex uint32 `json:"children_position_index"`
} //@name CreateFolderRequest

// CreateDocCommentRequest ... 创建文档评论请求
type CreateDocCommentRequest struct {
	Content string `json:"content"`
} //@name CreateDocCommentRequest

// UpdateDocCommentRequest ... 修改文档评论请求
type UpdateDocCommentRequest struct {
	Content   string `json:"content"`
	ProjectId uint32 `json:"project_id"`
} //@name UpdateDocCommentRequest

// DocComment
type DocComment struct {
	Cid      uint32 `json:"cid"`
	Uid      uint32 `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Time     string `json:"time"`
	Content  string `json:"content"`
} //@name DocComment

// DocCommentListResponse ... 获取文档评论响应
type DocCommentListResponse struct {
	Count       uint32       `json:"count"`
	CommentList []DocComment `json:"commentlist"`
} //@name DocCommentListResponse

// GetProjectInfoResponse ... 获取项目详情响应
type GetProjectInfoResponse struct {
	ProjectID    uint32              `json:"project_id"`
	ProjectName  string              `json:"project_name"`
	Intro        string              `json:"intro"`
	UserCount    uint32              `json:"user_count"`
	DocChildren  []*FileChildrenItem `json:"doc_children"`
	FileChildren []*FileChildrenItem `json:"file_children"`
	Time         string              `json:"time"`
	CreatorName  string              `json:"creator_name"`
} //@name GetProjectInfoResponse

// ProjectRequest ... 修改项目详情请求
type ProjectUpdateRequest struct {
	ProjectName string `json:"project_name"`
	Intro       string `json:"intro"`
} //@name ProjectUpdateRequest

// ProjectListItem
type ProjectListItem struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
} //@name ProjectListItem

// MemberListItem
type MemberListItem struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	GroupName string `json:"group_name"`
	Role      uint32 `json:"role"`
} //@name MemberListItem

// GetMemberResponse ... 获取项目成员响应
type GetMemberResponse struct {
	Count uint32            `json:"count"`
	List  []*MemberListItem `json:"list"`
} //@name GetMemberResponse

// UpdateMemberRequest ... 修改项目成员
type UpdateMemberRequest struct {
	UserList []uint32 `json:"user_list"` // users' ids
} //@name UpdateMemberRequest

// GetProjectListResponse ... 获取项目列表请求
type GetProjectListResponse struct {
	Count uint32             `json:"count"`
	List  []*ProjectListItem `json:"list"`
} //@name GetProjectListResponse

// FileChildrenItem ... 文件树 包括和文件 文档
type FileChildrenItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type bool   `json:"type"` // 判断是不是 folder 0->file 1->folder
} //@name FileChildrenItem

// ChildrenInfo ... 文件文档共用
type ChildrenInfo struct {
	Id          uint32 `json:"id"`
	Type        bool   `json:"type"`
	Name        string `json:"name"`
	CreatTime   string `json:"creat_time"`
	CreatorName string `json:"creator_name"`
	Path        string `json:"path"`
} //@name ChildrenInfo

// GetFileChildrenResponse ... 文件文档共用
type GetFileChildrenResponse struct {
	Count        uint32          `json:"count"`
	FileChildren []*ChildrenInfo `json:"file_children"`
} //@name GetFileChildrenResponse

// UpdateFileChildrenRequest ... 文件文档共用
type UpdateFileChildrenRequest struct {
	FileChildren []*FileChildrenItem `json:"file_children"`
}

// UpdateProjectChildrenRequest ... 用于更改项目的子树，相较于前者多了一个 bool 字段判断修改 docChildren 还是 FileChildren
type UpdateProjectChildrenRequest struct {
	FileChildren []*FileChildrenItem
	Type         bool // 判断是 doc 还是 file
}

// CreateFileRequest ... 新建文件请求
type CreateFileRequest struct {
	FileName              string `json:"file_name"`
	Url                   string `json:"url"`
	FatherId              uint32 `json:"father_id"`
	ChildrenPositionIndex uint32 `json:"children_position_index"`
} //@name CreateFileRequest

// GetFileDetailResponse ... 获取文件详情响应
type GetFileDetailResponse struct {
	Id         uint32 `json:"file_id"`
	Url        string `json:"url"`
	Creator    string `json:"creator"`
	CreateTime string `json:"create_time"`
} //@name GetFileDetailResponse

// DeleteFileRequest ... 删除文件请求
type DeleteFileRequest struct {
	FileName string `json:"file_name"`
} //@name DeleteFileRequest

// CreateDocRequest ... 创建文档请求
type CreateDocRequest struct {
	Title                 string `json:"title"`
	Content               string `json:"content"`
	DocName               string `json:"doc_name"`
	FatherID              uint32 `json:"father_id"`               // 父节点 id
	ChildrenPositionIndex uint32 `json:"children_position_index"` // 子节点的位置
} //@name CreateDocRequest

// GetDocDetailResponse ... 获取文档详情请求
type GetDocDetailResponse struct {
	Id           uint32 `json:"doc_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Creator      string `json:"creator"`
	CreateTime   string `json:"create_time"`
	LastEditor   string `json:"last_editor"`
	LastEditTime string `json:"last_edit_time"`
} //@name GetDocDetailResponse

// 可能 feed 有用
type DeleteDocRequest struct {
	DocName string `json:"doc_name"`
} //@name DeleteDocRequest

// UpdateDocRequest ... 修改文档请求
type UpdateDocRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
} //@name UpdateDocRequest

// FileInfoItem
type FileInfoItem struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
} //@name FileInfoItem

// GetFileInfoListResponse ... 获取文件信息，包括文件 文档 文件夹
type GetFileInfoListResponse struct {
	Count uint32          `json:"count"`
	List  []*FileInfoItem `json:"list"`
} //@name GetFileInfoListResponse

// GetProjectIdsForUserResponse ... 获取成员的所有项目响应
type GetProjectIdsForUserResponse struct {
	Ids []uint32 `json:"ids"`
} //@name GetProjectIdsForUserResponse

type SearchRequest struct {
	Type       uint32 `json:"type"`
	Keyword    string `json:"keyword"`
	UserId     uint32 `json:"user_id"`
	Offset     uint32 `json:"offset"`
	Limit      uint32 `json:"limit"`
	Pagination bool   `json:"pagination"`
} //@name SearchRequest

type SearchResult struct {
	Id          uint32 `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	UserName    string `json:"user_name"`
	ProjectName string `json:"project_name"`
	Time        string `json:"time"`
} //@name SearchResult

type SearchResponse struct {
	List  []*SearchResult `json:"list"`
	Count uint32          `json:"count"`
} // @name SearchResponse

// 转换 子文件 共用函数

func FormatChildren(strChildren string) []*FileChildrenItem {
	var list []*FileChildrenItem
	raw := strings.Split(strChildren, ",")
	for _, v := range raw {
		r := strings.Split(v, "-")
		if r[1] == "0" { // TODO
			list = append(list, &FileChildrenItem{
				Id:   r[0],
				Type: false,
			})
		} else {
			list = append(list, &FileChildrenItem{
				Id:   r[0],
				Type: true,
			})
		}
	}

	return list
}

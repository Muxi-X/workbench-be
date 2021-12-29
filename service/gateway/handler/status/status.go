package status

// LikeRequest 点赞进度请求
type LikeRequest struct {
	Liked bool `json:"liked"`
} // @name LikeRequest

// CreateRequest 创建进度请求
type CreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
} // @name CreateRequest

type Comment struct {
	Id       uint32 `json:"id"`
	Uid      uint32 `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Time     string `json:"time"`
	Content  string `json:"content"`
	Kind     uint32 `json:"kind"`
} // @name Comment

// CommentListResponse 评论列表响应
type CommentListResponse struct {
	Total       uint32    `json:"total"` // 评论总数
	Count       uint32    `json:"count"` // 一级评论数
	CommentList []Comment `json:"comment_list"`
} // @name CommentListResponse

// GetResponse 获得进度实体响应
type GetResponse struct {
	Sid       uint32 `json:"sid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserId    uint32 `json:"user_id"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	Time      string `json:"time"`
	Liked     bool   `json:"liked"`
	LikeCount uint32 `json:"like_count"`
} // @name GetResponse

// UpdateRequest 编辑进度请求
type UpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
} // @name UpdateRequest

// DeleteRequest 删除进度请求
type DeleteRequest struct {
	Title string `json:"title"`
} // @name DeleteRequest

type Status struct {
	Id           uint32 `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Time         string `json:"time"`
	CommentCount uint32 `json:"comment_count"`
	LikeCount    uint32 `json:"like_count"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
	Liked        bool   `json:"liked"`
} // @name Status

// StatusListResponse 获取进度列表响应
type StatusListResponse struct {
	Count  uint32   `json:"count"`
	Status []Status `json:"stauts"`
} // @name StatusListResponse

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	Content string `json:"content"`
	Kind    uint32 `json:"kind"`
} // @name CreateCommentRequest

// UpdateCommentRequest ... 修改评论请求
type UpdateCommentRequest struct {
	Content  string `json:"content"`
	Kind     uint32 `json:"kind"`
	StatusId uint32 `json:"status_id"`
} // @name UpdateCommentRequest

// DeleteCommentRequest 删除评论请求
type DeleteCommentRequest struct {
	Title string `json:"title"`
	Kind  uint32 `json:"kind"`
} // @name DeleteCommentRequest

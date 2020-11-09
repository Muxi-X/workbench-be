package status

// LikeRequest 点赞进度请求
type LikeRequest struct {
	Liked bool `json:"liked"`
}

// CreateRequest 创建进度请求
type CreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Comment struct {
	Cid      uint32 `json:"cid"`
	Uid      uint32 `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Time     string `json:"time"`
	Content  string `json:"content"`
}

// CommentListResponse 评论列表响应
type CommentListResponse struct {
	Count       uint32    `json:"count"`
	CommentList []Comment `json:"commentlist"`
}

// GetResponse 获得进度实体响应
type GetResponse struct {
	Sid          uint32 `json:"sid"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	UserId       uint32 `json:"userid"`
	Time         string `json:"time"`
	LikeCount    uint32 `json:"like_count"`
	CommentCount uint32 `json:"comment_count"`
	Liked        bool   `json:"liked"`
	UserName     string `json:"user_name"`
}

// UpdateRequest 编辑进度请求
type UpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// DeleteRequest 删除进度请求
type DeleteRequest struct {
	Title string `json:"title"`
}

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
}

// ListResponse 获取进度列表响应
type ListResponse struct {
	Count  uint32   `json:"count"`
	Status []Status `json:"stauts"`
}

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	Content string `json:"content"`
}

// DeleteCommentRequest 删除评论请求
type DeleteCommentRequest struct {
	Title    string `json:"title"`
	StatusId uint32 `json:"status_id"`
}

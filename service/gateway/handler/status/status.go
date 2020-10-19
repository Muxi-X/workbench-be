package status

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

type CommentListResponse struct {
	Count       uint32    `json:"count"`
	CommentList []Comment `json:"commentlist"`
}
type GetResponse struct {
	Sid      uint32 `json:"sid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserId   uint32 `json:userid"`
	Time     string `json:"time"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}

type UpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeleteRequest struct {
	Title string `json:title"`
}

type Status struct {
	Id       uint32 `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserId   uint32 `json:"userid"`
	Time     string `json:"time"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}

type ListResponse struct {
	Count  uint32   `json:"count"`
	Status []Status `json:"stauts"`
}

type CreateCommentRequest struct {
	Content string `json:"content"`
}

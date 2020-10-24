package status

type createRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	StatusId uint32 `json:"stauts_id"`
}

type comment struct {
	Cid      uint32 `json:"cid"`
	Uid      uint32 `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Time     string `json:"time"`
	Content  string `json:"content"`
}

type getResponse struct {
	Sid         uint32    `json:"sid"` // status id
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	UserId      uint32    `json:"user_id"`
	Time        string    `json:"time"`
	Avatar      string    `json:"avatar"`
	Username    string    `json:"username"`
	Count       uint32    `json:"count"`
	CommentList []comment `json:"comment_list"`
}

type updateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type deleteRequest struct {
	Title string `json:"title"`
}

type status struct {
	Id       uint32 `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserId   uint32 `json:"user_id"`
	Time     string `json:"time"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}

type listResponse struct {
	Count  uint32   `json:"count"`
	Status []status `json:"stauts"`
}

type createCommentRequest struct {
	Content string `json:"content"`
}

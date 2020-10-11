package feed

// TO DO：不需要，role 从 Authorization 获取 或 从 user-service 获取
type ListRequest struct {
	Role uint32 `json:"role"`
}

type User struct {
	Name      string `json:"name"`
	Id        uint32 `json:"id"`
	AvatarUrl string `json:"avatar_url"`
}

type Source struct {
	Kind        uint32 `json:"kind"`
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	ProjectId   uint32 `json:"projectid"`
	ProjectName string `json:"projectname"`
}

type FeedItem struct {
	Id          uint32  `json:"id"`
	Action      string  `json:"action"`
	ShowDivider bool    `json:"show_divider"` // 分割线
	Date        string  `json:"date"`
	Time        string  `json:"time"`
	User        *User   `json:"user"`
	Source      *Source `json:"source"`
}

type ListResponse struct {
	Count uint32      `json:"count"`
	List  []*FeedItem `json:"list"`
}

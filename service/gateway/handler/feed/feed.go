package feed

type listRequest struct {
	Role uint32 `json:"role"`
}

type user struct {
	Name      string `json:"name"`
	Id        uint32 `json:"id"`
	AvatarUrl string `json:avatar_url":`
}

type source struct {
	Kind        uint32 `json:"kind"`
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	ProjectId   uint32 `json:"projectid"`
	ProjectName string `json:"projectname"`
}

type feedItem struct {
	Id          uint32 `json:"id"`
	Action      string `json:"action"`
	ShowDivider bool   `json:"show_divider"`
	Date        string `json:"date"`
	Time        string `json:time"`
	User        user   `json:"user"`
	Source      source `json:"source"`
}

type listResponse struct {
	Count    uint32     `json:"count"`
	FeedItem []feedItem `json":feeditem"`
}

package feed

type FeedUser struct {
	Name      string `json:"name"`
	Id        uint32 `json:"id"`
	AvatarUrl string `json:"avatar_url"`
} //@name FeedUser

type Source struct {
	Kind        uint32 `json:"kind"` // 类型，1 -> 团队，2 -> 项目，3 -> 文档，4 -> 文件，6 -> 进度（5 不使用）
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	ProjectId   uint32 `json:"project_id"`
	ProjectName string `json:"project_name"`
} //@name Source

type FeedItem struct {
	Id          uint32    `json:"id"`
	Action      string    `json:"action"`
	ShowDivider bool      `json:"show_divider"` // 分割线
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	FeedUser    *FeedUser `json:"user"`
	Source      *Source   `json:"source"`
} //@name FeedItem

type FeedListResponse struct {
	Count uint32      `json:"count"`
	List  []*FeedItem `json:"list"`
} //@name FeedListResponse

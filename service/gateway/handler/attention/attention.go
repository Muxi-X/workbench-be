package attention

type AttentionUser struct {
	Name string `json:"name"`
	Id   uint32 `json:"id"`
} //@name AttentionUser

type AttentionItem struct {
	Id   uint32         `json:"id"`
	Date string         `json:"date"`
	Time string         `json:"time"`
	User *AttentionUser `json:"user"`
	Doc  *Doc           `json:"doc"`
} //@name AttentionItem

type AttentionListResponse struct {
	Count uint32           `json:"count"`
	List  []*AttentionItem `json:"list"`
} //@name AttentionListResponse

type Doc struct {
	Id          uint32         `json:"id"`
	Name        string         `json:"name"`
	DocCreator  *AttentionUser `json:"creator"`
	ProjectId   uint32         `json:"project_id"`
	ProjectName string         `json:"project_name"`
} //@name Doc

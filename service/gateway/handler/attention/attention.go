package attention

type AttentionUser struct {
	Name string `json:"name"`
	Id   uint32 `json:"id"`
} // @name AttentionUser

type AttentionItem struct {
	Id   uint32         `json:"id"`
	Date string         `json:"date"`
	Time string         `json:"time"`
	User *AttentionUser `json:"user"`
	File *File          `json:"file"`
} // @name AttentionItem

// FileRequest ... 添加删除关注请求
type FileRequest struct {
	Id   uint32 `json:"id"`
	Kind uint32 `json:"kind"`
} // @name FileRequest

type AttentionListResponse struct {
	Count uint32           `json:"count"`
	List  []*AttentionItem `json:"list"`
} // @name AttentionListResponse

type File struct {
	Id          uint32         `json:"file_id"`
	Name        string         `json:"file_name"`
	FileCreator *AttentionUser `json:"file_creator"`
	ProjectId   uint32         `json:"file_project_id"`
	ProjectName string         `json:"file_project_name"`
	Kind        uint32         `json:"file_kind"`
} // @name File

package user

// LoginRequest login 请求
type LoginRequest struct {
	OauthCode string `json:"oauth_code"`
} // @name LoginRequest

// LoginResponse login 请求响应
type LoginResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
} // @name LoginResponse

// RegisterRequest 注册请求
type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
} // @name RegisterRequest

// GetInfoRequest 获取 info 请求
type GetInfoRequest struct {
	Ids []uint32 `json:"ids" binding:"required"`
} // @name GetInfoRequest

type UserInfo struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	RealName  string `json:"real_name"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
}

// GetInfoResponse 获取 info 响应
type GetInfoResponse struct {
	List []UserInfo `json:"list"`
} // @name GetInfoResponse

// GetProfileRequest 获取 profile 请求
type GetProfileRequest struct {
	Id uint32 `json:"id"`
} // @name GetProfileRequest

// UserProfile 获取 profile 响应
type UserProfile struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	RealName string `json:"real_name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
	Role     uint32 `json:"role"`
	Team     uint32 `json:"team"`
	Group    uint32 `json:"group"`
} // @name UserProfile

// ListRequest 获取 userList 请求
type ListRequest struct {
	Team  uint32 `json:"team"`
	Group uint32 `json:"group"`
} // @name ListRequest

type User struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	RealName string `json:"real_name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Role     uint32 `json:"role"`
	Team     uint32 `json:"team"`
	Group    uint32 `json:"group"`
} // @name User

// ListResponse 获取 userList 响应
type ListResponse struct {
	Count uint32 `json:"count"`
	List  []User `json:"list"`
} // @name ListResponse

// UpdateInfoRequest 更新 userInfo 请求
type UpdateInfoRequest struct {
	UserInfo
} // @name UpdateInfoRequest

// UpdateTeamGroupRequest
type UpdateTeamGroupRequest struct {
	Ids   []uint32 `json:"ids"`
	Value uint32   `json:"value"`
	Kind  uint32   `json:"kind"`
} // @name UpdateTeamGroupRequest

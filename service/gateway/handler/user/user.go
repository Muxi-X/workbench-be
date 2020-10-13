package user

type LoginRequest struct {
	OauthCode string `json:"oauth_code"`
}

type LoginResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetInfoRequest struct {
	Ids []uint32 `json:"ids" binding:"required"`
}

type UserInfo struct {
	Id        uint32 `json:"id"`
	Nick      string `json:"nick"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

type GetInfoResponse struct {
	List []UserInfo `json:"list"`
}

type GetProfileRequest struct {
	Id uint32 `json:"id"`
}

type UserProfile struct {
	Id     uint32 `josn:"id"`
	Nick   string `json:"nick"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Tel    string `json:"tel"`
	Role   uint32 `json:"role"`
	Team   uint32 `json:"team"`
	Group  uint32 `json:"group"`
}

type ListRequest struct {
	Team  uint32 `json:"team"`
	Group uint32 `json:"group"`
}

type User struct {
	Id     uint32 `json:"id"`
	Nick   string `json:"nick"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Role   uint32 `json:"role"`
	Team   uint32 `json:"team"`
	Group  uint32 `json:"group"`
}

type ListResponse struct {
	Count uint32 `json:"count"`
	List  []User `json:"list"`
}

type UpdateInfoRequest struct {
	UserInfo
}

type UpdateTeamGroupRequest struct {
	Ids   []uint32 `json:"ids"`
	Value uint32   `json:"value"`
	Kind  uint32   `json:"kind"`
}

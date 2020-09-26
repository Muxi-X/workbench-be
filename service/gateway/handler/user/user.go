package user

type loginRequest struct {
	OauthCode string `json:"oauth_code"`
}

type loginResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

type registerRequest struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type getInfoRequest struct {
	Ids []uint32 `json:"ids" binding:"required"`
}

type userInfo struct {
	Id        uint32 `json:"id"`
	Nick      string `json:"nick"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

type getInfoResponse struct {
	List []userInfo `json:"list"`
}

type getProfileRequest struct {
	Id uint32 `json:"id"`
}

type userProfile struct {
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

type listRequest struct {
	Team  uint32 `json:"team"`
	Group uint32 `json:"group"`
}

type user struct {
	Id     uint32 `json:"id"`
	Nick   string `json:"nick"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Role   uint32 `json:"role"`
	Team   uint32 `json:"team"`
	Group  uint32 `json:"group"`
}

type listResponse struct {
	Count uint32 `json:"count"`
	List  []user `json:"list"`
}

type updateInfoRequest struct {
	userInfo
}

type updateTeamGroupRequest struct {
	Ids   []uint32 `json:"ids"`
	Value uint32   `json:"value"`
	Kind  uint32   `json:"kind"`
}

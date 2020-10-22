package team

// 权限分级
const (
	NOBODY     = 0
	MEMBER     = 1
	ADMIN      = 3
	SUPERADMIN = 7
)

type member struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	TeamID    uint32 `json:"teamid"`
	GroupID   uint32 `json:"groupid"`
	GroupName string `json:"groupname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type group struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	UserCount uint32 `json:"usercount"`
}

type applyUserItem struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type updateGroupInfoRequest struct {
	Role         uint32 `json:"role"`
	NewGroupName string `json:"newgroupname"`
	GroupID      uint32 `json:"groupid"`
}

type createGroupRequest struct {
	Role      uint32   `json:"role"`
	GroupName string   `json:"groupname"`
	UserIDs   []uint32 `json:"userids"`
}

type deleteGroupRequest struct {
	Role    uint32 `json:"role"`
	GroupID uint32 `json:"groupid"`
}

type updateMembersRequest struct {
	GroupID  uint32   `json:"groupid"`
	UserList []uint32 `json:"userlist"`
	Role     uint32   `json:"role"`
}

type joinRequest struct {
	UserList []uint32 `json:"userlist"`
	TeamID   uint32   `json:"teamid"`
}

type removeRequest struct {
	UserList []uint32 `json:"userlist"`
	TeamID   uint32   `json:"teamid"`
}

type createInvitationRequest struct {
	TeamID  uint32 `json:"teamid"`
	Expired int64  `json:"expired"`
}

type parseInvitationRequest struct {
	Hash string `json:"hash"`
}

type createTeamRequest struct {
	TeamName  string `json:"teamname"`
	CreatorID uint32 `json:"creatorid"`
	Role      uint32 `json:"role"`
}

type updateTeamInfoRequest struct {
	TeamID  uint32 `json:"teamid"`
	NewName string `json:"newname"`
	Role    uint32 `json:"role"`
}

type applicationRequest struct {
	UserID uint32 `json:"userid"`
}

type dropTeamRequest struct {
	TeamID uint32 `json:"teamid"`
	Role   uint32 `json:"role"`
}

type groupListResponse struct {
	Count  uint32  `json:"count"`
	Groups []group `json:"groups"`
}

type memberListResponse struct {
	Count   uint32   `json:"count"`
	Members []member `json:"members"`
}

type createInvitationResponse struct {
	Hash string `json:"hash"`
}

type parseInvitationResponse struct {
	TeamID uint32 `json:"teamid"`
}

type applicationListResponse struct {
	ApplyList []applyUserItem `json:"applys"`
	Count     uint32          `json:"count"`
}

package team

// 权限分级
const (
	NOBODY     = 0
	MEMBER     = 1
	ADMIN      = 3
	SUPERADMIN = 7
)

type createGroupRequest struct {
	Role      uint32   `json:"role"`
	GroupName string   `json:"groupname"`
	UserIDs   []uint32 `json:"userids"`
}

type deleteGroupRequest struct {
	Role    uint32 `json:"role"`
	GroupID uint32 `json:"groupid"`
}

type updateGroupInfo struct {
	Role         uint32 `json:"role"`
	NewGroupName string `json:"newgroupname"`
	GroupID      uint32 `json:"groupid"`
}

type group struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	UserCount uint32 `json:"usercount"`
}

type groupListResponse struct {
	Count  uint32  `json:"count"`
	Groups []group `json:"groups"`
}

type member struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	TeamID    uint32 `json:"teamid"`
	GroupID   uint32 `json:"groupid"`
	GroupName string `json:"groupname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type memberListResponse struct {
	Count   uint32   `json:"count"`
	Members []member `json:"members"`
}

type updateMembersRequest struct {
	GroupID  uint32   `json:"groupid"`
	UserList []uint32 `json:"userlist"`
	Role     uint32   `json:"role"`
}

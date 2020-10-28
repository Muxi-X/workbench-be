package team

type member struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	TeamID    uint32 `json:"team_id"`
	GroupID   uint32 `json:"group_id"`
	GroupName string `json:"group_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type group struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	UserCount uint32 `json:"user_count"`
}

type applyUserItem struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type updateGroupInfoRequest struct {
	NewGroupName string `json:"new_group_name"`
	GroupID      uint32 `json:"group_id"`
}

type createGroupRequest struct {
	GroupName string   `json:"group_name"`
	UserList  []uint32 `json:"user_list"`
}

type updateMembersRequest struct {
	GroupID  uint32   `json:"group_id"`
	UserList []uint32 `json:"user_list"`
}

type joinRequest struct {
	UserList []uint32 `json:"user_list"`
	TeamID   uint32   `json:"team_id"`
}

type removeRequest struct {
	UserList []uint32 `json:"user_list"`
	TeamID   uint32   `json:"team_id"`
}

type createInvitationRequest struct {
	TeamID  uint32 `json:"team_id"`
	Expired int64  `json:"expired"`
}

type createTeamRequest struct {
	TeamName  string `json:"team_name"`
	CreatorID uint32 `json:"creator_id"`
}

type updateTeamInfoRequest struct {
	TeamID  uint32 `json:"team_id"`
	NewName string `json:"new_name"`
}

type applicationRequest struct {
	UserID uint32 `json:"user_id"`
}

type dropTeamRequest struct {
	TeamID uint32 `json:"team_id"`
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
	TeamID uint32 `json:"team_id"`
}

type applicationListResponse struct {
	ApplyList []applyUserItem `json:"apply_list"`
	Count     uint32          `json:"count"`
}

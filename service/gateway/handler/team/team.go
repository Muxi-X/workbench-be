package team

type Member struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	TeamID    uint32 `json:"team_id"`
	GroupID   uint32 `json:"group_id"`
	GroupName string `json:"group_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type Group struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	UserCount uint32 `json:"user_count"`
}

type ApplyUserItem struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateGroupInfoRequest struct {
	NewGroupName string `json:"new_group_name"`
}

type CreateGroupRequest struct {
	GroupName string   `json:"group_name"`
	UserList  []uint32 `json:"user_list"`
}

type UpdateMembersRequest struct {
	GroupID  uint32   `json:"group_id"`
	UserList []uint32 `json:"user_list"`
}

type JoinRequest struct {
	UserList []uint32 `json:"user_list"`
}

type RemoveRequest struct {
	UserList []uint32 `json:"user_list"`
	TeamID   uint32   `json:"team_id"`
}

type CreateTeamRequest struct {
	TeamName string `json:"team_name"`
}

type UpdateTeamInfoRequest struct {
	NewName string `json:"new_name"`
}

type DropTeamRequest struct {
	TeamID uint32 `json:"team_id"`
}

type DeleteApplicationRequest struct {
	ApplicationList []uint32 `json:"application_list"`
}

type GroupListResponse struct {
	Count  uint32  `json:"count"`
	Groups []Group `json:"groups"`
}

type MemberListResponse struct {
	Count   uint32   `json:"count"`
	Members []Member `json:"members"`
}

type CreateInvitationResponse struct {
	Hash string `json:"hash"`
}

type ParseInvitationResponse struct {
	TeamID uint32 `json:"team_id"`
}

type ApplicationListResponse struct {
	ApplyList []ApplyUserItem `json:"apply_list"`
	Count     uint32          `json:"count"`
}

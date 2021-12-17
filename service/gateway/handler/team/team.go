package team

type Member struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	TeamID    uint32 `json:"team_id"`
	GroupID   uint32 `json:"group_id"`
	GroupName string `json:"group_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Role      uint32 `json:"role"`
} // @name Member

type Group struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	UserCount uint32 `json:"user_count"`
} // @name Group

type ApplyUserItem struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
} // @name ApplyUserItem

type UpdateGroupInfoRequest struct {
	NewGroupName string `json:"new_group_name"`
} // @name UpdateGroupInfoRequest

type CreateGroupRequest struct {
	GroupName string   `json:"group_name"`
	UserList  []uint32 `json:"user_list"`
} // @name CreateGroupRequest

type UpdateMembersRequest struct {
	GroupID  uint32   `json:"group_id"`
	UserList []uint32 `json:"user_list"`
} // @name UpdateMembersRequest

type JoinRequest struct {
	UserList []uint32 `json:"user_list"`
} // @name JoinRequest

type RemoveRequest struct {
	UserList []uint32 `json:"user_list"`
} // @name RemoveRequest

type CreateTeamRequest struct {
	TeamName string `json:"team_name"`
} // @name CreateTeamRequest

type UpdateTeamInfoRequest struct {
	NewName string `json:"new_name"`
} // @name UpdateTeamInfoRequest

type DropTeamRequest struct {
	TeamID uint32 `json:"team_id"`
} // @name DropTeamRequest

type DeleteApplicationRequest struct {
	ApplicationList []uint32 `json:"application_list"`
} // @name DeleteApplicationRequest

type GroupListResponse struct {
	Count  uint32  `json:"count"`
	Groups []Group `json:"groups"`
} // @name GroupListResponse

type MemberListResponse struct {
	Count   uint32   `json:"count"`
	Members []Member `json:"members"`
} // @name MemberListResponse

type CreateInvitationResponse struct {
	Hash string `json:"hash"`
} // @name CreateInvitationResponse

type ParseInvitationResponse struct {
	TeamID uint32 `json:"team_id"`
} // @name ParseInvitationResponse

type ApplicationListResponse struct {
	ApplyList []ApplyUserItem `json:"apply_list"`
	Count     uint32          `json:"count"`
} // @name ApplicationListResponse

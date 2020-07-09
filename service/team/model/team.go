package model

const (
	MUXI = 1 //muxi
)

func JoinTeam(teamid uint32, userid uint32) error {
	users := []uint32{userid}
	return UpdateUsersGroupidOrTeamid(users, teamid, TEAM)
}

func RemoveformTeam(teamid uint32, userid uint32) error {
	users := []uint32{userid}
	return UpdateUsersGroupidOrTeamid(users, 0, TEAM)
}

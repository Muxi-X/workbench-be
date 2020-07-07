package model

import "github.com/Muxi-X/workbench-be/service/team/service"

const(
	MUXI = 1 //muxi
)

func JoinTeam(teamid uint32, userid uint32) error {
	users := []uint32{userid}
	return service.UpdateUsersGroupidOrTeamid(users,teamid,TEAM)
}

func RemoveformTeam(teamid uint32, userid uint32) error {
	users := []uint32{userid}
	return service.UpdateUsersGroupidOrTeamid(users,0,TEAM)
}

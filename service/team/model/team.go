package model

import (
	m "muxi-workbench/model"
)

type TeamModel struct {
	ID      uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name    string `json:"name" gorm:"column:name;" binding:"required"`
	Count   uint32 `json:"count" gorm:"column:count;" binding:"required"`
	CreatorId  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	Time    string `json:"time" gorm:"column:time;" binding:"required"`
}

const (
	MUXI = 1 //muxi
)

func (t *TeamModel) TableName() string {
	return "teams"
}

func DropTeam(id uint32) error {
	team := &TeamModel{}
	team.ID = id
	return m.DB.Self.Delete(&team).Error
}

func (t *TeamModel) Create() error {
	return m.DB.Self.Create(&t).Error
}

func (t *TeamModel) Update() error {
	return m.DB.Self.Save(t).Error
}

//get team by teamid
func GetTeam(id uint32) (*TeamModel, error) {
	t := &TeamModel{}
	d := m.DB.Self.Where("id = ?", id).First(&t)
	return t, d.Error
}

func TeamCountAdd(teamid uint32) error {
	t, err := GetTeam(teamid)
	if err != nil {
		return err
	}

	t.Count++
	if err := t.Update(); err != nil {
		return err
	}

	return nil
}

func TeamCountSub(teamid uint32) error {
	t, err := GetTeam(teamid)
	if err != nil {
		return err
	}

	t.Count--
	if err := t.Update(); err != nil {
		return err
	}

	return nil
}

func JoinTeam(teamid uint32, userid uint32) error {
	users := []uint32{userid}
	if err := UpdateUsersGroupidOrTeamid(users, teamid, TEAM); err != nil {
		return err
	}
	if err := TeamCountAdd(teamid); err != nil {
		return err
	}
	return nil
}

func RemoveformTeam(teamid uint32, userid uint32) error {
	users := []uint32{userid}
	if err := UpdateUsersGroupidOrTeamid(users, 0, TEAM); err != nil {
		return err
	}
	if err := TeamCountSub(teamid); err != nil {
		return err
	}
	return nil
}


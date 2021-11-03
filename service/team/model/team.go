package model

import (
	"errors"

	m "muxi-workbench/model"

	"github.com/jinzhu/gorm"
)

type TeamModel struct {
	ID        uint32 `json:"id" gorm:"column:id;not null"`
	Name      string `json:"name" gorm:"column:name;"`
	Count     uint32 `json:"count" gorm:"column:count;"`
	CreatorID uint32 `json:"creator" gorm:"column:creator;"`
	Time      string `json:"time" gorm:"column:time;"`
}

const (
	MUXI = 1 // muxi
)

const (
	ADD = 1  // 加法操作
	SUB = -1 // 减法操作
)

func (t *TeamModel) TableName() string {
	return "teams"
}

// DropTeam drop a team by id
func DropTeam(id uint32) error {
	team := &TeamModel{}
	team.ID = id
	return m.DB.Self.Delete(&team).Error
}

// Create team
func (t *TeamModel) Create() error {
	return m.DB.Self.Create(&t).Error
}

// Update team
func (t *TeamModel) Update() error {
	return m.DB.Self.Save(t).Error
}

// GetTeam get team by id
func GetTeam(id uint32) (*TeamModel, error) {
	t := &TeamModel{}
	d := m.DB.Self.Where("id = ?", id).First(&t)
	if d.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	return t, d.Error
}

// TeamCountOperation choose a operation for count and value
func TeamCountOperation(teamID uint32, value uint32, operation int) error {
	t, err := GetTeam(teamID)
	if err != nil {
		return err
	}

	if operation == ADD {
		t.Count += value
	}
	if operation == SUB {
		if t.Count < value {
			return errors.New("减数大于被减数，超出了uint32的范围")
		}
		t.Count -= value
	}

	if err := t.Update(); err != nil {
		return err
	}

	return nil
}

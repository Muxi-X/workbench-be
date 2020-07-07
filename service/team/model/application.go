package model

import (
	m "github.com/Muxi-X/workbench-be/model"
	"github.com/Muxi-X/workbench-be/pkg/constvar"
)

type ApplyModel struct {
	ID     uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserID uint32 `json:"user_id" gorm:"column:user_id;" binding:"required"`
}

type ApplyUserItem struct {
	ID    uint32
	Name  string
	Eamil string
}

func (a *ApplyModel) TableName() string{
	return "applys"
}

func (a *ApplyModel) Create() error {
	return m.DB.Self.Create(&a).Error
}

func DeleteApply(userid uint32) error {
	apply := &ApplyModel{}
	apply.UserID = userid
	return m.DB.Self.Delete(&apply).Error
}

func ListApplictions(offset uint32, limit uint32, pagination bool) (, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	applicationlist := make([]*ApplyModel, 0)

	query := m.DB.Self.Table("applys").Select("id").Order("id desc")

	if pagination {
		query = query.Offset(offset).Limit(limit)
	}

	var count uint64

	if err := query.Scan(&applicationlist).Count(&count).Error; err != nil {

	}


}
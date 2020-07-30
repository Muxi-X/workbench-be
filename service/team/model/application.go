package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type ApplyModel struct {
	ID     uint32 `json:"id" gorm:"column:id;not null"`
	UserID uint32 `json:"user_id" gorm:"column:user_id;"`
}

type ApplyUserItem struct {
	ID    uint32
	Name  string
	Eamil string
}

func (a *ApplyModel) TableName() string {
	return "applys"
}

// Create apply
func (a *ApplyModel) Create() error {
	return m.DB.Self.Create(&a).Error
}

// DeleteApply delete an apply by id
func DeleteApply(userID uint32) error {
	return m.DB.Self.Where("user_id = ?", userID).Delete(&ApplyModel{}).Error
}

// ListApplys list all applys
func ListApplys(offset uint32, limit uint32, pagination bool) ([]*ApplyModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	applicationlist := make([]*ApplyModel, 0)

	query := m.DB.Self.Table("applys").Select("id, user_id")

	if pagination {
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Scan(&applicationlist).Error; err != nil {
		return nil, 0, nil
	}

	count := len(applicationlist)
	return applicationlist, uint64(count), nil
}

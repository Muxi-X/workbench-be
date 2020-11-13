package model

import (
	"errors"
	"github.com/jinzhu/gorm"

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

// Check whether there is a which has the specific id
func (a *ApplyModel) Check() error {
	var tmpApply ApplyModel

	d := m.DB.Self.Where("user_id = ?", a.UserID).First(&tmpApply)
	if d.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return errors.New("该用户已申请！请勿重复提交！")
}

// DeleteApply delete applications by id
func DeleteApply(applyList []uint32) error {
	return m.DB.Self.Where("id in (?)", applyList).Delete(&ApplyModel{}).Error

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

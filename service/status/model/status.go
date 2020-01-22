package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type StatusModel struct {
	ID      uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Content string `json:"content" gorm:"column:content;" binding:"required"`
	Title   string `json:"title" gorm:"column:title;" binding:"required"`
	Time    string `json:"time" gorm:"column:time;" binding:"required"`
	Like    uint32 `json:"like" gorm:"column:like;" binding:"required"`
	Comment uint32 `json:"comment" gorm:"column:comment;" binding:"required"`
	UserID  uint32 `json:"userId" gorm:"column:user_id;" binding:"required"`
}

func (c *StatusModel) TableName() string {
	return "status"
}

// Create status
func (u *StatusModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// Delete status
func DeleteStatus(id uint32) error {
	status := &StatusModel{}
	status.ID = id
	return m.DB.Self.Delete(&status).Error
}

// Update status
func (u *StatusModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetStatus get a single status by id
func GetStatus(id uint32) (*StatusModel, error) {
	s := &StatusModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

// ListStatus list all status
func ListStatus(groupId, offset, limit int, lastId uint32) ([]*StatusModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	statusList := make([]*StatusModel, 0)
	var count uint64

	if err := m.DB.Self.Model(&StatusModel{}).Count(&count).Error; err != nil {
		return statusList, count, err
	}

	if lastId != 0 {
		if err := m.DB.Self.Where("id < ?", lastId).Offset(offset).Limit(limit).Order("id desc").Find(&statusList).Error; err != nil {
			return statusList, count, err
		}
	} else {
		if err := m.DB.Self.Offset(offset).Limit(limit).Order("id desc").Find(&statusList).Error; err != nil {
			return statusList, count, err
		}
	}

	return statusList, count, nil
}

// // Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
// func (u *UserModel) Compare(pwd string) (err error) {
// 	err = auth.Compare(u.Password, pwd)
// 	return
// }

// // Encrypt the user password.
// func (u *UserModel) Encrypt() (err error) {
// 	u.Password, err = auth.Encrypt(u.Password)
// 	return
// }

// // Validate the fields.
// func (u *UserModel) Validate() error {
// 	validate := validator.New()
// 	return validate.Struct(u)
// }

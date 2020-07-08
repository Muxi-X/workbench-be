package model

import m "muxi-workbench/model"

type UserModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name     string `json:"name" gorm:"column:name;" binding:"required"`
	RealName string `json:"realName" gorm:"column:real_name;" binding:"required"`
	Email    string `json:"email" gorm:"column:email;" binding:"required"`
	Avatar   string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	Tel      string `json:"tel" gorm:"column:tel;"`
	Role     uint32 `json:"role" gorm:"column:role;" binding:"required"`
	TeamID   uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	GroupID  uint32 `json:"groupId" gorm:"column:group_id;" binding:"required"`
}

func (u *UserModel) TableName() string {
	return "users"
}

// Create ... create user
func (u *UserModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// // Delete status
// func DeleteStatus(id uint32) error {
// 	status := &StatusModel{}
// 	status.ID = id
// 	return m.DB.Self.Delete(&status).Error
// }

// // Update status
// func (u *StatusModel) Update() error {
// 	return m.DB.Self.Save(u).Error
// }

// GetUser get a single user by id
func GetUser(id uint32) (*UserModel, error) {
	s := &UserModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

// GetUserByIds get user by id array
func GetUserByIds(ids []uint32) ([]*UserModel, error) {
	list := make([]*UserModel, 0)
	if err := m.DB.Self.Where("id IN (?)", ids).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

// // ListStatus list all status
// func ListStatus(groupId, offset, limit int, lastId uint32, filter *StatusModel) ([]*StatusModel, uint64, error) {
// 	if limit == 0 {
// 		limit = constvar.DefaultLimit
// 	}

// 	statusList := make([]*StatusModel, 0)
// 	var count uint64

// 	if err := m.DB.Self.Model(&StatusModel{}).Where(filter).Count(&count).Error; err != nil {
// 		return statusList, count, err
// 	}

// 	if lastId != 0 {
// 		if err := m.DB.Self.Where(filter).Where("id < ?", lastId).Offset(offset).Limit(limit).Order("id desc").Find(&statusList).Error; err != nil {
// 			return statusList, count, err
// 		}
// 	} else {
// 		if err := m.DB.Self.Where(filter).Offset(offset).Limit(limit).Order("id desc").Find(&statusList).Error; err != nil {
// 			return statusList, count, err
// 		}
// 	}

// 	return statusList, count, nil
// }

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

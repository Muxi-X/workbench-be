package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"

	"github.com/jinzhu/gorm"
)

type UserModel struct {
	ID           uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string `json:"name" gorm:"column:name;" binding:"required"`
	RealName     string `json:"realName" gorm:"column:real_name;" binding:"required"`
	Email        string `json:"email" gorm:"column:email;" binding:"required"`
	Avatar       string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	Tel          string `json:"tel" gorm:"column:tel;"`
	Role         uint32 `json:"role" gorm:"column:role;" binding:"required"`
	TeamID       uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	GroupID      uint32 `json:"groupId" gorm:"column:group_id;" binding:"required"`
	EmailService uint32 `json:"emailService" gorm:"column:email_service;" binding:"required"`
	Message      uint32 `json:"message" gorm:"column:message;" binding:"required"`
}

func (u *UserModel) TableName() string {
	return "users"
}

// Create ... create user
func (u *UserModel) Create() error {
	return m.DB.Self.Create(u).Error
}

// Save save user.
func (u *UserModel) Save() error {
	return m.DB.Self.Save(u).Error
}

// GetUser get a single user by id
func GetUser(id uint32) (*UserModel, error) {
	user := &UserModel{}
	d := m.DB.Self.Where("id = ?", id).First(user)
	if d.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, d.Error
}

// GetUserByIds get user by id array
func GetUserByIds(ids []uint32) ([]*UserModel, error) {
	list := make([]*UserModel, 0)
	if err := m.DB.Self.Where("id IN (?)", ids).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

// GetUserByEmail get a user by email.
func GetUserByEmail(email string) (*UserModel, error) {
	u := &UserModel{}
	err := m.DB.Self.Where("email = ?", email).First(u).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return u, err
}

// GetUserByName get a user by name.
func GetUserByName(name string) (*UserModel, error) {
	u := &UserModel{}
	err := m.DB.Self.Where("name = ?", name).First(u).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return u, err
}

// ListUser list users
func ListUser(offset, limit, lastId uint32, filter *UserModel) ([]*UserModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	list := make([]*UserModel, 0)

	query := m.DB.Self.Model(&UserModel{}).Where(filter).Offset(offset).Limit(limit)

	if lastId != 0 {
		query = query.Where("id < ?", lastId).Order("id desc")
	}

	if err := query.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

// UpdateTeamAndGroup update user's team_id and group_id.
func UpdateTeamAndGroup(ids []uint32, value, kind uint32) (err error) {
	query := m.DB.Self.Table("users").Where("id In (?)", ids)
	if kind == 1 {
		err = query.Update("team_id", value).Error
	} else if kind == 2 {
		err = query.Update("group_id", value).Error
	}
	return
}

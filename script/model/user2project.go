package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// UserToProjectModel ... 用户和项目的中间表
type UserToProjectModel struct {
	ID     uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserID uint32 `json:"userId" gorm:"column:user_id;" binding:"required"`
}
type User struct {
	ID   uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Role uint32 `json:"role" gorm:"column:role;not null" binding:"required"`
}

func GetAllUser2project() []*UserToProjectModel {
	list := make([]*UserToProjectModel, 0)
	DB.Self.Table("user2projects").Scan(&list)
	return list
}

func ClearUser(val *UserToProjectModel) {
	var user User
	err := DB.Self.Table("users").Where("id = ?", val.UserID).First(&user).Error
	if err == gorm.ErrRecordNotFound || user.Role == 0 {
		fmt.Println(val.ID, " ", val.UserID, " ", user.Role)
		DB.Self.Table("user2projects").Where("id = ?", val.ID).Delete(val)
	}
}

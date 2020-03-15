package model

type UserToProjectModel struct {
	ID        uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserID    uint32 `json:"userId" gorm:"column:user_id;" binding:"required"`
	ProjectID uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

func (u *UserToProjectModel) TableName() string {
	return "user2projects"
}

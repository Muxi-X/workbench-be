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

type StatusListItem struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Title    string `json:"title" gorm:"column:title;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Like     uint32 `json:"like" gorm:"column:like;" binding:"required"`
	Comment  uint32 `json:"comment" gorm:"column:comment;" binding:"required"`
	UserID   uint32 `json:"userId" gorm:"column:user_id;" binding:"required"`
	Avatar   string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	UserName string `json:"username" gorm:"column:name;" binding:"required"`
	GroupID  uint32 `json:"groupId" gorm:"column:group_id;" binding:"required"`
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
func ListStatus(groupID, offset, limit, lastID uint32, filter *StatusModel) ([]*StatusListItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	statusList := make([]*StatusListItem, 0)

	query := m.DB.Self.Table("status").Select("status.*, users.name, users.avatar, users.group_id").Where(filter).Joins("left join users on users.id = status.user_id").Offset(offset).Limit(limit).Order("status.id desc")

	if lastID != 0 {
		query = query.Where("status.id < ?", lastID)
	}

	if groupID != 0 {
		query = query.Where("users.group_id = ?", groupID)
	}

	if filter.UserID != 0 {
		query = query.Where("status.user_id = ?", filter.UserID)
	}

	var count uint64

	if err := query.Scan(&statusList).Count(&count).Error; err != nil {
		return statusList, count, err
	}

	return statusList, count, nil
}

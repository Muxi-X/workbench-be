package model

import (
	"strconv"

	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"

	"github.com/jinzhu/gorm"
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
func DeleteStatus(db *gorm.DB, id, uid uint32) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 删除 status
	status := &StatusModel{}
	status.ID = id
	if err := m.DB.Self.Where("user_id=?", strconv.Itoa(int(uid))).Delete(&status).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除 comment
	if err := m.DB.Self.Where("statu_id=?", strconv.Itoa(int(id))).Delete(&CommentsModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除 user2status
	if err := m.DB.Self.Where("status_id=?", strconv.Itoa(int(id))).Delete(&UserToStatusModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
func ListStatus(groupID, teamID, offset, limit, lastID uint32, filter *StatusModel) ([]*StatusListItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	statusList := make([]*StatusListItem, 0)

	query := m.DB.Self.Table("status").Select("status.*, users.name, users.avatar, users.group_id").Where(filter).Joins("left join users on users.id = status.user_id").Offset(offset).Limit(limit).Order("status.id desc")

	if lastID != 0 {
		query = query.Where("status.id < ?", lastID)
	}

	if teamID != 0 {
		query = query.Where("users.team_id = ?", teamID)
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

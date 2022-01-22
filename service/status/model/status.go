package model

import (
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

// FilterParams provide filter's params.
type FilterParams struct {
	UserName string
	GroupId  uint32
	Key      string
}

func (s *StatusModel) TableName() string {
	return "status"
}

// Create status
func (s *StatusModel) Create() error {
	return m.DB.Self.Create(&s).Error
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
	if err := tx.Where("user_id = ?", uid).Delete(&status).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 软删除 一级comment
	if err := tx.Table("comments_status").Where("statu_id = ? AND kind = 0", id).Update("re", 1).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除 user2status
	if err := tx.Where("status_id = ?", id).Delete(&UserToStatusModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Update status
func (s *StatusModel) Update() error {
	return m.DB.Self.Save(s).Error
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

	if err := query.Scan(&statusList).Error; err != nil {
		return statusList, uint64(0), err
	}

	return statusList, uint64(len(statusList)), nil
}

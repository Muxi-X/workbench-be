package model

import (
	"strconv"

	m "muxi-workbench/model"

	"github.com/jinzhu/gorm"
)

type UserToStatusModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserID   uint32 `json:"userId" gorm:"column:user_id;" binding:"required"`
	StatusID uint32 `json:"statusId" gorm:"column:status_id;" binding:"required"`
}

func (c *UserToStatusModel) TableName() string {
	return "user2status"
}

func GetStatusLikeRecord(userID, statusID uint32) (*UserToStatusModel, error) {
	record := &UserToStatusModel{}
	d := m.DB.Self.Table("user2status").Where("user_id = ? AND status_id = ?", userID, statusID).First(&record)
	return record, d.Error
}

func GetStatusLikeRecordForUser(userID uint32) ([]*UserToStatusModel, error) {
	statusLikeList := make([]*UserToStatusModel, 0)
	d := m.DB.Self.Table("user2status").Where("user_id = ?", userID).Order("status_id desc").Scan(&statusLikeList)
	return statusLikeList, d.Error
}

func AddStatusLike(db *gorm.DB, u *UserToStatusModel, m *StatusModel) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(m).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func CancelStatusLike(db *gorm.DB, u *UserToStatusModel, m *StatusModel) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	record := &UserToStatusModel{}
	record.ID = u.ID
	if err := tx.Where("user_id = ? AND status_id = ?", strconv.Itoa(int(u.StatusID)), strconv.Itoa(int(u.UserID))).Delete(&record).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(m).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

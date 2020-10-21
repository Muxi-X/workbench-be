package model

import (
	m "muxi-workbench/model"
	"strconv"
)

type StatusLikeModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserID   uint32 `json:"userId" gorm:"column:user_id;" binding:"required"`
	StatusID uint32 `json:"statusId" gorm:"column:status_id;" binding:"required"`
}

func (c *StatusLikeModel) TableName() string {
	return "statuslike"
}

func (u *StatusLikeModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

func GetStatusLikeRecord(userID, statusID uint32) (*StatusLikeModel, error) {
	record := &StatusLikeModel{}
	d := m.DB.Self.Table("statuslike").Where("user_id = ? AND status_id = ?", userID, statusID).First(&record)
	return record, d.Error
}

func DeleteStatusLikeRecord(userID, statusID, ID uint32) error {
	record := &StatusLikeModel{}
	record.ID = ID
	return m.DB.Self.Where("user_id = ? AND status_id = ?", strconv.Itoa(int(statusID)), strconv.Itoa(int(userID))).Delete(&record).Error
}

func GetStatusLikeRecordForUser(userID uint32) ([]*StatusLikeModel, error) {
	statusLikeList := make([]*StatusLikeModel, 0)
	d := m.DB.Self.Table("status").Where("user_id = ?", userID).Find(&statusLikeList)
	return statusLikeList, d.Error
}

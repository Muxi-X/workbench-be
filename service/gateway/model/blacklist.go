package model

import (
	"time"

	m "muxi-workbench/model"
)

type BlacklistModel struct {
	ID        uint32    `json:"id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt int64     `json:"expires_at"`
}

func (b *BlacklistModel) TableName() string {
	return "blacklist"
}

// Create creates a new record into blacklist
func (b *BlacklistModel) Create() error {
	return m.DB.Self.Create(b).Error
}

// DeleteExpiredBlacklist deletes all expired blacklist records if expired,
// judge expired records by expires_at <= now(timestamp, int64)
func DeleteExpiredBlacklist() error {
	now := time.Now().Unix()
	return m.DB.Self.Where("expires_at <= ?", now).Delete(&BlacklistModel{}).Error
}

// HasExistedInBlacklist checks if the token exists in DB's blacklist
func HasExistedInBlacklist(token string) (bool, error) {
	var count int
	err := m.DB.Self.Table("blacklist").Where("token = ?", token).Count(&count).Error
	return count != 0, err
}

// GetAllBlacklist gets all blacklist records
func GetAllBlacklist() ([]*BlacklistModel, error) {
	var list []*BlacklistModel
	d := m.DB.Self.Find(&list)
	return list, d.Error
}

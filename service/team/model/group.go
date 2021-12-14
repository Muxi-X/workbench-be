package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"

	"github.com/jinzhu/gorm"
)

type GroupModel struct {
	ID     uint32 `json:"id" gorm:"column:id;not null"`
	Name   string `json:"name" gorm:"column:name;"`
	Order  uint32 `json:"order" gorm:"column:order; default:'null'"`
	Count  uint32 `json:"count" gorm:"column:count;"`
	Leader uint32 `json:"leader" gorm:"column:leader;"`
	Time   string `json:"time" gorm:"column:time;"`
}

const (
	NOBODY     = 0 // 无权限用户
	NORMAL     = 1 // 普通用户
	ADMIN      = 3 // 管理员
	SUPERADMIN = 7 // 超管
)

const (
	TEAM    = 1 // 对象:团队
	GROUP   = 2 // 对象:组别
	NOGROUP = 0
	NOTEAM  = 0
)

func (g *GroupModel) TableName() string {
	return "groups"
}

// Create group
func (g *GroupModel) Create() error {
	return m.DB.Self.Create(&g).Error
}

// DeleteGroup  delete group by id
func DeleteGroup(id uint32) error {
	group := GroupModel{ID: id}
	return m.DB.Self.Delete(&group).Error
}

// Update group
func (g *GroupModel) Update() error {
	return m.DB.Self.Save(g).Error
}

// GetGroup get a group by ID
func GetGroup(id uint32) (*GroupModel, error) {
	g := &GroupModel{}
	d := m.DB.Self.Where("id = ?", id).First(&g)
	if d.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	return g, d.Error
}

// ListGroup list all groups
func ListGroup(offset uint32, limit uint32, pagination bool) ([]*GroupModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	groupList := make([]*GroupModel, 0)

	query := m.DB.Self.Table("groups")

	if pagination {
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Scan(&groupList).Error; err != nil {
		return nil, 0, nil
	}

	count := len(groupList)
	return groupList, uint64(count), nil
}

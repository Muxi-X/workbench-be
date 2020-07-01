package model

import(
	"github.com/Muxi-X/workbench-be/pkg/constvar"
	m "github.com/Muxi-X/workbench-be/model"
	)

type MemberModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	TeamID     uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	GroupID    uint32 `json:"groupId" gorm:"column:group_id;" binding:"required"`
    GroupName  string
	Role       uint32 `json:"role" gorm:"column:role;" binding:"required"`
	Email      string `json:"email" gorm:"column:email;" binding:"required"`
	Avatar     string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	Name       string `json:"name" gorm:"column:name;" binding:"required"`
}

//list all members of a group
func ListMembersOfAGroup(groupID uint32, limit uint32, offset uint32,pagination bool) ([]*MemberModel, uint64,error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	memberlist := make([]*MemberModel, 0)

	query := m.DB.Self.Table("users").Select("group_id").Order("role desc")

	if pagination {
		query = query.Offset(offset).Limit(limit)
	}

	var count uint64

	if err := query.Scan(&memberlist).Count(&count).Error; err != nil {
		return memberlist, count, err
	}

	return memberlist, count, nil
}


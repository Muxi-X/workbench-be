package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"

	"github.com/jinzhu/gorm"
)

// UserToProjectModel ... 用户和项目的中间表
type UserToProjectModel struct {
	ID        uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserID    uint32 `json:"userId" gorm:"column:user_id;" binding:"required"`
	ProjectID uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

// TableName ... 物理表名
func (u *UserToProjectModel) TableName() string {
	return "user2projects"
}

// UserCount ... 项目用户数
type UserCount struct {
	Count uint32 `json:"count" gorm:"column:count(DISTINCT user_id);" binding:"required"`
}

// UserID ... 项目用户 id 列表
type UserID struct {
	ID uint32 `json:"userId" gorm:"column:user_id;" binding:"required"`
}

// GetProjectUserCount ... 获取项目人数
func GetProjectUserCount(id uint32) (uint32, error) {
	count := &UserCount{}
	d := m.DB.Self.Table("user2projects").Select("count(DISTINCT user_id)").Where("project_id = ?", id).First(&count)
	return count.Count, d.Error
}

// MemberListItem ... 项目成员列表项
type MemberListItem struct {
	ID      uint32 `json:"id" gorm:"column:user_id;not null" binding:"required"`
	Name    string `json:"name" gorm:"column:name;" binding:"required"`
	Avatar  string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	GroupID uint32 `json:"groupName" gorm:"column:group_id;" binding:"required"`
	Role    uint32 `json:"role" gorm:"column:role;" binding:"required"`
}

// GetUserToProjectByUser ... 根据用户 id 获取项目-成员关系
func GetUserToProjectByUser(id uint32) ([]*UserToProjectModel, error) {
	list := make([]*UserToProjectModel, 0)
	d := m.DB.Self.Table("user2projects").Where("user_id = ?", id).Find(&list)
	return list, d.Error
}

// GetUserListByProject ... 获取一个项目的 id 列表
func GetUserListByProject(id uint32) ([]*UserID, error) {
	list := make([]*UserID, 0)
	d := m.DB.Self.Table("user2projects").Where("project_id=?", id).Select("user_id").Order("user_id asc").Scan(&list)
	return list, d.Error
}

// UpdateMembers modify member list, add and delete by id list
func UpdateMembers(db *gorm.DB, projectID uint32, addList, delList []uint32) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 插入新的数据
	u := make([]*UserToProjectModel, 0)
	for i := 0; i < len(addList); i++ {
		u = append(u, &UserToProjectModel{
			ProjectID: projectID,
			UserID:    addList[i],
		})
	}

	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除没有的数据
	m := &UserToProjectModel{}
	if err := tx.Where("project_id=? AND user_id in", projectID, delList).Delete(m).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetProjectMemberList ... 项目成员列表
func GetProjectMemberList(projectID, offset, limit, lastID uint32, pagination bool) ([]*MemberListItem, uint64, error) {
	var count uint64

	list := make([]*MemberListItem, 0)

	query := m.DB.Self.Table("user2projects").Where("user2projects.project_id = ?", projectID).Select("user2projects.*, users.name, users.avatar, users.role, users.group_id").Joins("left join users on user2projects.user_id = users.id").Order("users.id")

	if pagination {
		if limit == 0 {
			limit = constvar.DefaultLimit
		}

		query = query.Offset(offset).Limit(limit)

		if lastID != 0 {
			query = query.Where("user.id < ?", lastID)
		}

		count = uint64(limit)
	}

	if err := query.Scan(&list).Error; err != nil {
		return list, count, err
	}

	if !pagination {
		count = uint64(len(list))
	}

	return list, count, nil
}

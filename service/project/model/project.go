package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	"time"

	g "github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// ProjectModel project table's structure
type ProjectModel struct {
	ID           uint32         `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string         `json:"name" gorm:"column:name;" binding:"required"`
	Intro        string         `json:"intro" gorm:"column:intro;" binding:"required"`
	Time         string         `json:"time" gorm:"column:time;" binding:"required"`
	Count        uint32         `json:"count" gorm:"column:count;" binding:"required"`
	TeamID       uint32         `json:"teamId" gorm:"column:team_id;" binding:"required"`
	FileChildren string         `json:"fileChildren" gorm:"column:file_children;" binding:"required"`
	DocChildren  string         `json:"docChildren" gorm:"column:doc_children;" binding:"required"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;" binding:"required"`
}

// ProjectListItem ProjectList service item
type ProjectListItem struct {
	ID   uint32 `json:"id" gorm:"column:project_id;not null" binding:"required"`
	Name string `json:"name" gorm:"column:name;" binding:"required"`
}

// ProjectName save the name of project
type ProjectName struct {
	Name string `json:"name" gorm:"column:name;" binding:"required"`
}

type ProjectChildren struct {
	DocChildren  string `json:"doc_children" gorm:"column:doc_children;" binding:"required"`
	FileChildren string `json:"file_children" gorm:"column:file_children;" binding:"required"`
}

// TableName return table name
func (u *ProjectModel) TableName() string {
	return "projects"
}

// Create ... 创建项目
func (u *ProjectModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// RecoverProject ... 恢复软删除
func RecoverProject(id uint32) error {
	return m.DB.Self.Unscoped().Table("projects").Where("id = ?", id).Update("deleted_at", "").Error
}

// DeleteProject ... 删除项目
// 事务
func DeleteProject(db *g.DB, id uint32) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	project := &ProjectModel{
		ID: id,
	}

	if err := db.Delete(project).Error; err != nil {
		tx.Rollback()
		return err
	}

	trashbin := &TrashbinModel{
		FileId:   id,
		FileType: constvar.ProjectCode,
		Name:     project.Name,
	}

	// 获取时间
	day := viper.GetInt("trashbin.expired")
	t := time.Now().Unix()
	trashbin.ExpiresAt = t + int64(time.Hour*24*time.Duration(day))

	// 插入回收站
	if err := trashbin.Create(); err != nil {
		tx.Rollback()
		return err
	}

	var res []string
	if err := GetProjectChildFolder(trashbin.FileId, &res); err != nil {
		tx.Rollback()
		return err
	}

	if len(res) != 0 {
		if err := m.SAddToRedis(constvar.Trashbin, res); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// GetProjectName return project's name
func GetProjectName(id uint32) (string, error) {
	record := &ProjectName{}
	err := m.DB.Self.Table("projects").Where("id=?", id).Select("name").First(record).Error
	return record.Name, err
}

// Update ... 更新项目
func (u *ProjectModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetProject ... 获取项目
func GetProject(id uint32) (*ProjectModel, error) {
	s := &ProjectModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

// ListProject list all project
func ListProject(userID, offset, limit, lastID uint32, pagination bool) ([]*ProjectListItem, uint64, error) {
	var count uint64

	projectList := make([]*ProjectListItem, 0)

	query := m.DB.Self.Table("user2projects").Where("user2projects.user_id = ?", userID).Select("user2projects.*, projects.name").Joins("left join projects on user2projects.project_id = projects.id").Order("projects.id")

	if pagination {
		if limit == 0 {
			limit = constvar.DefaultLimit
		}

		query = query.Offset(offset).Limit(limit).Count(&count)

		if lastID != 0 {
			query = query.Where("projects.id < ?", lastID)
		}
	}

	if err := query.Scan(&projectList).Error; err != nil {
		return projectList, count, err
	}

	return projectList, count, nil
}

func GetProjectChildrenById(id uint32) (*ProjectChildren, error) {
	s := &ProjectChildren{}
	d := m.DB.Self.Table("projects").Select("doc_children", "file_children").Where("id = ?", id).Find(&s)
	return s, d.Error
}

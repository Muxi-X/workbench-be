package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"

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
	CreatorId    uint32         `json:"creator_id" gorm:"column:creator_id;"`
}

type ProjectDetail struct {
	ID           uint32         `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string         `json:"name" gorm:"column:name;" binding:"required"`
	Intro        string         `json:"intro" gorm:"column:intro;" binding:"required"`
	Time         string         `json:"time" gorm:"column:time;" binding:"required"`
	Count        uint32         `json:"count" gorm:"column:count;" binding:"required"`
	TeamID       uint32         `json:"teamId" gorm:"column:team_id;" binding:"required"`
	FileChildren string         `json:"fileChildren" gorm:"column:file_children;" binding:"required"`
	DocChildren  string         `json:"docChildren" gorm:"column:doc_children;" binding:"required"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;" binding:"required"`
	CreatorId    uint32         `json:"creator_id" gorm:"column:creator_id;"`
	CreatorName  string         `json:"creator_name" gorm:"column:creator_name"`
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

// DeleteProject ... 删除项目
func DeleteProject(id uint32) error {
	project := &ProjectModel{
		ID: id,
	}

	return m.DB.Self.Delete(project).Error
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

// GetProjectDetail ... 获取项目详情
func GetProjectDetail(id uint32) (*ProjectDetail, error) {
	s := &ProjectDetail{}
	d := m.DB.Self.Table("projects").Where("projects.id = ?", id).Select("projects.*, users.name creator_name").Joins("left join users on users.id = creator_id").First(&s)
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

/* --- 移动文件 --- */

// UpdateFilePosition ... 移动文件，事务
func UpdateFilePosition(file interface{}, fatherId, oldFatherId uint32,
	fileType uint8, isFatherProject, isOldFatherProject bool, childrenPositionIndex uint32) error {
	tx := m.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 将目标文件 father_id 改变

	// 修改父文件的 children 插入子节点
	// 需要 isProject fatherId childrenPositionIndex obj(需要先断言)

	// 删除之前父文件的 children 子节点
	// 需要 isProject oldFatherId id isFolder
	switch fileType {
	case constvar.DocCode:
		doc := file.(DocModel)
		doc.FatherId = fatherId
		if err := tx.Update(doc).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := AddChildren(tx, isFatherProject, fatherId,
			childrenPositionIndex, doc); err != nil {
			tx.Rollback()
			return err
		}
		if err := DeleteChildren(tx, isOldFatherProject, oldFatherId, doc.ID,
			constvar.DocCode); err != nil {
			tx.Rollback()
			return err
		}
	case constvar.FileCode:
		file := file.(FileModel)
		file.FatherId = fatherId
		if err := file.Update(); err != nil {
			tx.Rollback()
			return err
		}
		if err := AddChildren(tx, isFatherProject, fatherId,
			childrenPositionIndex, file); err != nil {
			tx.Rollback()
			return err
		}
		if err := DeleteChildren(tx, isOldFatherProject, oldFatherId, file.ID,
			constvar.FileCode); err != nil {
			tx.Rollback()
			return err
		}
	case constvar.DocFolderCode:
		folder := file.(FolderForDocModel)
		folder.FatherId = fatherId
		if err := tx.Update(folder).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := AddChildren(tx, isFatherProject, fatherId,
			childrenPositionIndex, folder); err != nil {
			tx.Rollback()
			return err
		}
		if err := DeleteChildren(tx, isOldFatherProject, oldFatherId, folder.ID,
			constvar.DocCode); err != nil {
			tx.Rollback()
			return err
		}
	case constvar.FileFolderCode:
		folder := file.(FolderForFileModel)
		folder.FatherId = fatherId
		if err := tx.Update(folder).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := AddChildren(tx, isFatherProject, fatherId,
			childrenPositionIndex, folder); err != nil {
			tx.Rollback()
			return err
		}
		if err := DeleteChildren(tx, isOldFatherProject, oldFatherId, folder.ID,
			constvar.FileCode); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type AttentionModel struct {
	Id       uint32 `json:"id" gorm:"column:id"`
	UserId   uint32 `json:"user_id" gorm:"column:user_id"`
	FileId   uint32 `json:"file_id" gorm:"column:file_id"`
	TimeDay  string `json:"time_day" gorm:"column:time_day"`
	TimeHm   string `json:"time_hm" gorm:"column:time_hm"`
	FileKind uint32 `json:"file_kind" gorm:"column:file_kind"`
}

type AttentionDetail struct {
	Id       uint32 `json:"id"`
	UserId   uint32 `json:"user_id"`
	UserName string `json:"user_name"`
	File     `json:"file"`
	TimeDay  string `json:"time_day"`
	TimeHm   string `json:"time_hm"`
}

type File struct {
	Name        string `json:"file_name" gorm:"column:file_name"`
	Id          uint32 `json:"file_id" gorm:"column:file_id"`
	CreatorId   uint32 `json:"file_creator_id"`
	CreatorName string `json:"file_creator_name" gorm:"column:creator"`
	ProjectName string `json:"file_project_name"`
	ProjectId   uint32 `json:"file_project_id"`
	Kind        uint32 `json:"file_kind" gorm:"column:file_kind"`
}

func (*AttentionModel) TableName() string {
	return "user2attentions"
}

// Create a new attention
func (a *AttentionModel) Create() error {
	return m.DB.Self.Create(a).Error
}

// Delete a being attention
func (a *AttentionModel) Delete() error {
	if a.UserId == 0 { // 删除所有关注该id的attentions
		return m.DB.Self.Where("file_id = ? and file_kind = ?", a.FileId, a.FileKind).Delete(a).Error
	} else {
		return m.DB.Self.Where("user_id = ? and file_id = ? and file_kind = ?", a.UserId, a.FileId, a.FileKind).Delete(a).Error
	}
}

// GetByUserAndFile ...get attention info by user_id and file_id
func (a *AttentionModel) GetByUserAndFile() error {
	return m.DB.Self.Where("user_id = ? and file_id = ? and file_kind = ?", a.UserId, a.FileId, a.FileKind).First(a).Error
}

// FilterParams provide filter's params.
type FilterParams struct {
	UserId uint32
}

// List ... 查找attention
func List(lastId, limit uint32, filter *FilterParams) ([]*AttentionDetail, error) {
	var attentions []*AttentionDetail

	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	query := m.DB.Self.Table("user2attentions").Select("user2attentions.*, users.name user_name").Joins("left join users on users.id = user_id").Order("user2attentions.id desc").Limit(limit)

	// 查找用户的attention
	if filter.UserId != 0 {
		query = query.Where("user2attentions.user_id = ?", filter.UserId)
	}

	// 分页
	if lastId != 0 {
		query = query.Where("user2attentions.id < ?", lastId)
	}

	if err := query.Scan(&attentions).Error; err != nil {
		return nil, err
	}

	return attentions, nil
}

// GetDocDetail ... 获取文档详情
func GetDocDetail(id uint32) (*File, error) {
	s := &File{}
	d := m.DB.Self.Table("docs").Where("docs.id = ? AND re = 0", id).Select("docs.*, docs.id file_id, docs.filename file_name, users.name as creator, projects.name project_name").Joins("left join users on users.id = docs.creator_id").Joins("left join projects on project_id = projects.id").First(&s)
	return s, d.Error
}

// GetFileDetail ... 获取文件详情
func GetFileDetail(id uint32) (*File, error) {
	s := &File{}
	d := m.DB.Self.Table("files").Where("files.id = ? AND re = 0", id).Select("files.*, files.id file_id, files.filename file_name, users.name as creator, projects.name project_name").Joins("left join users on users.id = files.creator_id").Joins("left join projects on project_id = projects.id").First(&s)
	return s, d.Error
}

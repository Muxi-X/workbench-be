package model

import (
	m "muxi-workbench/model"
)

type AttentionModel struct {
	Id       uint32 `json:"id" gorm:"column:id"`
	UserId   uint32 `json:"user_id" gorm:"column:user_id"`
	Username string `json:"user_name"`
	Doc      `json:"doc"`
	TimeDay  string `json:"time_day" gorm:"column:time_day"`
	TimeHm   string `json:"time_hm" gorm:"column:time_hm"`
}

type Doc struct {
	Name        string `json:"doc_name"`
	Id          uint32 `json:"doc_id" gorm:"column:doc_id"`
	CreatorId   uint32 `json:"doc_creator_id"`
	CreatorName string `json:"doc_creator_name"`
	ProjectName string `json:"doc_project_name"`
	ProjectId   uint32 `json:"doc_project_id"`
}

func (*AttentionModel) TableName() string {
	return "attentions"
}

// Create a new attention
func (a *AttentionModel) Create() error {
	return m.DB.Self.Create(a).Error
}

// Delete a being attention
func (a *AttentionModel) Delete() error {
	return m.DB.Self.Delete(a).Error
}

// FilterParams provide filter's params.
type FilterParams struct {
	UserId uint32
}

// List ... 查找attention
func List(lastId, limit uint32, filter *FilterParams) ([]*AttentionModel, error) {
	var data []*AttentionModel

	query := m.DB.Self.Table("attentions").Select("attentions.*").Order("attentions.id desc").Limit(limit)

	// 查找用户的attention
	if filter.UserId != 0 {
		query = query.Where("attentions.userid = ?", filter.UserId)
	}

	// 分页
	if lastId != 0 {
		query = query.Where("attentions.id < ?", lastId)
	}

	if err := query.Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

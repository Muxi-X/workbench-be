package model

import (
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type AttentionModel struct {
	Id      uint32 `json:"id" gorm:"column:id"`
	UserId  uint32 `json:"user_id" gorm:"column:user_id"`
	DocId   uint32 `json:"doc_id" gorm:"column:doc_id"`
	TimeDay string `json:"time_day" gorm:"column:time_day"`
	TimeHm  string `json:"time_hm" gorm:"column:time_hm"`
}

type AttentionDetail struct {
	Id       uint32 `json:"id"`
	UserId   uint32 `json:"user_id"`
	Username string `json:"user_name"`
	Doc      `json:"doc"`
	TimeDay  string `json:"time_day"`
	TimeHm   string `json:"time_hm"`
}

type Doc struct {
	Name        string `json:"doc_name"`
	Id          uint32 `json:"doc_id"`
	CreatorId   uint32 `json:"doc_creator_id"`
	CreatorName string `json:"doc_creator_name"`
	ProjectName string `json:"doc_project_name"`
	ProjectId   uint32 `json:"doc_project_id"`
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
	return m.DB.Self.Where("user_id = ? and doc_id = ?", a.UserId, a.DocId).Delete(a).Error
}

// GetByUserAndDoc ...get attention info by user_id and doc_id
func (a *AttentionModel) GetByUserAndDoc() error {
	return m.DB.Self.Where("user_id = ? and doc_id = ?", a.UserId, a.DocId).First(a).Error
}

// FilterParams provide filter's params.
type FilterParams struct {
	UserId uint32
}

// List ... 查找attention
func List(lastId, limit uint32, filter *FilterParams) ([]*AttentionDetail, error) {
	var data []*AttentionModel

	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	query := m.DB.Self.Table("user2attentions").Select("user2attentions.*").Order("user2attentions.id desc").Limit(limit)

	// 查找用户的attention
	if filter.UserId != 0 {
		query = query.Where("user2attentions.user_id = ?", filter.UserId)
	}

	// 分页
	if lastId != 0 {
		query = query.Where("user2attentions.id < ?", lastId)
	}

	if err := query.Scan(&data).Error; err != nil {
		return nil, err
	}
	var attentions []*AttentionDetail
	for _, d := range data {
		attention := &AttentionDetail{
			Id:      d.Id,
			UserId:  d.UserId,
			TimeDay: d.TimeDay,
			TimeHm:  d.TimeHm,
			Doc: Doc{
				Id: d.DocId,
			},
		}
		attentions = append(attentions, attention)
	}
	return attentions, nil
}

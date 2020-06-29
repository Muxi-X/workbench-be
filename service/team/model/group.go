package model

import(
	m "github.com/Muxi-X/workbench-be/model"
)

type GroupModel struct {
	ID      uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name    string `json:"name" gorm:"column:name;" binding:"required"`
	Order   uint32 `json:"order" gorm:"column:order;" binding:"required"`
	Counter uint32 `json:"counter" gorm:"column:counter;" binding:"required"`
	Leader  uint32 `json:"leader" gorm:"column:leader;" binding:"required"`
	Time    string `json:"time" gorm:"column:time;" binding:"required"`
}

func (g *GroupModel) TableName() string{
	return "groups"
}

//create group
func (g *GroupModel) Create() error {
	return m.DB.Self.Create(&g).Error
}

//delete group
func DeleteGroup(id uint32) error {
	group := &GroupModel{}
	group.ID = id
	return m.DB.Self.Delete(&group).Error
}

//update group
func (g *GroupModel) Update() error {
	return m.DB.Self.Save(g).Error
}

func GetGroup(id uint32) (*GroupModel, error) {
	g := &GroupModel{}
	d := m.DB.Self.Where("id = ?", id).First(&g)
	return g, d.Error
}

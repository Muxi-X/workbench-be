package model

import(
	m "github.com/Muxi-X/workbench-be/model"
	"github.com/Muxi-X/workbench-be/pkg/constvar"
)

type GroupModel struct {
	ID      uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name    string `json:"name" gorm:"column:name;" binding:"required"`
	Order   uint32 `json:"order" gorm:"column:order;" binding:"required"`
	Counter uint32 `json:"counter" gorm:"column:counter;" binding:"required"`
	Leader  uint32 `json:"leader" gorm:"column:leader;" binding:"required"`
	Time    string `json:"time" gorm:"column:time;" binding:"required"`
}

const (
	NOBODY = 0     // 无权限用户
	NORMAL = 1     // 普通用户
	ADMIN = 3      // 管理员
	SUPERADMIN = 7 // 超管

)

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
//get group by groupid
func GetGroup(id uint32) (*GroupModel, error) {
	g := &GroupModel{}
	d := m.DB.Self.Where("id = ?", id).First(&g)
	return g, d.Error
}

//get groupid by groupname
func GetGroupId(name string) (uint32, error) {
	g := &GroupModel{}
	d := m.DB.Self.Where("name = ?", name).First(&g)
	return g.ID, d.Error
}

//list all of group
func ListGroup(offset uint32, limit uint32, pagination bool) ([]*GroupModel,uint64,error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	grouplist := make([]*GroupModel, 0)

	//order 根据groups.id降序输出，即优先输出最近创建group
	query := m.DB.Self.Table("groups").Select("id").Order("id desc")

	if pagination {
		query = query.Offset(offset).Limit(limit)
	}

    var count uint64

	if err := query.Scan(&grouplist).Count(&count).Error; err != nil {
		return grouplist, count, err
	}

	return grouplist, count, nil
}
























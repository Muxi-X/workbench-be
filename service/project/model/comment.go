package model

type Commenter interface {
	TableName() string
	Create(model CommentModel) error
	Update(string) error
	Delete(uint32) error
	GetModelById(uint32) error
	Verify(uint32) bool
	List(targetID, offset, limit, lastID uint32) ([]*CommentListItem, uint64, error)
}

type CommentModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Kind     uint32 `json:"kind" gorm:"column:kind;" binding:"required"` // 0 是一级，1 是二级
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Creator  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	TargetID uint32 `json:"target_id" gorm:"column:target_id;" binding:"required"`
	Re       bool   `json:"re" gorm:"column:re;" binding:"required"`
}

type CommentListItem struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Kind     uint32 `json:"kind" gorm:"column:kind;" binding:"required"` // 0 是一级，1 是二级
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Creator  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	Avatar   string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	UserName string `json:"username" gorm:"column:name;" binding:"required"`
}

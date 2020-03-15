package model

type DocModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	Content    string `json:"content" gorm:"column:content;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	Top        bool   `json:"top" gorm:"column:top;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	DeleteTime string `json:"deleteTime" gorm:"column:delete_time;" binding:"required"`
	TeamID     uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	EditorID   uint32 `json:"editorId" gorm:"column:editor_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

func (u *DocModel) TableName() string {
	return "docs"
}

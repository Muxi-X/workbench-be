package model

type FileModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	RealName   string `json:"realName" gorm:"column:realname;" binding:"required"`
	URL        string `json:"url" gorm:"column:url;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	Top        bool   `json:"top" gorm:"column:top;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	DeleteTime string `json:"deleteTime" gorm:"column:delete_time;" binding:"required"`
	TeamID     uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

func (u *FileModel) TableName() string {
	return "files"
}

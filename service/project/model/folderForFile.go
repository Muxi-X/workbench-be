package model

type FolderForFileModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:filename;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	CreatorID  string `json:"creatorID" gorm:"column:create_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

func (u *FolderForFileModel) TableName() string {
	return "foldersformds"
}

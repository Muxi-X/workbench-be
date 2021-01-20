package model

type FolderForFileModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:name;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	CreatorID  string `json:"creatorID" gorm:"column:create_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	Children   string `json:"children" gorm:"column:children;" binding:"required"`
}

func (u *FolderForFileModel) TableName() string {
	return "foldersforfiles"
}

func (u *FolderForFileModel) Update() error {
	return DB.Self.Save(u).Error
}

func GetFolderForFileModel(id uint32) (*FolderForFileModel, error) {
	s := &FolderForFileModel{}
	d := DB.Self.Table("foldersforfiles").Where("id = ?", id).First(&s)
	return s, d.Error
}

package model

type FolderForDocModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:name;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	CreatorID  string `json:"creatorID" gorm:"column:create_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	Children   string `json:"children" gorm:"column:children;" binding:"required"`
	FatherId   uint32 `json:"father_id" gorm:"column:father_id;" binding:"required"`
}

func (u *FolderForDocModel) TableName() string {
	return "foldersformds"
}

func (u *FolderForDocModel) Update() error {
	return DB.Self.Save(u).Error
}

func GetFolderForDocModel(id uint32) (*FolderForDocModel, error) {
	s := &FolderForDocModel{}
	d := DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

package model

type ProjectModel struct {
	ID           uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string `json:"name" gorm:"column:name;" binding:"required"`
	Intro        string `json:"intro" gorm:"column:intro;" binding:"required"`
	Time         string `json:"time" gorm:"column:time;" binding:"required"`
	Count        uint32 `json:"count" gorm:"column:count;" binding:"required"`
	TeamID       uint32 `json:"teamId" gorm:"column:team_id;" binding:"required"`
	FileTree     string `json:"fileTree" gorm:"column:fileTree;" binding:"required"`
	DocTree      string `json:"docTree" gorm:"column:docTree;" binding:"required"`
	FileChildren string `json:"fileChildren" gorm:"column:file_children;" binding:"required"`
	DocChildren  string `json:"docChildren" gorm:"column:doc_children;" binding:"required"`
}

type ProjectTree struct {
	FileTree string `json:"fileTree" gorm:"column:fileTree;" binding:"required"`
	DocTree  string `json:"docTree" gorm:"column:docTree;" binding:"required"`
}

type FileTreeNode struct {
	Folder        bool           `json:"folder"`
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Router        []string       `json:"router"`
	Selected      bool           `json:"selected"`
	FinalSelected bool           `json:"finalselected"`
	Children      []FileTreeNode `json:"child"`
}

func (u *ProjectModel) TableName() string {
	return "projects"
}

func GetProject(id uint32) (*ProjectModel, error) {
	s := &ProjectModel{}
	d := DB.Self.Where("id = ?", id).First(&s)
	return s, d.Error
}

func GetProjectTree(id uint32) (*ProjectTree, error) {
	s := &ProjectTree{}
	d := DB.Self.Table("projects").Where("id = ?", id).Select("fileTree,docTree").Scan(&s)
	return s, d.Error
}

type ChildrenTree struct {
	Ids  []string `json:"Ids"`
	Type string   `json:"type"` // 0->file 1->fileTree 2->doc 3->docTree
}

func (u *ProjectModel) Update() error {
	return DB.Self.Save(u).Error
}

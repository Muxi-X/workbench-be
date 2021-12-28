package model

type CommentModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Kind     uint32 `json:"kind" gorm:"column:kind;" binding:"required"` // 0 是进度，1 是文档或者文件
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Creator  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	DocID    uint32 `json:"docId" gorm:"column:doc_id;"`
	FileID   uint32 `json:"fileId" gorm:"column:file_id;"`
	StatusID uint32 `json:"statusId" gorm:"column:statu_id;"`
}

type CommentListItem struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Creator  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	Avatar   string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	UserName string `json:"username" gorm:"column:name;" binding:"required"`
}

func (c *CommentModel) TableName() string {
	return "comments"
}

// Create comments
func (c *CommentModel) Create() error {
	return DB.Self.Create(&c).Error
}

// Update comment
func (c *CommentModel) Update() error {
	return DB.Self.Save(c).Error
}

// delete comment
func DeleteComment(id, uid uint32) error {
	s := &CommentModel{}
	s.ID = id
	d := DB.Self.Where("creator = ?", uid).Delete(s)
	return d.Error
}

func GetCommentModelById(id uint32) (*CommentModel, error) {
	s := &CommentModel{}
	d := DB.Self.Where("id = ?", id).First(s)
	return s, d.Error
}

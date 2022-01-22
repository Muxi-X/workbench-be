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

func (c *CommentModel) TableName() string {
	return "comments"
}

func GetAllComments() []*CommentModel {
	var comments []*CommentModel
	DB.Self.Find(&comments)
	return comments
}

type NewCommentModel struct {
	ID       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Kind     uint32 `json:"kind" gorm:"column:kind;" binding:"required"` // 0 是一级，1 是二级
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Time     string `json:"time" gorm:"column:time;" binding:"required"`
	Creator  uint32 `json:"creator" gorm:"column:creator;" binding:"required"`
	TargetID uint32 `json:"target_id" gorm:"column:target_id;" binding:"required"`
	Re       bool   `json:"re" gorm:"column:re;" binding:"required"`
}

func CreateDocComment(comment NewCommentModel) error {
	return DB.Self.Table("comments_docs").Create(&comment).Error
}

func CreateFileComment(comment NewCommentModel) error {
	return DB.Self.Table("comments_files").Create(&comment).Error
}

func CreateStatusComment(comment NewCommentModel) error {
	return DB.Self.Table("comments_status").Create(&comment).Error
}

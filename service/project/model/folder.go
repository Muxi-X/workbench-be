package model

import (
	"errors"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
)

type FolderModel struct {
	ID         uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name       string `json:"name" gorm:"column:name;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	CreateTime string `json:"createTime" gorm:"column:create_time;" binding:"required"`
	CreatorID  uint32 `json:"creatorID" gorm:"column:create_id;" binding:"required"`
	ProjectID  uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	Children   string `json:"children" gorm:"column:children;" binding:"required"`
	FatherId   uint32 `json:"father_id" gorm:"column:father_id;" binding:"required"`
}

type FolderInfo struct {
	ID        uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name      string `json:"name" gorm:"column:name;" binding:"required"`
	ProjectID uint32 `json:"project_id" gorm:"column:project_id;" binding:"required"`
}

// FolderDetail ... 文件夹详情（doc/file）
type FolderDetail struct {
	Creator string `json:"creator" gorm:"column:creator;not null" binding:"required"`
	FolderModel
}

type FolderChildren struct {
	Children  string `json:"children" gorm:"column:children;" binding:"required"`
	ProjectID uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
}

func (f *FolderModel) Create(typeId uint8) error {
	if typeId == constvar.DocFolderCode {
		folder := FolderForDocModel{
			FolderModel: *f,
		}
		return folder.Create()
	} else if typeId == constvar.FileFolderCode {
		folder := FolderForFileModel{
			FolderModel: *f,
		}
		return folder.Create()
	} else {
		return errors.New("wrong type_id")
	}
}

func (f *FolderModel) Update(typeId uint8) error {
	if typeId == constvar.DocFolderCode {
		folder := FolderForDocModel{
			FolderModel: *f,
		}
		return folder.Update()
	} else if typeId == constvar.FileFolderCode {
		folder := FolderForFileModel{
			FolderModel: *f,
		}
		return folder.Update()
	} else {
		return errors.New("wrong type_id")
	}
}

// CreateFolder ... 事务
func CreateFolder(folder FolderModel, childrenPositionIndex uint32, typeId uint8) (uint32, error) {
	tx := m.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := folder.Create(typeId); err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	// 获取 fatherId
	fatherId := folder.FatherId
	isFatherProject := false
	if folder.FatherId == 0 {
		isFatherProject = true
		fatherId = folder.ProjectID
	}

	var err error
	if typeId == constvar.DocFolderCode {
		err = AddChildren(tx, isFatherProject, fatherId, childrenPositionIndex, &FolderForDocModel{
			FolderModel: folder,
		})
	} else if typeId == constvar.FileFolderCode {
		err = AddChildren(tx, isFatherProject, fatherId, childrenPositionIndex, &FolderForFileModel{
			FolderModel: folder,
		})
	} else {
		err = errors.New("wrong type_id")
	}

	if err != nil {
		tx.Rollback()
		return uint32(0), err
	}

	return folder.ID, tx.Commit().Error
}

package model

import (
	"fmt"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// TrashbinModel ... 回收站表
type TrashbinModel struct {
	Id        uint32 `json:"id" gorm:"column:id;" binding:"required"`
	FileId    uint32 `json:"file_id" gorm:"column:file_id;" binding:"required"`
	FileType  uint8  `json:"file_type" gorm:"column:file_type;" binding:"required"`
	Name      string `json:"name" gorm:"column:name;" binding:"required"`
	Re        bool   `json:"re" gorm:"column:re;" binding:"required"`
	ExpiresAt int64  `json:"expires_at" gorm:"column:expires_at;" binding:"required"`
}

// TrashbinListItem
type TrashbinListItem struct {
	FileId   uint32 `json:"file_id" gorm:"column:file_id;" binding:"required"`
	FileType uint8  `json:"file_type" gorm:"column:file_type;" binding:"required"`
	Name     string `json:"name" gorm:"column:name;" binding:"required"`
}

func (u *TrashbinModel) TableName() string {
	return "trashbin"
}

func (u *TrashbinModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// DeleteTrashbin ... 用户删除回收站的文件,修改 re 字段使对用户不可见
func DeleteTrashbin(fileId uint32, fileType uint8) error {
	return m.DB.Self.Table("trashbin").Where("file_id AND file_type= ?", fileId, fileType).Update("re", "1").Error
}

// DeleteTrashbinRecord 删除记录
func DeleteTrashbinRecord() error {
	t := time.Now().Unix()
	return m.DB.Self.Table("trashbin").Where("re = 1 or expires_at <= ?", t).Delete(&TrashbinModel{}).Error
}

// DeleteTrashbinRecordById ... 用户恢复文件调用
func DeleteTrashbinRecordById(id uint32, fileType uint8) error {
	return m.DB.Self.Table("trashbin").Where("file_id = ? AND file_type = ?", id, fileType).Delete(&TrashbinModel{}).Error
}

// GetTrashbinDeletedAndExpired ... 获取需要删除的文件列表
func GetTrashbinDeletedAndExpired() ([]*TrashbinListItem, error) {
	item := make([]*TrashbinListItem, 0)
	t := time.Now().Unix()
	d := m.DB.Self.Table("trashbin").Select("file_id,file_type,name").Where("re = ? or expires_at <= ?", 1, t).Scan(&item)
	if d.Error != nil {
		return nil, d.Error
	}

	return item, nil
}

// ListTrashbin list trashbin
// 用户调用查看
func ListTrashbin(offset, limit uint32) ([]*TrashbinListItem, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	trashbinList := make([]*TrashbinListItem, 0)

	query := m.DB.Self.Table("trashbin").Select("file_id,file_type,name").Where("re = ", 0).Offset(offset).Limit(limit).Order("id desc")

	if err := query.Scan(&trashbinList).Error; err != nil {
		return trashbinList, err
	}

	return trashbinList, nil
}

// --Synchronize part

// GetAllTrashbin ... 被 SynchronizeTrashbinToRedis 调用
func GetAllTrashbin() ([]*TrashbinListItem, error) {
	item := make([]*TrashbinListItem, 0)
	d := m.DB.Self.Table("trashbin").Select("file_id,file_type,name").Find(&item)
	if d.Error != nil {
		return nil, d.Error
	}

	return item, nil
}

// --Get childFolder part

// GetProjectChildFolder ... 递归获取子文件夹
func GetProjectChildFolder(id uint32, res *[]string) error {
	// 先查询子树
	project := &ProjectModel{}
	d := m.DB.Self.Where("id = ?", id).First(&project)
	if d.Error != nil {
		return d.Error
	}

	// 并入结果集
	*res = append(*res, fmt.Sprintf("%d-%d", id, constvar.ProjectCode))

	docChildren := project.DocChildren
	docRaw := strings.Split(docChildren, ",")
	for _, v := range docRaw {
		r := strings.Split(v, "-")
		if r[1] == "1" {
			// 调用自身
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = GetDocChildFolder(uint32(childId), res)
			if err != nil {
				return err
			}
		}
	}

	// 删除 file
	fileChildren := project.FileChildren
	fileRaw := strings.Split(fileChildren, ",")
	for _, v := range fileRaw {
		r := strings.Split(v, "-")
		if r[1] == "1" {
			// 调用自身
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = GetFileChildFolder(uint32(childId), res)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetDocChildFolder(id uint32, res *[]string) error {
	// 先查询子树
	docFolder := &FolderForDocModel{}
	d := m.DB.Self.Where("id = ?", id).First(&docFolder)
	if d.Error != nil {
		return d.Error
	}

	// 并入结果集
	*res = append(*res, fmt.Sprintf("%d-%d", id, constvar.DocFolderCode))

	children := docFolder.Children
	raw := strings.Split(children, ",")
	for _, v := range raw {
		r := strings.Split(v, "-")
		if r[1] == "1" {
			// 调用自身
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = GetDocChildFolder(uint32(childId), res)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetFileChildFolder(id uint32, res *[]string) error {
	// 先查询子树
	fileFolder := &FolderForFileModel{}
	d := m.DB.Self.Where("id = ?", id).First(&fileFolder)
	if d.Error != nil {
		return d.Error
	}

	// 并入结果集
	*res = append(*res, fmt.Sprintf("%d-%d", id, constvar.FileFolderCode))

	children := fileFolder.Children
	raw := strings.Split(children, ",")
	for _, v := range raw {
		r := strings.Split(v, "-")
		if r[1] == "1" {
			// 调用自身
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = GetFileChildFolder(uint32(childId), res)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// --Goroutine part

// TidyTrashbin ... 事务,被协程调用
func TidyTrashbin(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查找用户删除的文件
	deletedList, err := GetTrashbinDeletedAndExpired()
	if err != nil {
		tx.Rollback()
		return err
	}

	var res []string
	for _, v := range deletedList {
		// 修改原表 re 字段 和 获取子文件
		switch v.FileType {
		case 0:
			err = DeleteProjectTrashbin(v.FileId, &res)
		case 1:
			err = DeleteDocTrashbin(v.FileId)
			// 单独并入结果集
			res = append(res, fmt.Sprintf("%d-%d", v.FileId, constvar.DocCode))
		case 2:
			err = DeleteFileTrashbin(v.FileId)
			// 单独并入结果集
			res = append(res, fmt.Sprintf("%d-%d", v.FileId, constvar.FileCode))
		case 3:
			err = DeleteDocFolderTrashbin(v.FileId, &res)
		case 4:
			err = DeleteFileFolderTrashbin(v.FileId, &res)
		}

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 同步 redis
	if err := m.SRemToRedis(constvar.Trashbin, res); err != nil {
		tx.Rollback()
		return err
	}

	// 删除 trashbin 记录
	if err = DeleteTrashbinRecord(); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// --Delete part

// DeleteProjectTrashbin ... 递归调用,子协程调用
func DeleteProjectTrashbin(id uint32, res *[]string) error {
	// 先查询子树
	project := &ProjectModel{}
	d := m.DB.Self.Where("id = ?", id).First(&project)
	if d.Error != nil {
		return d.Error
	}

	// 并入结果集
	*res = append(*res, fmt.Sprintf("%d-%d", id, constvar.ProjectCode))

	docChildren := project.DocChildren
	docRaw := strings.Split(docChildren, ",")
	for _, v := range docRaw {
		r := strings.Split(v, "-")
		if r[1] == "0" {
			// 调用 delete doc
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = DeleteDocTrashbin(uint32(childId))
			if err != nil {
				return err
			}
		} else {
			// 调用自身
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = DeleteDocFolderTrashbin(uint32(childId), res)
			if err != nil {
				return err
			}
		}
	}

	// 删除 file
	fileChildren := project.FileChildren
	fileRaw := strings.Split(fileChildren, ",")
	for _, v := range fileRaw {
		r := strings.Split(v, "-")
		if r[1] == "0" {
			// 调用 delete doc
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = DeleteFileTrashbin(uint32(childId))
			if err != nil {
				return err
			}
		} else {
			// 调用自身
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = DeleteFileFolderTrashbin(uint32(childId), res)
			if err != nil {
				return err
			}
		}
	}

	return m.DB.Self.Delete(&project).Error
}

func DeleteDocTrashbin(id uint32) error {
	return m.DB.Self.Table("docs").Where("id = ?", id).Update("re", "1").Error
}

func DeleteFileTrashbin(id uint32) error {
	return m.DB.Self.Table("files").Where("id = ?", id).Update("re", "1").Error
}

func DeleteDocFolderTrashbin(id uint32, res *[]string) error {
	// 先查询子树
	docFolder := &FolderForDocModel{}
	d := m.DB.Self.Where("id = ?", id).First(&docFolder)
	if d.Error != nil {
		return d.Error
	}

	// 并入结果集
	*res = append(*res, fmt.Sprintf("%d-%d", id, constvar.DocFolderCode))

	children := docFolder.Children
	raw := strings.Split(children, ",")
	for _, v := range raw {
		r := strings.Split(v, "-")
		if r[1] == "0" {
			// 调用 delete doc
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = DeleteDocTrashbin(uint32(childId))
			if err != nil {
				return err
			}
		} else {
			// 调用自身
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = DeleteDocFolderTrashbin(uint32(childId), res)
			if err != nil {
				return err
			}
		}
	}

	return m.DB.Self.Table("folderformds").Where("id = ?", id).Update("re", "1").Error
}

func DeleteFileFolderTrashbin(id uint32, res *[]string) error {
	// 先查询子树
	fileFolder := &FolderForFileModel{}
	d := m.DB.Self.Where("id = ?", id).First(&fileFolder)
	if d.Error != nil {
		return d.Error
	}

	// 并入结果集
	*res = append(*res, fmt.Sprintf("%d-%d", id, constvar.FileFolderCode))

	children := fileFolder.Children
	raw := strings.Split(children, ",")
	for _, v := range raw {
		r := strings.Split(v, "-")
		if r[1] == "0" {
			// 调用 delete doc
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = DeleteFileTrashbin(uint32(childId))
			if err != nil {
				return err
			}
		} else {
			// 调用自身
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = DeleteFileFolderTrashbin(uint32(childId), res)
			if err != nil {
				return err
			}
		}
	}

	return m.DB.Self.Table("folderforfile").Where("id = ?", id).Update("re", "1").Error
}

// --Remove part

// RemoveTrashbin ... 从回收站恢复文件
// 事务
func RemoveTrashbin(db *gorm.DB, fileId uint32, fileType uint8, fatherType bool, fatherId, childrenPositionIndex uint32) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 删除回收站记录
	if err := DeleteTrashbinRecordById(fileId, fileType); err != nil {
		tx.Rollback()
		return err
	}

	// 同步 redis
	// 需要找到子文件
	var res []string
	var err error
	switch fileType {
	case 0:
		err = GetProjectChildFolder(fileId, &res)
	case 1:
		res = append(res, fmt.Sprintf("%d-%d", fileId, constvar.DocCode))
	case 2:
		res = append(res, fmt.Sprintf("%d-%d", fileId, constvar.FileCode))
	case 3:
		err = GetDocChildFolder(fileId, &res)
	case 4:
		err = GetFileChildFolder(fileId, &res)
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	// 同步 redis
	if len(res) == 0 {
		if err = m.SRemToRedis(constvar.Trashbin, res); err != nil {
			tx.Rollback()
			return err
		}
	}

	// 恢复文件树
	// 分类 project doc file docfolder filefolder
	switch fileType {
	case 0:
		err = RecoverProject(fileId)
	case 1:
		err = AddDocChildren(fatherType, fatherId, childrenPositionIndex, &DocModel{ID: fileId})
	case 2:
		err = AddFileChildren(fatherType, fatherId, childrenPositionIndex, &FileModel{ID: fileId})
	case 3:
		err = AddDocChildren(fatherType, fatherId, childrenPositionIndex, &FolderForDocModel{ID: fileId})
	case 4:
		err = AddFileChildren(fatherType, fatherId, childrenPositionIndex, &FolderForFileModel{ID: fileId})
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

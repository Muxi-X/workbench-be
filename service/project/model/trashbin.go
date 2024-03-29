package model

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// TrashbinModel ... 回收站表
type TrashbinModel struct {
	Id         uint32 `json:"id" gorm:"column:id;" binding:"required"`
	FileId     uint32 `json:"file_id" gorm:"column:file_id;" binding:"required"`
	FileType   uint8  `json:"file_type" gorm:"column:file_type;" binding:"required"`
	Name       string `json:"name" gorm:"column:name;" binding:"required"`
	Re         bool   `json:"re" gorm:"column:re;" binding:"required"`
	ExpiresAt  int64  `json:"expires_at" gorm:"column:expires_at;" binding:"required"`
	DeleteTime string `json:"delete_time" gorm:"column:delete_time;" binding:"required"`
	CreateTime string `json:"create_time" gorm:"column:create_time;" binding:"required"`
	ProjectID  uint32 `json:"project_id" gorm:"column:project_id;" binding:"required"`
}

// TrashbinListItem ...
type TrashbinListItem struct {
	FileId     uint32 `json:"file_id" gorm:"column:file_id;" binding:"required"`
	FileType   uint8  `json:"file_type" gorm:"column:file_type;" binding:"required"`
	Name       string `json:"name" gorm:"column:name;" binding:"required"`
	DeleteTime string `json:"delete_time" gorm:"column:delete_time;" binding:"required"`
	CreateTime string `json:"create_time" gorm:"column:create_time;" binding:"required"`
}

func (trashbin *TrashbinModel) TableName() string {
	return "trashbin"
}

func (trashbin *TrashbinModel) Create() error {
	return m.DB.Self.Create(&trashbin).Error
}

// DeleteTrashbin ... 用户删除回收站的文件,修改 re 字段使对用户不可见
func DeleteTrashbin(fileId uint32, fileType uint8, projectId uint32) error {
	return m.DB.Self.Table("trashbin").Where("file_id = ? AND file_type = ? AND project_id = ? ", fileId, fileType, projectId).Update("re", "1").Error
}

// DeleteTrashbinRecord 删除记录
func DeleteTrashbinRecord() error {
	t := time.Now().Unix()
	return m.DB.Self.Table("trashbin").Where("re = 1 or expires_at <= ?", t).Delete(&TrashbinModel{}).Error
}

// DeleteTrashbinRecordByIdAndProjectId ... 用户恢复文件调用
func DeleteTrashbinRecordByIdAndProjectId(id uint32, fileType uint8, projectId uint32) error {
	return m.DB.Self.Table("trashbin").Where("file_id = ? AND file_type = ? AND project_id = ?", id, fileType, projectId).Delete(&TrashbinModel{}).Error
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
func ListTrashbin(offset, limit, projectId uint32) ([]*TrashbinListItem, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	trashbinList := make([]*TrashbinListItem, 0)

	query := m.DB.Self.Table("trashbin").Where("re = 0  AND project_id = ?", projectId).Offset(offset).Limit(limit).Order("id desc")

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

func GetChildFolder(id uint32, res *[]string, typeId uint8) error {
	// 先查询子树
	folder := &FolderModel{}
	var err error
	if typeId == constvar.DocCode {
		err = m.DB.Self.Table("foldersformds").Where("id = ?", id).First(&folder).Error
	} else if typeId == constvar.FileCode {
		err = m.DB.Self.Table("foldersforfiles").Where("id = ?", id).First(&folder).Error
	} else {
		return errors.New("wrong type_id")
	}
	if err != nil {
		return err
	}
	children := folder.Children

	// 并入结果集
	*res = append(*res, fmt.Sprintf("%d-%d", id, typeId))
	raw := strings.Split(children, ",")
	if len(raw) <= 1 {
		return nil
	}
	for _, v := range raw {
		r := strings.Split(v, "-")
		if r[1] == "1" {
			// 调用自身
			childId, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			err = GetChildFolder(uint32(childId), res, typeId)
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
		case constvar.DocCode:
			err = DeleteDocTrashbin(v.FileId)
			// 单独并入结果集
			res = append(res, fmt.Sprintf("%d-%d", v.FileId, constvar.DocCode))
		case constvar.FileCode:
			err = DeleteFileTrashbin(v.FileId)
			// 单独并入结果集
			res = append(res, fmt.Sprintf("%d-%d", v.FileId, constvar.FileCode))
		case constvar.DocFolderCode:
			err = DeleteDocFolderTrashbin(v.FileId, &res)
		case constvar.FileFolderCode:
			err = DeleteFileFolderTrashbin(v.FileId, &res)
		default:
			err = errors.New("wrong type_id")
		}

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(res) == 0 {
		return tx.Commit().Error
	}

	// 同步 redis
	// if err := m.SRemToRedis(constvar.Trashbin, res); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// 删除 trashbin 记录
	if err = DeleteTrashbinRecord(); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// --Delete part

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

	return m.DB.Self.Table("foldersformds").Where("id = ?", id).Update("re", "1").Error
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

	return m.DB.Self.Table("foldersforfiles").Where("id = ?", id).Update("re", "1").Error
}

// --Recover part

// RecoverTrashbin ... 从回收站恢复文件
// 事务
func RecoverTrashbin(db *gorm.DB, fileId uint32, fileType uint8, isFatherProject bool, fatherId, childrenPositionIndex, projectID uint32) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 删除回收站记录
	if err := DeleteTrashbinRecordByIdAndProjectId(fileId, fileType, projectID); err != nil {
		tx.Rollback()
		return err
	}

	// 同步 redis
	// 需要找到子文件
	var res []string
	var err error
	switch fileType {
	case constvar.DocCode:
		res = append(res, fmt.Sprintf("%d-%d", fileId, constvar.DocCode))
	case constvar.FileCode:
		res = append(res, fmt.Sprintf("%d-%d", fileId, constvar.FileCode))
	case constvar.DocFolderCode:
		err = GetChildFolder(fileId, &res, fileType)
	case constvar.FileFolderCode:
		err = GetChildFolder(fileId, &res, fileType)
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	// 同步 redis
	if len(res) != 0 {
		if err = m.SRemToRedis(constvar.Trashbin, res); err != nil {
			tx.Rollback()
			return err
		}
	}

	// 恢复文件树
	// 分类 project doc file docfolder filefolder
	switch fileType {
	case constvar.DocCode:
		err = AddChildren(tx, isFatherProject, fatherId, childrenPositionIndex, &DocModel{ID: fileId})
	case constvar.FileCode:
		err = AddChildren(tx, isFatherProject, fatherId, childrenPositionIndex, &FileModel{ID: fileId})
	case constvar.DocFolderCode:
		err = AddChildren(tx, isFatherProject, fatherId, childrenPositionIndex, &FolderForDocModel{FolderModel: FolderModel{
			ID: fileId,
		}})
	case constvar.FileFolderCode:
		err = AddChildren(tx, isFatherProject, fatherId, childrenPositionIndex, &FolderForFileModel{FolderModel: FolderModel{
			ID: fileId,
		}})
	default:
		err = errors.New("wrong type_id")
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// --adjust part 用来判断文件是否在 redis 回收站列表

// AdjustSelfIfExist ... 判断自身是否存在，多用于 folder
func AdjustSelfIfExist(id uint32, code uint8) (bool, error) {
	target := fmt.Sprintf("%d-%d", id, code) // code 表示文件编码
	isDeleted, err := m.SIsmembersFromRedis(constvar.Trashbin, target)
	return isDeleted, err
}

// AdjustSelfAndFatherIfExist ... 判断自身和父节点是否存在，多用于 file
func AdjustSelfAndFatherIfExist(id, fatherId uint32, code, fatherCode uint8) (bool, error) {
	self := fmt.Sprintf("%d-%d", id, code)
	father := fmt.Sprintf("%d-%d", fatherId, fatherCode)
	isDeleted, err := m.SIsmembersFromRedis(constvar.Trashbin, self, father)
	return isDeleted, err
}

// AdjustFolderListIfExist ... 过滤文件夹列表
func AdjustFolderListIfExist(list []uint32, code uint8) ([]uint32, error) {
	var scope []uint32
	for _, v := range list {
		isDeleted, err := AdjustSelfIfExist(v, code)
		if err != nil {
			return nil, err
		}
		if !isDeleted { // 存在 redis 返回 1, 说明被删
			scope = append(scope, v)
		}
	}

	return scope, nil
}

func AdjustFileListIfExist(list []uint32, fatherId uint32, code, fatherCode uint8) ([]uint32, error) {
	// 先判断父文件夹是否被删
	// 这是用来获取一个文件夹下的全部文件，父 id 都是一样的
	isDeleted, err := AdjustSelfIfExist(fatherId, fatherCode)
	if err != nil {
		return nil, err
	}
	if isDeleted {
		return nil, nil
	}

	var scope []uint32
	for _, v := range list {
		// 父 id 已经判断，这里判断自身就可以了
		isDeleted, err := AdjustSelfIfExist(v, code)
		if err != nil {
			return nil, err
		}
		if !isDeleted { // 存在 redis 返回 1, 说明被删
			scope = append(scope, v)
		}
	}

	return scope, nil
}

func (trashbin *TrashbinModel) DeleteChildren(fatherId uint32, isFatherProject bool) error {
	tx := m.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 获取时间
	day := viper.GetInt("trashbin.expired")
	t := time.Now().Unix()
	trashbin.ExpiresAt = t + int64(time.Hour*24*time.Duration(day))

	if err := tx.Create(trashbin).Error; err != nil {
		tx.Rollback()
		return err
	}

	typeId := trashbin.FileType
	if trashbin.FileType >= 3 { // folder
		// 获取子文件，同步 redis
		typeId -= 2
		var res []string
		if err := GetChildFolder(trashbin.FileId, &res, typeId); err != nil {
			tx.Rollback()
			return err
		}
		if len(res) != 0 {
			if err := m.SAddToRedis(constvar.Trashbin, res); err != nil {
				tx.Rollback()
				return err
			}
		}
	} else {
		// 同步 redis
		// 不需要找子文件夹
		if err := m.SAddToRedis(constvar.Trashbin,
			fmt.Sprintf("%d-%d", trashbin.FileId, typeId)); err != nil {
			tx.Rollback()
			return err
		}
	}

	// 修改文件树
	if err := DeleteChildren(tx, isFatherProject, fatherId, trashbin.FileId, typeId); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

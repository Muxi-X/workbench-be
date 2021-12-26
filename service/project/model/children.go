package model

import (
	"errors"
	"fmt"
	"muxi-workbench-project/errno"
	"muxi-workbench/pkg/constvar"
	"strings"

	"github.com/jinzhu/gorm"
)

// 放修改文件树的函数，包括 add 和 delete

// addChildren ... 返回新 children
func addChildren(children string, id, childrenPositionIndex, isFolder uint32) (string, error) {
	if childrenPositionIndex == 0 { // 插在开头
		if children != "" {
			children = fmt.Sprintf("%d-%d,%s", id, isFolder, children)
		} else {
			children = fmt.Sprintf("%d-%d", id, isFolder)
		}
		return children, nil
	}

	// 根据 childrenPositionIndex 判断插入位置，从 0 计数
	index := 0
	count := 0
	for k, v := range children {
		if v == ',' {
			count++
		}

		if count == int(childrenPositionIndex) {
			index = k + 1 // index+1 取逗号后一位，即下一个文件的开头
			break
		}
	}

	if count == int(childrenPositionIndex) { // 插在中间的情况
		children = fmt.Sprintf("%s%d-%d,%s", children[:index], id, isFolder, children[index:])
	} else if count+1 == int(childrenPositionIndex) { // 插在结尾
		children = fmt.Sprintf("%s,%d-%d", children, id, isFolder)
	} else {
		return "", errors.New("invalid children position index")
	}

	return children, nil
}

func deleteChildren(children string, id uint32, isFolder uint8) (string, error) {
	file := fmt.Sprintf("%d-%d", id, isFolder)
	index := strings.Index(children, file)
	if index == -1 {
		return "", errno.ErrFileNotFound
	}

	nextIndex := strings.Index(children[index:], ",")
	if nextIndex == -1 { // 找不到，说明当前文件是最后的文件
		children = children[:index]
		if len(children) > 0 && children[len(children)-1] == ',' { // 去掉末尾逗号
			children = children[:len(children)-1]
		}
	} else { // 找到了，说明是中间的文件
		children = children[:index] + children[index+nextIndex+1:]
	}

	return children, nil
}

// DeleteChildren ... 删除 树
func DeleteChildren(tx *gorm.DB, isFatherProject bool, fatherId, id uint32, typeId uint8) error {
	// 修改文件树
	isFolder := (typeId - 1) % 2

	var getFolderModel func(uint32) (*FolderModel, error)

	var code uint8
	switch typeId {
	case constvar.DocCode:
		code = 1
	case constvar.FileCode:
		code = 2
	case constvar.DocFolderCode:
		getFolderModel = GetFolderForDocModel
	case constvar.FileFolderCode:
		getFolderModel = GetFolderForFileModel
	default:
		return errors.New("wrong type_id")
	}

	if isFatherProject {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}

		if code == 1 {
			newChildren, err := deleteChildren(item.DocChildren, id, isFolder)
			if err != nil {
				return err
			}

			item.DocChildren = newChildren
		} else if code == 2 {
			newChildren, err := deleteChildren(item.FileChildren, id, isFolder)
			if err != nil {
				return err
			}

			item.FileChildren = newChildren
		}
		if err := tx.Save(item).Error; err != nil {
			return err
		}
	} else {
		item, err := getFolderModel(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := deleteChildren(item.Children, id, isFolder)
		if err != nil {
			return err
		}

		item.Children = newChildren

		if err := tx.Save(item).Error; err != nil {
			return err
		}
	}

	return nil
}

func AddChildren(tx *gorm.DB, isFatherProject bool, fatherId, childrenPositionIndex uint32, obj interface{}) error {
	var id uint32
	var code uint8
	var isFolder uint32

	var getFolderModel func(uint32) (*FolderModel, error)

	switch obj.(type) {
	case *DocModel:
		id = obj.(*DocModel).ID
		code = constvar.DocCode
		getFolderModel = GetFolderForDocModel
	case *FileModel:
		id = obj.(*FileModel).ID
		code = constvar.FileCode
		getFolderModel = GetFolderForFileModel
	case *FolderForDocModel:
		id = obj.(*FolderForDocModel).ID
		code = constvar.DocFolderCode
		getFolderModel = GetFolderForDocModel
		isFolder = uint32(1)
	case *FolderForFileModel:
		id = obj.(*FolderForFileModel).ID
		code = constvar.FileFolderCode
		getFolderModel = GetFolderForFileModel
		isFolder = uint32(1)
	default:
		return errors.New("wrong type_id")
	}

	if isFatherProject {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}
		if code == constvar.DocCode || code == constvar.DocFolderCode {
			newChildren, err := addChildren(item.DocChildren, id, childrenPositionIndex, isFolder)
			if err != nil {
				return err
			}

			item.DocChildren = newChildren
		} else if code == constvar.FileCode || code == constvar.FileFolderCode {
			newChildren, err := addChildren(item.FileChildren, id, childrenPositionIndex, isFolder)
			if err != nil {
				return err
			}

			item.FileChildren = newChildren
		}
		if err := tx.Save(item).Error; err != nil {
			return err
		}
	} else {
		item, err := getFolderModel(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := addChildren(item.Children, id, childrenPositionIndex, isFolder)
		if err != nil {
			return err
		}

		item.Children = newChildren

		if code == constvar.DocCode || code == constvar.DocFolderCode {
			err = tx.Table("foldersformds").Save(item).Error
		} else if code == constvar.FileCode || code == constvar.FileFolderCode {
			err = tx.Table("foldersforfiles").Save(item).Error
		}
		if err != nil {
			return err
		}

	}

	return nil
}

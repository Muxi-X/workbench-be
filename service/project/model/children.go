package model

import (
	"errors"
	"fmt"
	"muxi-workbench-project/errno"
	"strings"

	"github.com/jinzhu/gorm"
)

// 放修改文件树的函数，包括 add 和 delete

// AddChildren ... 返回新 children
func AddChildren(children string, id, childrenPositionIndex, isFolder uint32) (string, error) {
	if childrenPositionIndex == 0 { // 插在开头
		children = fmt.Sprintf("%d-%d,%s", id, isFolder, children)
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
		}
	}

	if count == int(childrenPositionIndex) { // 插在中间的情况
		children = fmt.Sprintf("%s%d-%d,%s", children[:index], id, isFolder, children[index:])
	} else if count+1 == int(childrenPositionIndex) { // 插在结尾
		children = fmt.Sprintf("%s,%d-%d", children, id, isFolder)
	} else {
		return "", errors.New("Invalid children position index.")
	}

	return children, nil
}

func DeleteChildren(children string, id uint32, isFolder uint8) (string, error) {
	file := fmt.Sprintf("%d-%d", id, isFolder)
	index := strings.Index(children, file)
	if index == -1 {
		return "", errno.ErrFileNotFound
	}

	nextIndex := strings.Index(children[index:], ",")
	if nextIndex == -1 { // 找不到，说明当前文件是最后的文件
		children = children[:index]
	} else { // 找到了，说明是中间的文件
		children = children[:index] + children[nextIndex:]
	}

	return children, nil
}

// AddDocChildren ... 新增 doc 文件树
func AddDocChildren(tx *gorm.DB, isFatherProject bool, fatherId, childrenPositionIndex uint32, obj interface{}) error {
	var id uint32
	var isFolder uint32 // 0->file 1->folder

	switch obj.(type) {
	case *DocModel:
		id = obj.(*DocModel).ID
	case *FolderForDocModel:
		id = obj.(*FolderForDocModel).ID
		isFolder = uint32(1)
	}

	if isFatherProject {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := AddChildren(item.DocChildren, id, childrenPositionIndex, isFolder)
		if err != nil {
			return err
		}

		item.DocChildren = newChildren

		if err := tx.Save(item).Error; err != nil {
			return err
		}
	} else {
		item, err := GetFolderForDocModel(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := AddChildren(item.Children, id, childrenPositionIndex, isFolder)
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

// DeleteDocChildren ... 删除 文档 树
func DeleteDocChildren(tx *gorm.DB, isFatherProject bool, fatherId, id uint32, isFolder uint8) error {
	// 修改文件树
	if isFatherProject {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := DeleteChildren(item.DocChildren, id, isFolder)
		if err != nil {
			return err
		}

		item.DocChildren = newChildren

		if err := tx.Save(item).Error; err != nil {
			return err
		}
	} else {
		item, err := GetFolderForDocModel(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := DeleteChildren(item.Children, id, isFolder)
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

func AddFileChildren(tx *gorm.DB, isFatherProject bool, fatherId, childrenPositionIndex uint32, obj interface{}) error {
	var id uint32
	var isFolder uint32

	switch obj.(type) {
	case *DocModel:
		id = obj.(*FileModel).ID
	case *FolderForDocModel:
		id = obj.(*FolderForFileModel).ID
		isFolder = uint32(1)
	}

	if isFatherProject {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := AddChildren(item.FileChildren, id, childrenPositionIndex, isFolder)
		if err != nil {
			return err
		}

		item.FileChildren = newChildren

		if err := tx.Save(item).Error; err != nil {
			return err
		}
	} else {
		item, err := GetFolderForFileModel(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := AddChildren(item.Children, id, childrenPositionIndex, isFolder)
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

func DeleteFileChildren(tx *gorm.DB, isFatherProject bool, fatherId, id uint32, isFolder uint8) error {
	if isFatherProject {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := DeleteChildren(item.FileChildren, id, isFolder)
		if err != nil {
			return err
		}

		item.DocChildren = newChildren

		if err := tx.Save(item).Error; err != nil {
			return err
		}
	} else {
		item, err := GetFolderForFileModel(fatherId)
		if err != nil {
			return err
		}

		newChildren, err := DeleteChildren(item.Children, id, isFolder)
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

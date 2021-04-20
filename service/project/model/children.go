package model

import (
	"errors"
	"fmt"
	"muxi-workbench-project/errno"
	"strings"
)

// 放修改文件树的函数，包括 add 和 delete

// AddChildren ... 返回新 children
func AddChildren(children string, id, childrenPositionIndex, isFolder uint32) (string, error) {
	// 根据 childrenPositionIndex 判断插入位置，从 0 计数
	index := int(childrenPositionIndex) * 4
	if index-1 < len(children) {
		children = fmt.Sprintf("%s%d-%d,%s", children[:index], id, isFolder, children[index:])
	} else if index-1 == len(children) {
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

	// index + 4 刚好取到下一个片段的第一个元素，以此区分是否在末尾
	if index+4 < len(children) {
		children = children[:index] + children[index+4:]
	} else if index+4 > len(children) {
		children = children[:index]
	} else {
		return "", errno.ErrInvalidIndex
	}

	return children, nil
}

// AddDocChildren ... 新增 doc 文件树
func AddDocChildren(isFatherProject bool, fatherId, childrenPositionIndex uint32, obj interface{}) error {
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

		if err := item.Update(); err != nil {
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

		if err := item.Update(); err != nil {
			return err
		}
	}

	return nil
}

// DeleteDocChildren ... 删除 文档 树
func DeleteDocChildren(isFatherProject bool, fatherId, id uint32, isFolder uint8) error {
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

		if err := item.Update(); err != nil {
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

		if err := item.Update(); err != nil {
			return err
		}
	}

	return nil
}

func AddFileChildren(isFatherProject bool, fatherId, childrenPositionIndex uint32, obj interface{}) error {
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

		if err := item.Update(); err != nil {
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

		if err := item.Update(); err != nil {
			return err
		}
	}

	return nil
}

func DeleteFileChildren(isFatherProject bool, fatherId, id uint32, isFolder uint8) error {
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

		if err := item.Update(); err != nil {
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

		if err := item.Update(); err != nil {
			return err
		}
	}

	return nil
}

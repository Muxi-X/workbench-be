package model

import (
	"errors"
	"fmt"
)

// 放修改文件树的函数，包括 add 和 delete

// AddDocChildren ... 新增 doc 文件树
func AddDocChildren(fatherType bool, fatherId, childrenPositionIndex uint32, obj interface{}) error {
	var id uint32
	var isfolder uint32 // 0->file 1->folder

	switch obj.(type) {
	case *DocModel:
		id = obj.(*DocModel).ID
	case *FolderForDocModel:
		id = obj.(*FolderForDocModel).ID
		isfolder = uint32(1)
	}

	if fatherType {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}

		// 根据 childrenPositionIndex 判断插入位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.DocChildren) {
			item.DocChildren = fmt.Sprintf("%s%d-%d,%s", item.DocChildren[:index], id, isfolder, item.DocChildren[index:])
		} else if index-1 == len(item.DocChildren) {
			item.DocChildren = fmt.Sprintf("%s,%d-%d", item.DocChildren, id, isfolder)
		} else {
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			return err
		}
	} else {
		item, err := GetFolderForDocModel(fatherId)
		if err != nil {
			return err
		}

		// 根据 childrenPositionIndex 判断插入位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.Children) {
			item.Children = fmt.Sprintf("%s%d-%d,%s", item.Children[:index], id, isfolder, item.Children[index:])
		} else if index-1 == len(item.Children) {
			item.Children = fmt.Sprintf("%s,%d-%d", item.Children, id, isfolder)
		} else {
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			return err
		}
	}

	return nil
}

// DeleteDocChildren ... 删除 文档 树
func DeleteDocChildren(fatherType bool, fatherId, childrenPositionIndex uint32) error {
	// 修改文件树
	if fatherType {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}

		// 根据 childrenPositionIndex 判断删除位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.DocChildren) {
			item.DocChildren = item.DocChildren[:index] + item.DocChildren[index+1:]
		} else if index-1 == len(item.DocChildren) {
			item.DocChildren = item.DocChildren[:index]
		} else {
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			return err
		}
	} else {
		item, err := GetFolderForDocModel(fatherId)
		if err != nil {
			return err
		}

		// 根据 childrenPositionIndex 判断删除位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.Children) {
			item.Children = item.Children[:index] + item.Children[index+1:]
		} else if index-1 == len(item.Children) {
			item.Children = item.Children[:index]
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			return err
		}
	}

	return nil
}

func AddFileChildren(fatherType bool, fatherId, childrenPositionIndex uint32, obj interface{}) error {
	var id uint32
	var isfolder uint32

	switch obj.(type) {
	case *DocModel:
		id = obj.(*FileModel).ID
	case *FolderForDocModel:
		id = obj.(*FolderForFileModel).ID
		isfolder = uint32(1)
	}

	if fatherType {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}

		// 根据 childrenPositionIndex 判断插入位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.FileChildren) {
			item.FileChildren = fmt.Sprintf("%s%d-%d,%s", item.FileChildren[:index], id, isfolder, item.FileChildren[index:])
		} else if index-1 == len(item.DocChildren) {
			item.FileChildren = fmt.Sprintf("%s,%d-%d", item.FileChildren, id, isfolder)
		} else {
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			return err
		}
	} else {
		item, err := GetFolderForFileModel(fatherId)
		if err != nil {
			return err
		}

		// 根据 childrenPositionIndex 判断插入位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.Children) {
			item.Children = fmt.Sprintf("%s%d-%d,%s", item.Children[:index], id, isfolder, item.Children[index:])
		} else if index-1 == len(item.Children) {
			item.Children = fmt.Sprintf("%s,%d-%d", item.Children, id, isfolder)
		} else {
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			return err
		}
	}

	return nil
}

func DeleteFileChildren(fatherType bool, fatherId, childrenPositionIndex uint32) error {
	if fatherType {
		// 查询结果，解析 children 再更新
		item, err := GetProject(fatherId)
		if err != nil {
			return err
		}

		// 根据 childrenPositionIndex 判断删除位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.FileChildren) {
			item.FileChildren = item.FileChildren[:index] + item.FileChildren[index+1:]
		} else if index-1 == len(item.FileChildren) {
			item.FileChildren = item.FileChildren[:index]
		} else {
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			return err
		}
	} else {
		item, err := GetFolderForFileModel(fatherId)
		if err != nil {
			return err
		}

		// 根据 childrenPositionIndex 判断删除位置，从 0 计数
		index := int(childrenPositionIndex) * 4
		if index-1 < len(item.Children) {
			item.Children = item.Children[:index] + item.Children[index+1:]
		} else if index-1 == len(item.Children) {
			item.Children = item.Children[:index]
			return errors.New("Invalid children position index.")
		}

		if err := item.Update(); err != nil {
			return err
		}
	}

	return nil
}

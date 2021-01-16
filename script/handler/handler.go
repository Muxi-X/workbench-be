package handler

import (
	"PROJECT_SCRIPT/log"
	"PROJECT_SCRIPT/model"
	"encoding/json"
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

// 获取数据
func getData(id uint32) (*model.FileTreeNode, *model.FileTreeNode, error) {
	s, err := model.GetProjectTree(id)
	if err != nil {
		log.Error("getProjectTree fatal",
			zap.String("cause:", err.Error()))
		return nil, nil, err
	}

	file := &model.FileTreeNode{}
	json.Unmarshal([]byte(s.FileTree), file)

	doc := &model.FileTreeNode{}
	json.Unmarshal([]byte(s.DocTree), doc)

	return doc, file, nil
}

func writeString(s1 string, s2 string) string {
	var result string
	if s1 == "" {
		result = fmt.Sprintf("%s", s2)
	} else {
		result = fmt.Sprintf("%s,%s", s1, s2)
	}
	return result
}

func writeItem(tree *model.FileTreeNode) (string, []model.FileTreeNode) {
	var file string
	var fileFolder string // 返回给下个函数
	fileFolderItem := make([]model.FileTreeNode, 0)
	for _, v := range tree.Children {
		if v.Folder {
			fileFolder = writeString(fileFolder, v.Id)
			fileFolderItem = append(fileFolderItem, v)
		} else {
			file = writeString(file, v.Id)
		}
	}

	fileChildren := fmt.Sprintf("%s;%s", fileFolder, file)

	return fileChildren, fileFolderItem
}

// 更新信息到 project 表
func updateDataToProject(doc, file *model.FileTreeNode, id uint32) ([]model.FileTreeNode, []model.FileTreeNode, error) {
	// 插入当前节点到 project docChildren 和 fileChildren 事务
	s, err := model.GetProject(id)
	if err != nil {
		// error
		log.Error("get total project fatal",
			zap.String("cause:", err.Error()))
		return nil, nil, err
	}

	u, err := model.GetProjectTree(id)
	if err != nil {
		log.Error("getProjectTree fatal",
			zap.String("cause:", err.Error()))
		return nil, nil, err
	}

	// 写入 fileChildren，文件夹写在前面，文件写在后面->1,5;2,4,5,6
	fileChildren, fileFolderItem := writeItem(file)

	// 写入 docChildren
	docChildren, docFolderItem := writeItem(doc)

	fmt.Println(fileChildren, docChildren)

	s.FileChildren = fileChildren
	s.DocChildren = docChildren
	s.FileTree = u.FileTree
	s.DocTree = u.DocTree

	err = s.Update()
	if err != nil {
		log.Error("update project fatal",
			zap.String("cause:", err.Error()))
	}

	return fileFolderItem, docFolderItem, nil
}

func writeFolderItem(tree []model.FileTreeNode) string {
	var item string
	var itemFolder string
	for _, v := range tree {
		if v.Folder {
			itemFolder = writeString(itemFolder, v.Id)
		} else {
			item = writeString(item, v.Id)
		}
	}

	children := fmt.Sprintf("%s;%s", itemFolder, item)
	return children
}

func insertDatabaseForFile(tree []model.FileTreeNode, id string) {
	// 将当前节点子节点都插入
	idInt, _ := strconv.Atoi(id)
	s, err := model.GetFolderForFileModel(uint32(idInt))
	if err != nil {
		//error
	}

	children := writeFolderItem(tree)

	s.Children = children

	err = s.Update()
	if err != nil {
		fmt.Println(err)
	}
}

// 写入信息到 fileFolder
func updateDataToFileFolder(fileTree []model.FileTreeNode) {
	for fileTree != nil {
		v := fileTree[0]
		if v.Children != nil {
			insertDatabaseForFile(v.Children, v.Id) // 插入结果，将item全部插入
			fileTree = append(fileTree, v.Children...)
		}
		if len(fileTree) == 1 {
			break
		}
		fileTree = fileTree[1:]
	}
}

func insertDatabaseForDoc(tree []model.FileTreeNode, id string) {
	// 将当前节点子节点都插入
	idInt, _ := strconv.Atoi(id)
	s, err := model.GetFolderForDocModel(uint32(idInt))
	if err != nil {
		fmt.Println(err)
	}

	children := writeFolderItem(tree)

	s.Children = children

	err = s.Update()
	if err != nil {
		fmt.Println(err)
	}
}

// 写入信息到 docFolder
func updateDataToDocFolder(fileTree []model.FileTreeNode) {
	for fileTree != nil {
		v := fileTree[0]
		if v.Children != nil {
			insertDatabaseForDoc(v.Children, v.Id) // 插入结果，将item全部插入
			fileTree = append(fileTree, v.Children...)
		}
		if len(fileTree) == 1 {
			break
		}
		fileTree = fileTree[1:]
	}
}

// Start ... 开始迁移
func Start() {
	doc, file, err := getData(uint32(3))
	if err != nil {
		panic(err)
	}

	fileTree, docTree, err := updateDataToProject(doc, file, uint32(3))
	if err != nil {
		panic(err)
	}

	// fmt.Println(docTree)
	updateDataToFileFolder(fileTree)
	updateDataToDocFolder(docTree)
}

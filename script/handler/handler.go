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
		log.Info("no record with this id",
			zap.String("id:", strconv.Itoa(int(id))))
		return nil, nil, err
	}

	file := &model.FileTreeNode{}
	json.Unmarshal([]byte(s.FileTree), file)

	doc := &model.FileTreeNode{}
	json.Unmarshal([]byte(s.DocTree), doc)

	return doc, file, nil
}

func writeString(s1 string, s2 string, itemType string) string {
	var result string
	if s1 == "" {
		result = fmt.Sprintf("%s-%s", s2, itemType)
	} else {
		result = fmt.Sprintf("%s,%s-%s", s1, s2, itemType)
	}
	return result
}

// writeItem ... 写入格式 file->0 , folder->1
func writeItem(tree *model.FileTreeNode) (string, []model.FileTreeNode) {
	var fileChildren string
	fileFolderItem := make([]model.FileTreeNode, 0) // 返回给下个函数 , fileChildren 中的文件夹部分
	for _, v := range tree.Children {
		if v.Folder {
			fileChildren = writeString(fileChildren, v.Id, "1")
			fileFolderItem = append(fileFolderItem, v)
		} else {
			fileChildren = writeString(fileChildren, v.Id, "0")
		}
	}

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

	log.Info("have got children",
		zap.String("children:", fileChildren+docChildren))

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
	var children string
	for _, v := range tree {
		if v.Folder {
			children = writeString(children, v.Id, "1")
		} else {
			children = writeString(children, v.Id, "0")
		}
	}

	return children
}

func insertDatabaseForFile(tree []model.FileTreeNode, id string) {
	// 将当前节点子节点都插入
	idInt, _ := strconv.Atoi(id)
	s, err := model.GetFolderForFileModel(uint32(idInt))
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

// 写入信息到 fileFolder
func updateDataToFileFolder(fileTree []model.FileTreeNode) {
	for len(fileTree) != 0 {
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
	for len(fileTree) != 0 {
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
	for i := 2; i <= 23; i++ {
		doc, file, err := getData(uint32(i))
		if err != nil {
			continue
		}

		fileTree, docTree, err := updateDataToProject(doc, file, uint32(i))
		if err != nil {
			log.Error("update record error",
				zap.String("cause:", err.Error()))
			panic(err)
		}

		// fmt.Println(docTree)
		updateDataToFileFolder(fileTree)
		updateDataToDocFolder(docTree)
	}
}

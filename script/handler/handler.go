package handler

import (
	"PROJECT_SCRIPT/constvar"
	"PROJECT_SCRIPT/log"
	"PROJECT_SCRIPT/model"
	"encoding/json"
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

// getData ... 获取一个 project 的数据，返回整个项目的 文件树 和 文档树
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
// project 专用，中间记录需要更新的 folder
func writeItem(tree *model.FileTreeNode) (string, []model.FileTreeNode) {
	var fileChildren string
	fileFolderItem := make([]model.FileTreeNode, 0) // 返回给下个函数 , fileChildren 中的文件夹部分
	for _, v := range tree.Children {
		if v.Folder {
			fileChildren = writeString(fileChildren, v.Id, "1") // 取出第一层，剩下的层数在 []children 字段里
			fileFolderItem = append(fileFolderItem, v)
			// 记录 folder，fatherId 默认为 0，不需要记录
		} else {
			// project 下的 children 默认 fatherId 为 0，不需要记录
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

// writeFolderItem ... 格式化字段
func writeFolderItem(tree []model.FileTreeNode, id uint32, code uint8, childCode uint8) (string, []model.HandleFatherIdSet) {
	var children string
	updateSet := make([]model.HandleFatherIdSet, 0)
	for _, v := range tree {
		if v.Folder {
			children = writeString(children, v.Id, "1")
			// 并入更新 folder father_id 的待处理集
			// 这里用 BFS，tree 的结构也不好改变，所以只能提取出来再更新一遍
			idInt, _ := strconv.Atoi(v.Id)
			updateSet = append(updateSet, model.HandleFatherIdSet{
				Type:     code,
				Id:       uint32(idInt),
				FatherId: id,
			})
		} else {
			children = writeString(children, v.Id, "0")
			// 并入更新 file father_id 的待处理集
			idInt, _ := strconv.Atoi(v.Id)
			updateSet = append(updateSet, model.HandleFatherIdSet{
				Type:     childCode,
				Id:       uint32(idInt),
				FatherId: id,
			})
		}
	}

	return children, updateSet
}

func insertDatabaseForFileFolder(tree []model.FileTreeNode, id string) []model.HandleFatherIdSet {
	// 将当前节点子节点都插入
	idInt, _ := strconv.Atoi(id)
	s, err := model.GetFolderForFileModel(uint32(idInt))
	if err != nil {
		fmt.Println(err)
	}

	children, updateSet := writeFolderItem(tree, uint32(idInt), constvar.FileFolderCode, constvar.FileCode)

	s.Children = children

	err = s.Update()
	if err != nil {
		fmt.Println(err)
	}

	return updateSet
}

// updateDataToFileFolder ... 写入信息到 fileFolder, BFS
func updateDataToFileFolder(fileTree []model.FileTreeNode) []model.HandleFatherIdSet {
	set := make([]model.HandleFatherIdSet, 0)
	for len(fileTree) != 0 {
		v := fileTree[0]
		if v.Children != nil {
			updateSet := insertDatabaseForFileFolder(v.Children, v.Id) // 插入结果，将item全部插入
			fileTree = append(fileTree, v.Children...)
			set = append(set, updateSet...)
		}
		if len(fileTree) == 1 { // 防止下面取出元素时越界
			break
		}
		fileTree = fileTree[1:]
	}

	return set
}

func insertDatabaseForDocFolder(tree []model.FileTreeNode, id string) []model.HandleFatherIdSet {
	// 将当前节点子节点都插入
	idInt, _ := strconv.Atoi(id)
	s, err := model.GetFolderForDocModel(uint32(idInt))
	if err != nil {
		fmt.Println(err)
	}

	children, updateSet := writeFolderItem(tree, uint32(idInt), constvar.DocFolderCode, constvar.DocCode)

	s.Children = children

	err = s.Update()
	if err != nil {
		fmt.Println(err)
	}

	return updateSet
}

// updateDataToDocFolder ... 写入信息到 docFolder, BFS 遍历
func updateDataToDocFolder(fileTree []model.FileTreeNode) []model.HandleFatherIdSet {
	set := make([]model.HandleFatherIdSet, 0)
	for len(fileTree) != 0 {
		v := fileTree[0]
		if v.Children != nil {
			updateSet := insertDatabaseForDocFolder(v.Children, v.Id) // 插入结果，将item全部插入
			fileTree = append(fileTree, v.Children...)
			set = append(set, updateSet...)
		}
		if len(fileTree) == 1 {
			break
		}
		fileTree = fileTree[1:]
	}
	return set
}

// updateFatherId ... 更新 father_id 集合
func updateFatherId(set []model.HandleFatherIdSet) error {
	var err error
	for _, v := range set {
		switch v.Type {
		case constvar.DocCode:
			err = model.UpdateFatherId(v.Id, v.FatherId, "docs")
		case constvar.DocFolderCode:
			err = model.UpdateFatherId(v.Id, v.FatherId, "foldersformds")
		case constvar.FileCode:
			err = model.UpdateFatherId(v.Id, v.FatherId, "files")
		case constvar.FileFolderCode:
			err = model.UpdateFatherId(v.Id, v.FatherId, "foldersforfiles")
		}

		if err != nil {
			return err
		}
	}
	return nil
}

// Start ... 开始迁移
// 新增字段 father_id
// 不需要更新 project 的下一层，因为默认就是 0
func Start() {
	// 获取最大 id 数
	maxId, err := model.GetProjectMaxId()
	if err != nil {
		panic(err)
	}

	for i := 0; i <= maxId; i++ {
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
		setFile := updateDataToFileFolder(fileTree)
		setDoc := updateDataToDocFolder(docTree)
		set := append(setFile, setDoc...)

		// 更新 fahter_id 集合
		err = updateFatherId(set)
		if err != nil {
			log.Error("update record error",
				zap.String("cause:", err.Error()))
			panic(err)
		}
	}

	log.Info("finsih")
}

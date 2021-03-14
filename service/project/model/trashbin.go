package model

import (
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"strconv"
	"strings"
)

// --Delete part

// DeleteProjectTrashbin ... 递归调用
func DeleteProjectTrashbin(id uint32) error {
	// 先查询子树
	project := &ProjectModel{}
	d := m.DB.Self.Where("id = ?", id).First(&project)
	if d.Error != nil {
		return d.Error
	}

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
			err = DeleteDocFolderTrashbin(uint32(childId))
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
			err = DeleteFileFolderTrashbin(uint32(childId))
			if err != nil {
				return err
			}
		}
	}

	return m.DB.Self.Unscoped().Delete(&project).Error
}

func DeleteDocTrashbin(id uint32) error {
	doc := &DocModel{}
	doc.ID = id
	return m.DB.Self.Unscoped().Delete(&doc).Error
}

func DeleteFileTrashbin(id uint32) error {
	file := &FileModel{}
	file.ID = id
	return m.DB.Self.Unscoped().Delete(&file).Error
}

func DeleteDocFolderTrashbin(id uint32) error {
	// 先查询子树
	docFolder := &FolderForDocModel{}
	d := m.DB.Self.Where("id = ?", id).First(&docFolder)
	if d.Error != nil {
		return d.Error
	}

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
			err = DeleteDocFolderTrashbin(uint32(childId))
			if err != nil {
				return err
			}
		}
	}

	return m.DB.Self.Unscoped().Delete(&docFolder).Error
}

func DeleteFileFolderTrashbin(id uint32) error {
	// 先查询子树
	fileFolder := &FolderForFileModel{}
	d := m.DB.Self.Where("id = ?", id).First(&fileFolder)
	if d.Error != nil {
		return d.Error
	}

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
			err = DeleteFileFolderTrashbin(uint32(childId))
			if err != nil {
				return err
			}
		}
	}

	return m.DB.Self.Unscoped().Delete(&fileFolder).Error
}

// --Remove part

func RemoveProjectTrashbin(id uint32) error {
	// 更改字段为 NULL
	d := m.DB.Self.Table("projects").Where("id = ?", id).Update("deleted_at", nil)
	if d.Error != nil {
		return d.Error
	}

	return nil
}

func RemoveDocTrashbin(id uint32) error {
	s := &DocModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	if d.Error != nil {
		return d.Error
	}

	// 更改字段
	s.Re = false

	if err := s.Update(); err != nil {
		return err
	}
	return nil
}

func RemoveFileTrashbin(id uint32) error {
	s := &FileModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	if d.Error != nil {
		return d.Error
	}

	// 更改字段
	s.Re = false

	if err := s.Update(); err != nil {
		return err
	}
	return nil
}

func RemoveDocFolderTrashbin(id uint32) error {
	s := &FolderForDocModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	if d.Error != nil {
		return d.Error
	}

	// 更改字段
	s.Re = false

	if err := s.Update(); err != nil {
		return err
	}
	return nil
}

func RemoveFileFolderTrashbin(id uint32) error {
	s := &FolderForFileModel{}
	d := m.DB.Self.Where("id = ?", id).First(&s)
	if d.Error != nil {
		return d.Error
	}

	// 更改字段
	s.Re = false

	if err := s.Update(); err != nil {
		return err
	}
	return nil
}

// --Get part

func GetProjectTrashbin() ([]*pb.Trashbin, error) {
	s := make([]*ProjectModel, 0)
	d := m.DB.Self.Unscoped().Scan(&s)

	if d.Error != nil {
		return nil, d.Error
	}

	var item = make([]*pb.Trashbin, 0)
	for _, v := range s {
		if v.DeletedAt.Valid {
			item = append(item, &pb.Trashbin{
				Id:   v.ID,
				Type: "0",
				Name: v.Name,
			})
		}
	}

	return item, nil
}

func GetDocTrashbin() ([]*pb.Trashbin, error) {
	s := make([]*DocModel, 0)
	d := m.DB.Self.Table("docs").Where("re = ?", 1).Scan(&s)

	if d.Error != nil {
		return nil, d.Error
	}

	var item = make([]*pb.Trashbin, 0)
	for _, v := range s {
		item = append(item, &pb.Trashbin{
			Id:   v.ID,
			Type: "1",
			Name: v.Name,
		})
	}

	return item, nil
}

func GetFileTrashbin() ([]*pb.Trashbin, error) {
	s := make([]*FileModel, 0)
	d := m.DB.Self.Table("file").Where("re = ?", 1).Scan(&s)

	if d.Error != nil {
		return nil, d.Error
	}

	var item = make([]*pb.Trashbin, 0)
	for _, v := range s {
		item = append(item, &pb.Trashbin{
			Id:   v.ID,
			Type: "2",
			Name: v.Name,
		})
	}

	return item, nil
}

func GetDocFolderTrashbin() ([]*pb.Trashbin, error) {
	s := make([]*FileModel, 0)
	d := m.DB.Self.Table("foldersformds").Where("re = ?", 1).Scan(&s)

	if d.Error != nil {
		return nil, d.Error
	}

	var item = make([]*pb.Trashbin, 0)
	for _, v := range s {
		item = append(item, &pb.Trashbin{
			Id:   v.ID,
			Type: "3",
			Name: v.Name,
		})
	}

	return item, nil
}

func GetFileFolderTrashbin() ([]*pb.Trashbin, error) {
	s := make([]*FolderForFileModel, 0)
	d := m.DB.Self.Table("foldersforfile").Where("re = ?", 1).Scan(&s)
	if d.Error != nil {
		return nil, d.Error
	}

	var item = make([]*pb.Trashbin, 0)
	for _, v := range s {
		item = append(item, &pb.Trashbin{
			Id:   v.ID,
			Type: "4",
			Name: v.Name,
		})
	}

	return item, nil
}

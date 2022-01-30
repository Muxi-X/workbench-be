package service

import (
	"context"
	"errors"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// UpdateFilePosition ... 移动文件
func (s *Service) UpdateFilePosition(ctx context.Context, req *pb.UpdateFilePositionRequest, res *pb.Response) error {
	// 判断 type 合法性
	fileType, fatherType, err := checkTypeIsValid(req.Type, req.FatherType)
	if err != nil {
		return e.BadRequestErr(errno.ErrInvalidFileType, err.Error())
	}
	// 检查文件是否被删
	isDeleted, err := checkFileIfDeleted(req.FileId, req.OldFatherId, req.FatherId, req.Type)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted {
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	// 用 fileType 和 fileId 找到目标文件
	fileItem, projectId, err := getFileItemByIdAndCode(req.FileId, fileType)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	if projectId != req.ProjectId {
		return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
	}

	isFatherProject := fatherType == constvar.ProjectCode

	isOldFatherProject := getOldFatherIsProject(fileItem, fileType)

	// 事务
	if err = model.UpdateFilePosition(fileItem, req.FatherId,
		req.OldFatherId, fileType, isFatherProject, isOldFatherProject,
		req.ChildrenPositionIndex); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

func getOldFatherIsProject(file interface{}, code uint8) bool {
	switch code {
	case constvar.DocCode:
		doc := file.(model.DocModel)
		if doc.FatherId != 0 {
			return false
		}
	case constvar.FileCode:
		file := file.(model.FileModel)
		if file.FatherId != 0 {
			return false
		}
	case constvar.DocFolderCode:
		file := file.(model.FolderForDocModel)
		if file.FatherId != 0 {
			return false
		}
	case constvar.FileFolderCode:
		file := file.(model.FolderForFileModel)
		if file.FatherId != 0 {
			return false
		}
	}

	return true
}

// 未使用
func getOldFather(file interface{}, code uint8) (interface{}, uint8, error) {
	var fileItem interface{}
	var oldFatherType uint8
	var err error
	switch code {
	case constvar.DocCode:
		doc := file.(model.DocModel)
		if doc.FatherId != 0 {
			oldFatherType = constvar.DocFolderCode
			fileItem, err = model.GetFolderForDocModel(doc.FatherId)
		} else {
			oldFatherType = constvar.ProjectCode
			fileItem, err = model.GetProject(doc.FatherId)
		}
	case constvar.FileCode:
		file := file.(model.FileModel)
		if file.FatherId != 0 {
			oldFatherType = constvar.DocFolderCode
			fileItem, err = model.GetFolderForDocModel(file.FatherId)
		} else {
			oldFatherType = constvar.ProjectCode
			fileItem, err = model.GetProject(file.FatherId)
		}
	default:
		err = errors.New("wrong type_id")
	}
	return fileItem, oldFatherType, err
}

func getFileItemByIdAndCode(id uint32, code uint8) (interface{}, uint32, error) {
	var fileItem interface{}
	var err error
	var projectId uint32
	switch code {
	case constvar.ProjectCode:
		fileItem, err = model.GetProject(id)
		projectId = fileItem.(model.ProjectModel).ID
	case constvar.DocCode:
		fileItem, err = model.GetDoc(id)
		projectId = fileItem.(model.DocModel).ProjectID
	case constvar.FileCode:
		fileItem, err = model.GetFile(id)
		projectId = fileItem.(model.FileModel).ProjectID
	case constvar.DocFolderCode:
		fileItem, err = model.GetFolderForDocModel(id)
		projectId = fileItem.(model.FolderForDocModel).ProjectID
	case constvar.FileFolderCode:
		fileItem, err = model.GetFolderForFileModel(id)
		projectId = fileItem.(model.FolderForFileModel).ProjectID
	default:
		err = errors.New("wrong type code")
	}

	return fileItem, projectId, err
}

func checkTypeIsValid(reqFileType, reqFatherType uint32) (uint8, uint8, error) {
	fileType, err := checkFileTypeIsValid(reqFileType)
	if err != nil {
		return uint8(0), uint8(0), err
	}
	if reqFatherType != 0 && reqFatherType != 1 {
		return uint8(0), uint8(0), errors.New("(father type be 0 -> project file, 1 -> other)")
	}

	var fatherType uint8
	if reqFatherType == 1 { // 非根目录
		fatherType = fileType + 2
	}
	return fileType, fatherType, nil
}

func checkFileTypeIsValid(reqType uint32) (uint8, error) {
	fileType := uint8(reqType)
	if fileType != constvar.DocCode || fileType != constvar.FileCode {
		return uint8(0), errors.New("file type must be 1 or 2")
	}
	return fileType, nil
}

func checkFileIfDeleted(fileId, oldFatherId, fatherId, fileType uint32) (bool, error) {
	// 判断新的 父 id 和自身 id 有无删除
	isDeleted, err := model.AdjustSelfAndFatherIfExist(fileId, oldFatherId, uint8(fileType), uint8(fileType+2))
	if err != nil {
		return false, err
	}
	if isDeleted {
		return isDeleted, nil
	}
	isDeleted, err = model.AdjustSelfIfExist(fatherId, uint8(fileType+2))
	if err != nil {
		return false, err
	}
	if isDeleted {
		return isDeleted, nil
	}

	return false, nil
}

package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
	"strconv"
	"strings"
	"time"
)

// DeleteFolder ... 删除folder
// 寻找子文件同步 redis 修改文件树 插入回收站
func (s *Service) DeleteFolder(ctx context.Context, req *pb.DeleteRequest, res *pb.Response) error {
	var item *model.FolderModel
	var err error
	if uint8(req.TypeId) == constvar.DocCode {
		item, err = model.GetFolderForDocModel(req.Id)
	} else if uint8(req.TypeId) == constvar.FileFolderCode {
		item, err = model.GetFolderForFileModel(req.Id)
	} else {
		return e.BadRequestErr(errno.ErrBind, "wrong type_id")
	}
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 权限判定
	if item.CreatorID != req.UserId {
		if req.Role <= constvar.Normal {
			return e.BadRequestErr(errno.ErrPermissionDenied, "")
		}
	}

	if item.ProjectID != req.ProjectId {
		return e.ServerErr(errno.ErrPermissionDenied, "project_id mismatch")
	}

	// 获取 fatherId
	isFatherProject := false
	var fatherId uint32
	if item.FatherId == 0 { // fatherId 为 0 则将 fatherId 设为 projectId
		isFatherProject = true
		fatherId = item.ProjectID
	} else {
		fatherId = item.FatherId
	}

	trashbin := &model.TrashbinModel{
		FileId:     req.Id,
		FileType:   uint8(req.TypeId),
		Name:       item.Name,
		DeleteTime: time.Now().Format("2006-01-02 15:04:05"),
		CreateTime: item.CreateTime,
		ProjectID:  req.ProjectId,
	}

	// 事务
	if err = trashbin.DeleteChildren(fatherId, isFatherProject); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	items, err := GetInfoByFolderId(item.ID, uint8(req.TypeId))
	for _, item := range items {
		err = DeleteAttentionsFromAttentionService(item.ID, req.TypeId, req.UserId)
		if err != nil {
			return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
		}
	}

	return nil
}

func GetInfoByFolderId(folderId uint32, typeId uint8) ([]*model.FileInfo, error) {
	var ids []uint32
	var getFolderModel func(uint32) (*model.FolderModel, error)
	var getInfoByIds func([]uint32) ([]*model.FileInfo, error)

	if typeId == constvar.DocFolderCode {
		getFolderModel = model.GetFolderForDocModel
		getInfoByIds = model.GetDocInfoByIds
	} else if typeId == constvar.FileFolderCode {
		getFolderModel = model.GetFolderForFileModel
		getInfoByIds = model.GetFileInfoByIds
	}

	var f func()
	f = func() {
		item, err := getFolderModel(folderId)

		if err != nil || len(item.Children) == 0 {
			return
		}

		raw := strings.Split(item.Children, ",")
		for _, v := range raw {
			r := strings.Split(v, "-")
			id, _ := strconv.Atoi(r[0])
			if r[1] == "0" {
				ids = append(ids, uint32(id))
			} else {
				folderId = uint32(id)
				f()
			}
		}
	}

	f()
	infos, err := getInfoByIds(ids)
	if err != nil {
		return infos, e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return infos, nil
}

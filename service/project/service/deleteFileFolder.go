package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
	"strconv"
	"strings"
	"time"
)

// DeleteFileFolder ... 删除文件夹
func (s *Service) DeleteFileFolder(ctx context.Context, req *pb.DeleteRequest, res *pb.Response) error {
	item, err := model.GetFolderForFileModel(req.Id)
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
	if item.FatherId == 0 { // fatherId 为 0 则是 project
		isFatherProject = true
		fatherId = item.ProjectID
	} else {
		fatherId = item.FatherId
	}

	trashbin := &model.TrashbinModel{
		FileId:     req.Id,
		FileType:   constvar.DocFolderCode,
		Name:       item.Name,
		DeleteTime: time.Now().Format("2006-01-02 15:04:05"),
		CreateTime: item.CreateTime,
		ProjectID:  req.ProjectId,
	}

	// 事务
	if err := model.DeleteFileFolder(m.DB.Self, trashbin, fatherId, isFatherProject); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 获取文件夹的file列表
	files, err := GetFilesByChildren(item.Children)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	for _, file := range files {
		err = DeleteAttentionsFromAttentionService(file.ID, uint32(constvar.FileCode), req.UserId)
		if err != nil {
			return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
		}
	}
	return nil
}

func GetFilesByChildren(children string) ([]*model.FileDetail, error) {
	if len(children) == 0 {
		return nil, nil
	}
	var files []*model.FileDetail
	raw := strings.Split(children, ",")
	for _, v := range raw {
		r := strings.Split(v, "-")
		id, _ := strconv.Atoi(r[0])
		if r[1] == "0" {
			file, err := model.GetFileDetail(uint32(id))
			if err != nil {
				return files, e.ServerErr(errno.ErrDatabase, err.Error())
			}
			files = append(files, file)
		}
	}
	return files, nil
}

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

// DeleteDocFolder ... 删除文档
// 寻找子文件同步 redis 修改文件树 插入回收站
func (s *Service) DeleteDocFolder(ctx context.Context, req *pb.DeleteRequest, res *pb.Response) error {
	item, err := model.GetFolderForDocModel(req.Id)
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
	if item.FatherId == 0 { // fatherId 为 0 则将 fatherId 设是 projectId
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
	err = model.DeleteDocFolder(m.DB.Self, trashbin, fatherId, isFatherProject)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 获取文档夹的doc列表
	docs, err := GetDocsByChildren(item.Children)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	for _, doc := range docs {
		err = DeleteAttentionsFromAttentionService(doc.ID, uint32(constvar.DocCode), req.UserId)
		if err != nil {
			return e.ServerErr(errno.ErrGetDataFromRPC, err.Error())
		}
	}
	return nil
}

func GetDocsByChildren(children string) ([]*model.DocDetail, error) {
	if len(children) == 0 {
		return nil, nil
	}
	var docs []*model.DocDetail
	raw := strings.Split(children, ",")
	for _, v := range raw {
		r := strings.Split(v, "-")
		id, _ := strconv.Atoi(r[0])
		if r[1] == "0" {
			doc, err := model.GetDocDetail(uint32(id))
			if err != nil {
				return docs, e.ServerErr(errno.ErrDatabase, err.Error())
			}
			docs = append(docs, doc)
		}
	}
	return docs, nil
}

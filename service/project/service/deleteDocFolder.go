package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
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
	}

	// 事务
	err = model.DeleteDocFolder(m.DB.Self, trashbin, fatherId, isFatherProject)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 删除文档folder后删除对应attentions TODO 测试

	// 获取文档夹的信息
	list, err := model.GetFolderForDocInfoByIds([]uint32{req.Id})
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	for _, doc := range list {
		err = DeleteAttentionsFromAttentionService(doc.ID, uint32(constvar.DocCode), req.UserId)
		if err != nil {
			return e.ServerErr(errno.ErrDatabase, err.Error())
		}
	}
	return nil
}

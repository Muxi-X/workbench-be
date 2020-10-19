package status

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommentList ... 获取动态详情
func CommentList(c *gin.Context) {
	log.Info("Status commentList function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从 Query Param 中获取 lastid 和 limit
	var limit int
	var lastId int
	var sid int
	var page int
	var err error

	sid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	lastId, err = strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 page
	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 listcomment 请求并发送
	listComReq := &pbs.CommentListRequest{
		StatusId: uint32(sid),
		Offset:   uint32(page * limit),
		Limit:    uint32(limit),
		Lastid:   uint32(lastId),
	}

	listComResp, err := service.StatusClient.ListComment(context.Background(), listComReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	resp := CommentListResponse{
		Count: listComResp.Count,
	}

	for i := 0; i < len(listComResp.List); i++ {
		resp.CommentList = append(resp.CommentList, Comment{
			Cid:      listComResp.List[i].Id,
			Uid:      listComResp.List[i].UserId,
			Username: listComResp.List[i].UserName,
			Avatar:   listComResp.List[i].Avatar,
			Time:     listComResp.List[i].Time,
			Content:  listComResp.List[i].Content,
		})
	}

	SendResponse(c, nil, resp)
}

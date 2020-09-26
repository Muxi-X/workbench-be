package status

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
)

// 需要调用 get 和 listcomment
// 不需要获取 userid
func Get(c *gin.Context) {
	log.Info("Status get function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从 Query Param 中获取 lastid 和 limit
	var limit int
	var lastId int
	var sid int
	var page int
	var err error

	sid, err = strconv.Atoi(c.Param("sid"))
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

	// 构造 get status 请求并发送
	getReq := &pbs.GetRequest{
		Id: uint32(sid),
	}

	getResp, err2 := service.StatusClient.Get(context.Background(), getReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error(), GetLine())
		return
	}

	// 构造 listcomment 请求并发送
	listComReq := &pbs.CommentListRequest{
		StatusId: uint32(sid),
		Offset:   uint32(page),
		Limit:    uint32(limit),
		Lastid:   uint32(lastId),
	}

	listComResp, err3 := service.StatusClient.ListComment(context.Background(), listComReq)
	if err3 != nil {
		SendError(c, errno.InternalServerError, nil, err3.Error(), GetLine())
		return
	}

	// 构造返回 response
	resp := getResponse{
		Sid:      uint32(sid),
		Title:    getResp.Status.Title,
		Content:  getResp.Status.Content,
		UserId:   getResp.Status.UserId,
		Time:     getResp.Status.Time,
		Avatar:   getResp.Status.Avatar,
		Username: getResp.Status.UserName,
		Count:    listComResp.Count,
	}

	for i := 0; i < len(listComResp.List); i++ {
		resp.Commentlist = append(resp.Commentlist, comment{
			Cid:      listComResp.List[i].Id,
			Uid:      listComResp.List[i].UserId,
			Username: listComResp.List[i].UserName,
			Avatar:   listComResp.List[i].Avatar,
			Time:     listComResp.List[i].Time,
			Content:  listComResp.List[i].Content,
		})
	}

	// 返回 response
	SendResponse(c, nil, resp)
}

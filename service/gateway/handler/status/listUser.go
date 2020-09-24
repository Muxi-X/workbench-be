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

// 只用调用一次 list  lastid limit 要从 query param 获取 要获取uid
// 不需要从 token 获取 userid
func ListUser(c *gin.Context) {
	log.Info("Status list user function call",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 gid 和 limit lastid
	var limit int
	var lastid int
	var uid int
	var page int
	var err error

	uid, err = strconv.Atoi(c.Param("uid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		log.Fatal("status listUser, get param:uid fatal",
			zap.String("reason", err.Error()))
		return
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		log.Fatal("status listUser, get param:limit fatal",
			zap.String("reason", err.Error()))
		return
	}

	lastid, err = strconv.Atoi(c.DefaultQuery("lastid", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		log.Fatal("status listUser, get param:lastid fatal",
			zap.String("reason", err.Error()))
		return
	}

	// 获取 page
	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		log.Fatal("status listUser, get param:page fatal",
			zap.String("reason", err.Error()))
		return
	}

	// 构造 list 请求
	listReq := &pbs.ListRequest{
		Lastid: uint32(lastid),
		Offset: uint32(page),
		Limit:  uint32(limit),
		Group:  0,
		Uid:    uint32(uid),
	}

	listResp, err2 := service.StatusClient.List(context.Background(), listReq)
	if err2 != nil {
		SendBadRequest(c, errno.ErrBind, nil, err2.Error())
		log.Fatal("status listUser, get response from status server fatal",
			zap.String("reason", err2.Error()))
	}

	// 构造返回 response
	var resp listResponse
	for i := 0; i < len(listResp.List); i++ {
		resp.Status = append(resp.Status, status{
			Id:       listResp.List[i].Id,
			Title:    listResp.List[i].Title,
			Content:  listResp.List[i].Content,
			UserId:   listResp.List[i].UserId,
			Time:     listResp.List[i].Time,
			Avatar:   listResp.List[i].Avatar,
			Username: listResp.List[i].UserName,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}

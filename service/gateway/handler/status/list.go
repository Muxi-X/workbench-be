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

// List ... 获取进度列表
func List(c *gin.Context) {
	log.Info("Status list function call",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var teamID uint32
	var team int
	var uid int
	var group int
	var limit int
	var lastID int
	var page int
	var err error

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	lastID, err = strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 gid
	group, err = strconv.Atoi(c.DefaultQuery("group", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 uid
	uid, err = strconv.Atoi(c.DefaultQuery("uid", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取是否筛选 team
	team, err = strconv.Atoi(c.DefaultQuery("team", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if team == 1 {
		// 还要获取用户 teamid
		teamID = c.MustGet("TeamID").(uint32)
	}

	userID := c.MustGet("UserID").(uint32)

	// 构造 list 请求
	listReq := &pbs.ListRequest{
		Lastid: uint32(lastID),
		Offset: uint32(page * limit),
		Limit:  uint32(limit),
		Group:  uint32(group), // 这里传 URL 里面获取的 group 参数，DefaultQuery("group", "0")
		Uid:    uint32(uid),   // 这里传 URL 里面获取的 group 参数，DefaultQuery("uid", "0")
		Team:   teamID,
		UserId: userID,
	}

	listResp, err := service.StatusClient.List(context.Background(), listReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var resp ListResponse
	for i := 0; i < len(listResp.List); i++ {
		resp.Status = append(resp.Status, Status{
			Id:       listResp.List[i].Id,
			Title:    listResp.List[i].Title,
			Content:  listResp.List[i].Content,
			Time:     listResp.List[i].Time,
			Avatar:   listResp.List[i].Avatar,
			Username: listResp.List[i].UserName,
			IfLike:   listResp.List[i].IfLike,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}

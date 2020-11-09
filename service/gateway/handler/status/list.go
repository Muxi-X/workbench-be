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
// @Summary list status api
// @Description 拉取进度列表
// @Tags status
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "status_id"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Param page query int false "page"
// @Param group query int false "group"
// @Param uid query int false "uid"
// @Param team query int false "team"
// @Security ApiKeyAuth
// @Success 200 {object} ListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /status [get]
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
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	lastID, err = strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 gid
	group, err = strconv.Atoi(c.DefaultQuery("group", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 uid
	uid, err = strconv.Atoi(c.DefaultQuery("uid", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取是否筛选 team
	team, err = strconv.Atoi(c.DefaultQuery("team", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	if team == 1 {
		// 还要获取用户 teamid
		teamID = c.MustGet("TeamID").(uint32)
	}

	userID := c.MustGet("UserID").(uint32)

	// 构造 list 请求
	listReq := &pbs.ListRequest{
		LastId: uint32(lastID),
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
	for _, item := range listResp.List {
		resp.Status = append(resp.Status, Status{
			Id:           item.Id,
			Title:        item.Title,
			Content:      item.Content,
			Time:         item.Time,
			CommentCount: item.Comment,
			LikeCount:    item.Like,
			Avatar:       item.Avatar,
			Username:     item.UserName,
			Liked:        item.Liked,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}

package project

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// 需要调用 list
// 需要从 token 获取 userid
func GetProjectList(c *gin.Context) {
	log.Info("project getProjectList function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 lastid page limit pagination
	var limit int
	var lastId int
	var page int
	var pagination int
	var err error

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	lastId, err = strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	pagination, err = strconv.Atoi(c.DefaultQuery("pagination", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	var pageBool bool
	if pagination == 1 {
		pageBool = true
	}

	// 获取 userid
	id := c.MustGet("userID").(uint32)

	// 构造请求
	getProListReq := &pbp.GetProjectListRequest{
		UserId:     id,
		Lastid:     uint32(lastId),
		Offset:     uint32(page),
		Limit:      uint32(limit * page),
		Pagination: pageBool,
	}

	getProListResp, err := service.ProjectClient.GetProjectList(context.Background(), getProListReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造response
	var resp getProjectListResponse
	for i := 0; i < len(getProListResp.List); i++ {
		resp.Projectlist = append(resp.Projectlist, projectListItem{
			Id:   getProListResp.List[i].Id,
			Name: getProListResp.List[i].Name,
			Logo: getProListResp.List[i].Logo,
		})
	}

	SendResponse(c, nil, resp)

}

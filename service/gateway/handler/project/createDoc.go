package project

import (
	"context"

	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateDoc creates a new doc
// 更新：前端不用传 fatherType 根据 fatherId 是否为 0 来判断
// @Summary create a new doc api
// @Description 通过父节点 id，和插入节点位置确定新建文档的位置。
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body CreateDocRequest true "create_doc_request"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/doc [post]
func CreateDoc(c *gin.Context) {
	log.Info("project createDoc function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 先建立 doc ，再修改 docChildren
	// 获得请求
	var req CreateDocRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	userID := c.MustGet("userID").(uint32)
	teamID := c.MustGet("teamID").(uint32)
	// 获取 projectID
	projectID := c.MustGet("projectID").(uint32)

	createDocReq := &pbp.CreateDocRequest{
		Title:                 req.Title,
		Content:               req.Content,
		ProjectId:             projectID,
		UserId:                userID,
		TeamId:                teamID,
		FatherId:              req.FatherID,
		ChildrenPositionIndex: req.ChildrenPositionIndex,
	}

	resp, err := service.ProjectClient.CreateDoc(context.Background(), createDocReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	/* --- 新增 feed --- */

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "创建",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        3,
			Id:          resp.Id,
			Name:        req.DocName,
			ProjectId:   projectID,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err = service.FeedClient.Push(context.Background(), pushReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

package attention

import (
	"context"
	"strconv"

	pb "muxi-workbench-attention/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// List ... lists attentions.
// @Summary get attention list api
// @Description 获取此用户的关注list
// @Tags attention
// @Accept application/json
// @Produce application/json
// @Param id path int true "user_id"
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Success 200 {object} AttentionListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /attention/list/{id} [get]
func List(c *gin.Context) {
	log.Info("attention List function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 limit
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 lastId
	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 userId
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	listReq := &pb.ListRequest{
		LastId: uint32(lastId),
		Limit:  uint32(limit),
		UserId: uint32(userId),
	}

	listResp, err := service.AttentionClient.List(context.Background(), listReq)
	if err != nil {
		SendError(c, errno.ErrAttentionList, nil, err.Error(), GetLine())
		return
	}

	// var list []*AttentionItem
	// for _, item := range listResp.List {
	// 	list = append(list, &AttentionItem{
	// 		Id:   item.Id,
	// 		Date: item.Date,
	// 		Time: item.Time,
	// 		User: &AttentionUser{
	// 			Name:      item.User.Name,
	// 			Id:        item.User.Id,
	// 		},
	// 		Doc:  &Doc{
	// 			Id:          item.Doc.Id,
	// 			Name:        item.Doc.Name,
	// 			DocCreator: &AttentionUser{
	// 				Name: item.Doc.DocCreator.Name,
	// 				Id:   item.Doc.DocCreator.Id,
	// 			},
	// 			ProjectId:   item.Doc.ProjectId,
	// 			ProjectName: item.Doc.ProjectName,
	// 		},
	// 	})
	// }
	//
	// SendResponse(c, nil, AttentionListResponse{
	// 	Count: listResp.Count,
	// 	List:  list,
	// })

	SendResponse(c, nil, listResp)
}

package status

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
)

// 需要调用 status create 和 feed push
// userid 从 token 获取
func CreateComment(c *gin.Context) {
	log.Info("Status createComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 sid
	var sid int
	var err error

	sid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取请求
	var req createCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	id := c.MustGet("userID").(uint32)

	_, err2 := service.StatusClient.CreateComment(context.Background(), &pbs.CreateCommentRequest{
		UserId:   id,
		StatusId: uint32(sid),
		Content:  req.Content,
	})
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 要通过 sid 获取 status 的 title
	getReq := &pbs.GetRequest{
		Id: uint32(sid),
	}

	getResp, err4 := service.StatusClient.Get(context.Background(), getReq)
	if err4 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error(), GetLine())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "评论",
		UserId: id,
		Source: &pbf.Source{
			Kind:        6,
			Id:          uint32(sid), // 暂时从前端获取
			Name:        getResp.Status.Title,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err3 := service.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		SendError(c, errno.InternalServerError, nil, err3.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

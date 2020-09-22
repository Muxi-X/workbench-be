package project

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
<<<<<<< HEAD
	"muxi-workbench-gateway/pkg/token"
=======
>>>>>>> master
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// 调用一次 update 和一次 feed push
<<<<<<< HEAD
// 需要从 token 获取 userid
=======
>>>>>>> master
func UpdateMembers(c *gin.Context) {
	log.Info("Project member update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 pid
	var pid int
	var err error

	pid, err = strconv.Atoi(c.Param("pid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取请求
	var req updateMemberRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

<<<<<<< HEAD
	// 获取 userid
	raw, ifexists := c.Get("context")
	if !ifexists {
		SendBadRequest(c, errno.ErrTokenInvalid, nil, "Context not exists")
	}
	ctx, ok := raw.(*token.Context)
	if !ok {
		SendError(c, errno.ErrValidation, nil, "Context assign failed")
	}

=======
>>>>>>> master
	// 构造请求
	// 这里 list 应该是 uint32 表示 uid
	updateMemReq := &pbp.UpdateMemberRequest{
		Id: uint32(pid),
	}
	for i := 0; i < len(req.Userlist); i++ {
		updateMemReq.List = append(updateMemReq.List, req.Userlist[i])
	}

	// 发送请求
	_, err2 := service.ProjectClient.UpdateMembers(context.Background(), updateMemReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "编辑",
<<<<<<< HEAD
		UserId: uint32(ctx.ID),
=======
		UserId: req.UserId,
>>>>>>> master
		Source: &pbf.Source{
			Kind:        2,
			Id:          0, // 暂时从前端获取
			Name:        "",
			ProjectId:   uint32(pid),
			ProjectName: req.ProjectName,
		},
	}

	// 向 feed 发送请求
	_, err3 := service.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		SendError(c, errno.InternalServerError, nil, err3.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}

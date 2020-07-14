package status

import (
	"context"
	//"fmt"
	"log"

	//tracer "muxi-workbench-status-client/tracer"
	pbf "muxi-workbench-feed/proto"
	pbs "muxi-workbench-status/proto"
	handler "muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

// 需要调用 feed push 和 status update
func Update(c *gin.Context) {
	log.Info("Status update function call.")

	// 获取 sid 和请求
	var sid int
	sid, err = strconv.Atoi(c.Param("sid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取请求
	var req updateRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造 updateReq 并发送请求
	updateReq := &pbs.UpdateRequest{
		Id:      sid,
		Title:   req.Title,
		Content: req.Content,
		UserId:  req.UserId,
	}

	_, err2 := StatusClient.Update(context.Background(), updateReq)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "编辑",
		UserId: req.UserId,
		Source: &pbf.Source{
			Kind:        6,
			Id:          sid, // 暂时从前端获取
			Name:        req.Title,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err3 := feed.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		log.Fatalf("Could not greet: %v", err3)
		SendError(c, errno.InternalServerError, nil, err3.Error())
		return
	}

	// 返回结果
	SendResponse(c, errno.OK, nil)
}

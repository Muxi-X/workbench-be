package status

import (
	"context"
	//"fmt"
	"log"

	//tracer "muxi-workbench-status-client/tracer"
	pbs "muxi-workbench-status/proto"
	handler "muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

// 需要调用 status like
func Like(c *gin.Context) {
	log.Info("Status like function call")

	// 获取sid
	var sid int
	var err error

	sid, err = strconv.Atoi(c.Param("sid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取请求
	var req pb.LikeRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 调用 like 请求
	_, err2 := StatusClient.Like(context.Background(), &pbs.LikeRequest{
		Id:     sid,
		UserId: req.UserId,
	})
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		endError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}

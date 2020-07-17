package handler

import (
	"context"
	//"fmt"
	"log"

	//tracer "muxi-workbench-project-client/tracer"
	pbp "muxi-workbench-project/proto"
	handler "muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

// 只调用一次 getdocdetail
func GetDocDetail(c *gin.Context) {
	log.Info("Doc detail get function call.")

	// 获取 did
	var did int
	var err error

	did, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	getDocDetailResp, err2 := ProjectClient.GetDocDetail(context.Background(), &pbp.GetRequest{
		Id: did,
	})
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	//构造返回结果
	resp := getDocDetailResponse{
		Id:           getDocDetailResp.Id,
		Title:        getDocDetailResp.Title,
		Content:      getDocDetailResp.Content,
		Creator:      getDocDetailResp.Creator,
		Createtime:   getDocDetailResp.CreateTime,
		Lasteditor:   getDocDetailResp.LastEditor,
		Lastedittime: getDocDetailResp.LastEditTime,
	}

	SendResponse(c, nil, resp)
}

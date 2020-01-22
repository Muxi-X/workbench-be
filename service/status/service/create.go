package service

import (
	"context"
	errno "muxi-workbench-status/errno"
	"muxi-workbench-status/model"
	pb "muxi-workbench-status/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// Create ... 创建动态
func (s *StatusService) Create(ctx context.Context, req *pb.CreateRequest, res *pb.Response) error {
	// get tracing info from context
	// md, ok := metadata.FromContext(ctx)
	// if !ok {
	// 	md = make(map[string]string)
	// }
	// var sp opentracing.Span
	//wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// // create new span and bind with context
	// sp = opentracing.StartSpan("Hello", opentracing.ChildOf(wireContext))
	// // record request
	// sp.SetTag("req", req)
	// defer func() {
	// 	// record response
	// 	sp.SetTag("res", res)
	// 	// before function return stop span, cuz span will counted how much time of this function spent
	// 	sp.Finish()
	// }()
	// if sc, ok := wireContext.(jaeger.SpanContext); ok {
	// 	sc.TraceID()
	// }
	// md, ok := metadata.FromContext(ctx)
	// fmt.Println(md["uber-trace-id"], ok)
	t := time.Now()

	status := model.StatusModel{
		UserID:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
		Time:    t.Format("2006-01-02 15:04:05"),
	}

	if err := status.Create(); err != nil {
		return e.ServerErr(errno.InternalServerError, err.Error())
	}

	return nil
}

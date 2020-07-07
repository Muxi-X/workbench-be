package handler

/*import (
     "context"
    "fmt"
    "github.com/opentracing/opentracing-go"
    "log"

    //tracer "muxi-workbench-status-client/tracer"

    //opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"

    pb "muxi-workbench-status/proto"

    handler "muxi-workbench/pkg/handler"

    micro "github.com/micro/go-micro"
)

var StatusService micro.Service

func statusInit(StatusService micro.Service){
    StatusService = micro.NewService(micro.Name("workbench.cli.status"),
        micro.WrapClient(
            opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
        ),
        micro.WrapCall(handler.ClientErrorHandlerWrapper()))
    StatusService.Init()
}


func StatusGet(c *gin.Context) {
    client := pb.NewStatusServiceClient("workbench.service.status", StatusService.Client())

    resp, err := client.Get(context.Background(), &pb.GetRequest{
      Id: 1,
    })

    if err != nil {
        log.Fatalf("Could not greet: %v", err)
    }
}

func StatusList(c *gin.Context){
    client := pb.NewStatusServiceClient("workbench.service.status", service.Client())

    resp, err := client.List(context.Background(), &pb.ListRequest{
      Offset: 0,
      Limit:  20,
      Lastid: 162,
      Group:  3,
      Uid:    0,
    })

     if err != nil {
        log.Fatalf("Could not greet: %v", err)
    }
}

func StatusCreate(c *gin.Context){
    client := pb.NewStatusServiceClient("workbench.service.status", service.Client())

    req := &pb.CreateRequest{
        UserId:  0,
        Title:   "哈哈哈哈😁",
        Content: "后废物废物分为",
    }

    _, err = client.Create(context.Background(), req)

    if err != nil {
        log.Fatalf("Could not greet: %v",err)
    }
}
//update
func StatusUpdate(c *gin.Context){
    client := pb.NewStatusServiceClient("workbench.service.status", service.Client())

}
//Delete
func StatusDelete(c *gin.Context){
    client := pb.NewStatusServiceClient("workbench.service.status", service.Client())

}
//Like
func StatusLike(c *gin.Context){
    client := pb.NewStatusServiceClient("workbench.service.status", service.Client())
}

func StatusCreateComment(c *gin.Context){
    client := pb.NewStatusServiceClient("workbench.service.status", service.Client())

    req := &pb.CreateCommentRequest{
        UserId:  0,
        StatusId: 3488,
        Content: "后废物废物分为",
    }

    _, err = client.CreateComment(context.Background(), req)

    if err != nil{
        log.Fatalf("Could not greet: %v",err)
    }
}

func StatusListComment(c *gin.Context){
    client := pb.NewStatusServiceClient("workbench.service.status", service.Client())

    req := &pb.CommentListRequest{
        StatusId: 3488,
        Offset: 0,
        Limit: 20,
        Lastid: 0,
    }

    resp, err := client.ListComment(context.Background(), req)

    if err != nil{
        log.Fatalf("Could not greet: %v",err)
    }

    fmt.Println(resp.List, resp.Count)
}*/

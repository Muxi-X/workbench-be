package handler

import(
    "context"
    //"fmt"
    "log"

    //"muxi-workbench-feed-client/tracer"
    pb "muxi-workbench-feed/proto"
    "muxi-workbench/pkg/handler"

    micro "github.com/micro/go-micro"
    opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
    "github.com/opentracing/opentracing-go"

    "github.com/gin-gonic/gin"
)

var FeedService micro.Service
var FeedClient pb.FeedServiceClient

func FeedInit(FeedService micro.Service,FeedClient pb.FeedServiceClient){
    FeedService=micro.NewService(micro.Name("workbench.cli.feed"),
        micro.WrapClient(
            opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
        ),
        micro.WrapCall(handler.ClientErrorHandlerWrapper()),
    )

    FeedService.Init()

    FeedClient = pb.NewFeedServiceClient("workbench.service.feed", FeedService.Client())

}

//Feed的List接口
func FeedList(c *gin.Context){
    //获取前端request数据
    var req pb.ListRequest
    if err := c.BindJSON(&req); err != nil{
        c.JSON(400, gin.H{
            "message":"Wrong",
        })
        return
    }

    /*req := &pb.ListRequest{
        LastId: 68,
        Limit:  5,
        Role:   3,
        UserId: 53,
        Filter: &pb.Filter{
            UserId:  0,
            GroupId: 3,
        },
    }*/

    resp, err2 := FeedClient.List(context.Background(), &req)
    if err2 != nil {
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }
    //fmt.Println(resp)
    c.JSON(200,resp)

}

//Feed的Push接口
func FeedPush(c *gin.Context){
    var req pb.PushRequest
    if err := c.BindJSON(&req);err != nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    /*addReq := &pb.PushRequest{
      Action: "创建",
      UserId: 2333,
      Source: &pb.Source{
          Kind:        6,
          Id:          6666,   // status id
          Name:        "测试数据", // 进度标题
          ProjectId:   0,      // 固定数据
          ProjectName: "",     // 固定数据
      },
    }*/
    resp, err2 := FeedClient.Push(context.Background(), &req)
    if err2 != nil {
      log.Fatalf("Could not greet: %v",err2)
      c.JSON(500,gin.H{
          "message":"wrong",
      })
      return
    }
    //fmt.Println(addResp)
    c.JSON(200,resp)
}

package handler

import (
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
var StatusClient pb.StatusServiceClient

func StatusInit(StatusService micro.Service,StatusClient pb.StatusServiceClient){
    StatusService = micro.NewService(micro.Name("workbench.cli.status"),
        micro.WrapClient(
            opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
        ),
        micro.WrapCall(handler.ClientErrorHandlerWrapper()))
    StatusService.Init()

    Statusclient = pb.NewStatusServiceClient("workbench.service.status", StatusService.Client())

}


func StatusGet(c *gin.Context) {
    var getId int
    if err:=c.BindJSON(&Id);err != nil{
        c.JSON(400,gin.H{
            "message":"wrong",
            })
    }
    resp, err := StatusClient.Get(context.Background(), &pb.GetRequest{
      Id: getId,
    })

    if err != nil {
        log.Fatalf("Could not greet: %v", err)
    }
}

func StatusList(c *gin.Context){
    var listRequest pb.ListRequest
    if err:=c.BindJSON(&listRequest);err != nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp, err := Status.Client.List(context.Background(), listRequest)
    if err != nil {
        log.Fatalf("Could not greet: %v", err)
    }
}

func StatusCreate(c *gin.Context){
    var req pb.CreateRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    /*req := &pb.CreateRequest{
        UserId:  0,
        Title:   "哈哈哈哈😁",
        Content: "后废物废物分为",
    }*/

    _, err = StatusClient.Create(context.Background(), req)

    if err != nil {
        log.Fatalf("Could not greet: %v",err)
    }
}

//update
func StatusUpdate(c *gin.Context){
    var updateRequest pb.UpdateRequest
    if err:=c.BindJSON(&updateRequest);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err=StatusClient.Update(context.Background(),updateRequest)

    if err != nil {
        log.Fatalf("Could not greet: %v",err)
    }
}
//Delete
func StatusDelete(c *gin.Context){
    var getRequest pb.GetResquest
    if err:=BindJSON(&getRequest);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err=StatusClient.Delete(context.Background(),getRequest)

    if err != nil{
        log.Fatalf("Could not greet: %v",err)
    }
}

//Like
func StatusLike(c *gin.Context){
    var likeRequest pb.LikeRequest
    if err:=c.BindJSON(&likeRquest);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err=StatusClient.Like(context.Background(),likeRequest)

    if err != nil{
        log.Fatalf("Could not greet: %v",err)
    }
}

func StatusCreateComment(c *gin.Context){
    var req pb.CreatCommentRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    /*req := &pb.CreateCommentRequest{
        UserId:  0,
        StatusId: 3488,
        Content: "后废物废物分为",
    }*/

    _, err = StatusClient.CreateComment(context.Background(), req)

    if err != nil{
        log.Fatalf("Could not greet: %v",err)
    }
}

func StatusListComment(c *gin.Context){
    var req pb.CommentListRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }
    /*req := &pb.CommentListRequest{
        StatusId: 3488,
        Offset: 0,
        Limit: 20,
        Lastid: 0,
    }*/

    resp, err := client.ListComment(context.Background(), req)

    if err != nil{
        log.Fatalf("Could not greet: %v",err)
    }

    fmt.Println(resp.List, resp.Count)
}

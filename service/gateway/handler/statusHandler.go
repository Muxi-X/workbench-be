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
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err != nil{
        c.JSON(400,gin.H{
            "message":"wrong",
            })
        return
    }
    resp, err2 := StatusClient.Get(context.Background(), req)

    if err2 != nil {
        log.Fatalf("Could not greet: %v", err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func StatusList(c *gin.Context){
    var req pb.ListRequest
    if err:=c.BindJSON(&req);err != nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp, err2 := StatusClient.List(context.Background(), req)
    if err2 != nil {
        log.Fatalf("Could not greet: %v", err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func StatusCreate(c *gin.Context){
    var req pb.CreateRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    /*req := &pb.CreateRequest{
        UserId:  0,
        Title:   "哈哈哈哈😁",
        Content: "后废物废物分为",
    }*/

    _, err2 = StatusClient.Create(context.Background(), req)

    if err2 != nil {
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,gin.H{
        "message":"ok",
    })
}

//update
func StatusUpdate(c *gin.Context){
    var updateRequest pb.UpdateRequest
    if err:=c.BindJSON(&updateRequest);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2=StatusClient.Update(context.Background(),updateRequest)

    if err2 != nil {
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,gin.H{
        "message":"ok",
    })
}
//Delete
func StatusDelete(c *gin.Context){
    var getRequest pb.GetResquest
    if err:=BindJSON(&getRequest);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2=StatusClient.Delete(context.Background(),getRequest)

    if err2 != nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,gin.H{
        "message":"ok",
    })
}

//Like
func StatusLike(c *gin.Context){
    var likeRequest pb.LikeRequest
    if err:=c.BindJSON(&likeRquest);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2=StatusClient.Like(context.Background(),likeRequest)

    if err2 != nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,gin.H{
        "message":"ok",
    })
}

func StatusCreateComment(c *gin.Context){
    var req pb.CreatCommentRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    /*req := &pb.CreateCommentRequest{
        UserId:  0,
        StatusId: 3488,
        Content: "后废物废物分为",
    }*/

    _, err2 = StatusClient.CreateComment(context.Background(), req)

    if err2 != nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,gin.H{
        "message":"ok",
    })
}

func StatusListComment(c *gin.Context){
    var req pb.CommentListRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }
    /*req := &pb.CommentListRequest{
        StatusId: 3488,
        Offset: 0,
        Limit: 20,
        Lastid: 0,
    }*/

    resp, err2 := StatusClient.ListComment(context.Background(), req)

    if err2 != nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    //fmt.Println(resp.List, resp.Count)
    c.JSON(200,resp)
}

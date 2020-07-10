package handler

import (
    "context"
    "fmt"
    "log"

    "github.com/opentracing/opentracing-go"

    tracer "muxi-workbench-project-client/tracer"

    opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"

    pb "muxi-workbench-project/proto"

    handler "muxi-workbench/pkg/handler"

    micro "github.com/micro/go-micro"
)

var ProjectService micro.Service
var ProjectClient pb.ProjectServiceClient

func ProjectInit(ProjectService micro.Service,ProjectClient pb.ProjectServiceClient){
    ProjectService=micro.NewService(micro.Name("workbench.cli.project"),
        micro.WrapClient(
            opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
        ),
        micro.WrapCall(handler.ClientErrorHandlerWrapper()),
    )

    ProjectService.Init()

    ProjectClient = pb.NewFeedServiceClient("workbench.service.project", FeedService.Client())
}

func GetProjectList(c *gin.Context){
    var req pb.GetProjectListRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient.List(context.Background(),req)
    if err2 != nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }
    //response
    c.JSON(200,resp)
}

func GetProjectInfo(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient.GetProjectInfo(context.Background(),req)
    if err2 !=nil{
        //do something
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    //response
    c.JSON(200,resp)
}

func UpdateProjectIfno(c *gin.Context){
    var req pb.ProjectInfo
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClinet.UpdateProjectInfo(context.Background(),req)
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

func DeleteProject(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err != nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClient.DeleteProject(context.Background(),req)
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

func GetFileTree(c *gin.Context){
    var req pb.GetRequest
    if err := c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient.GetFileTree(context.Background(),req)
    if err2 != nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func GetDocTree(c *gin.Context){
    var req pb.GetRequest
    if err := c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient.GetDocTree(context.Background(),req)
    if err2!=nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func UpdateFileTree(c *gin.Context){
    var req pb.UpdateTreeRequest
    if err := c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClient.UpdateFileTree(context.Background(),req)
    if err2!=nil{
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

func UpdateDocTree(c *gin.Context){
    var req pb.UpdateTreeRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClient(context.Background(),req)
    if err2!=nil{
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

func GetMembers(c *gin.Context){
    var req pb.GetMemberListRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient.GetMembers(context.Background(),req)
    if err2!=nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func UpdateMembers(c *gin.Context){
    var req pb.UpdateMemberRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClient(context.Background(),req)
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

func GetProjectIdsForUser(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
            })
        return
    }

    resp,err2:=ProjectClient(context.Background(),req)
    if err2 != nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func CreateFile(c *gin.Context){
    var req pb.CreateFileRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClient(context.Background(),req)
    if err2!=nil{
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

func DeleteFile(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClient(context.Background(),req)
    if err2 != nil{
        log.Fatal("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,gin.H{
        "message":"ok",
    })
}

func GetFileDetail(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient(context.Background(),req)
    if err2!=nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func GetFileInfoList(c *gin.Context){
    var req pb.GetInfoByIdsRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient(context.Background(),req)
    if err2!=nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func GetFileFolderInfoList(c *gin.Context){
    var req pb.GetInfoByIdsRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient(context.Background(),req)
    if err2!=nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func CreateDoc(c *gin.Context){
    var req pb.CreateDocRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClient(context.Background(),req)
    if err2!=nil{
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

func UpdateDoc(c *gin.Context){
    var req pb.UpdateDocRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClient(context.Background(),req)
    if err2!=nil{
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

func DeleteDoc(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    _,err2:=ProjectClient(context.Background(),req)
    if err2!=nil{
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

func GetDocDetail(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient(context.Background(),req)
    if err2!=nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func GetDocInfoList(c *gin.Context){
    var req pb.GetInfoByIdsRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient(context.Background(),req)
    if err!=nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

func GetDocFolderInfoList(c *gin.Context){
    var req pb.GetInfoByIdsRequest
    if err:=c.BindJSON(req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
        return
    }

    resp,err2:=ProjectClient(context.Background(),req)
    if err2 != nil{
        log.Fatalf("Could not greet: %v",err2)
        c.JSON(500,gin.H{
            "message":"wrong",
        })
        return
    }

    c.JSON(200,resp)
}

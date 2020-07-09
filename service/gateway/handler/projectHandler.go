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
    }

    resp,err:=ProjectClient.List(context.Background(),req)
    if err != nil{
        panic(err)//not panic
    }
    //response
}

func GetProjectInfo(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient.GetProjectInfo(context.Background(),req)
    if err !=nil{
        //do something
    }

    //response
}

func UpdateProjectIfno(c *gin.Context){
    var req pb.ProjectInfo
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClinet.UpdateProjectInfo(context.Background(),req)
    if err != nil{
        //
    }
    //
}

func DeleteProject(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err != nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClient.DeleteProject(context.Background(),req)
    if err != nil{
        //
    }
    //
}

func GetFileTree(c *gin.Context){
    var req pb.GetRequest
    if err := c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient.GetFileTree(context.Background(),req)
    if err != nil{
        //
    }

    //
}

func GetDocTree(c *gin.Context){
    var req pb.GetRequest
    if err := c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient.GetDocTree(context.Background(),req)
    if err!=nil{
        //
    }
    //
}

func UpdateFileTree(c *gin.Context){
    var req pb.UpdateTreeRequest
    if err := c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClient.UpdateFileTree(context.Background(),req)
    if err!=nil{
        //
    }
}

func UpdateDocTree(c *gin.Context){
    var req pb.UpdateTreeRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
}

func GetMembers(c *gin.Context){
    var req pb.GetMemberListRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient.GetMembers(context.Background(),req)
    if err!=nil{
        //
    }
    //
}

func UpdateMembers(c *gin.Context){
    var req pb.UpdateMemberRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClient(context.Background(),req)
    if err != nil{
        //
    }
}

func GetProjectIdsForUser(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
            })
    }

    resp,err:=ProjectClient(context.Background(),req)
    if err != nil{
        //
    }
    //
}

func CreateFile(c *gin.Context){
    var req pb.CreateFileRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
}

func DeleteFile(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClient(context.Background(),req)
    if err != nil{
        //
    }
}

func GetFileDetail(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
    //
}

func GetFileInfoList(c *gin.Context){
    var req pb.GetInfoByIdsRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
    //
}

func GetFileFolderInfoList(c *gin.Context){
    var req pb.GetInfoByIdsRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
    //
}

func CreateDoc(c *gin.Context){
    var req pb.CreateDocRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
}

func UpdateDoc(c *gin.Context){
    var req pb.UpdateDocRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
}

func DeleteDoc(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    _,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
}

func GetDocDetail(c *gin.Context){
    var req pb.GetRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
    //
}

func GetDocInfoList(c *gin.Context){
    var req pb.GetInfoByIdsRequest
    if err:=c.BindJSON(&req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient(context.Background(),req)
    if err!=nil{
        //
    }
    //
}

func GetDocFolderInfoList(c *gin.Context){
    var req pb.GetInfoByIdsRequest
    if err:=c.BindJSON(req);err!=nil{
        c.JSON(400,gin.H{
            "message":"wrong",
        })
    }

    resp,err:=ProjectClient(context.Background(),req)
    if err != nil{
        //
    }
    //
}

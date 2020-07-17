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

var ProjectService micro.Service
var ProjectClient pb.ProjectServiceClient

func ProjectInit(ProjectService micro.Service, ProjectClient pb.ProjectServiceClient) {
	ProjectService = micro.NewService(micro.Name("workbench.cli.project"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	ProjectService.Init()

	ProjectClient = pb.NewProjectServiceClient("workbench.service.project", ProjectService.Client())
}

type getProInfoResponse struct {
	Projectid   int    `json:"projectid"`
	Projectname string `json:"projectname"`
	Intro       string `json:"intro"`
	Usercount   int    `json:"usercount"`
}

type deleteRequest struct {
	UserId      int    `json:"userid"`
	Projectname string `json:"projectname"`
}

type updateRequest struct {
	UserId      int    `json:"userid"`
	Projectname string `json:projectname"`
	Intro       string `json:"intro"`
	Usercount   int    `json:"usercount"`
}

type projectListItem struct {
	Id   int    `json:"uid"`
	Name string `json:"username"`
	Logo string `json:logo"`
}

type memberListItem struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	GroupName string `json:"groupname"`
	Role      int    `json:"role"`
}

type getMemberResponse struct {
	Count      int              `json:"count"`
	Memberlist []memberListItem `json:"memberlist"`
}

type updateMemberRequest struct {
	UserId      int    `json:"userid"`
	ProjectName string `json:"projectname"`
	Userlist    []int  `json:"userlist"`
}

type getProjectListRequest struct {
	UserId int `json:"userid"`
}

type getProjectListResponse struct {
	Projectlist []projectListItem `json:"projectlist"`
}

type getFileTreeResponse struct {
	Filetree string `json:"filetree"`
}

type updateFileTreeRequest struct {
	UserId      int    `json:"userid"`
	Projectname string `json:"projectname"`
	Filetree    string `json:"filetree"`
}

type getDocTreeResponse struct {
	Doctree string `json:"doctree"`
}

type updateDocTreeRequest struct {
	UserId      int    `json:"userid"`
	Projectname string `json:"projectname"`
	Doctree     string `json:"doctree"`
}

type createFileRequest struct {
	UserId   int    `json:"userid"`
	Pid      int    `json:"projectid"`
	Filename string `json:"filename"`
	Hashname string `json:"hashname"`
	Url      string `json:"url"`
	Fid      int    `json:"fid"`
}

type deleteFileRequest struct {
	UserId   int    `json:"userid"`
	Filename string `json:"filename"`
}

type createDocRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Pid     int    `json:"projectid"`
	Docname string `json:"docname"`
	UserId  int    `json:"userid"`
}

type getDocDetailResponse struct {
	Id           int    `json:"docid"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Creator      string `json:"creator"`
	Createtime   string `json:"createtime"`
	Lasteditor   string `json:"lasteditor"`
	Lastedittime string `json:"lastedittime"`
}

type deleteDocRequest struct {
	UserId  int    `json:"userid"`
	Docname string `json:"docname"`
}

type updateDocRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"userId"`
}

// 下面是待抽离的 api
func GetProjectIdsForUser(c *gin.Context) {
	var req pb.GetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetProjectIdsForUser(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetFileDetail(c *gin.Context) {
	var req pb.GetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetFileDetail(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetFileInfoList(c *gin.Context) {
	var req pb.GetInfoByIdsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetFileInfoList(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetFileFolderInfoList(c *gin.Context) {
	var req pb.GetInfoByIdsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetFileFolderInfoList(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetDocInfoList(c *gin.Context) {
	var req pb.GetInfoByIdsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetDocInfoList(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

func GetDocFolderInfoList(c *gin.Context) {
	var req pb.GetInfoByIdsRequest
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{
			"message": "wrong",
		})
		return
	}

	resp, err2 := ProjectClient.GetDocFolderInfoList(context.Background(), &req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		c.JSON(500, gin.H{
			"message": "wrong",
		})
		return
	}

	c.JSON(200, resp)
}

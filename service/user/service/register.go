package service

import (
	"context"
	"fmt"

	errno "muxi-workbench-user/errno"
	model "muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	"muxi-workbench-user/util"
	e "muxi-workbench/pkg/err"
)

var (
	registerURL = ""
)

// Register ... 注册
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest, res *pb.Response) error {

	// muxi-auth-service 注册用户

	bodyData := map[string]string{"username": req.Email, "password": req.Password}

	body, err := util.SendHTTPRequest(registerURL, "POST", &util.RequestData{
		BodyData:    bodyData,
		ContentType: util.JsonData,
	})
	if err != nil {

	}
	fmt.Println(body)

	// 本地 user 数据库创建用户

	user := &model.UserModel{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := user.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

package service

import (
	"context"

	"muxi-workbench-user/errno"
	"muxi-workbench-user/model"
	"muxi-workbench-user/pkg/auth"
	pb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
)

// Register ... 注册
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest, res *pb.Response) error {

	// muxi-auth-service 注册用户
	if err := auth.RegisterRequest(req.Name, req.Email, req.Password); err != nil {
		return e.ServerErr(errno.ErrRegister, err.Error())
	}

	// 本地 user 数据库创建用户

	// 用户是否存在
	if CheckUserExisted(req.Name, req.Email) {
		return e.ServerErr(errno.ErrUserExisted, "")
	}

	// 创建用户
	user := &model.UserModel{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := user.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// CheckUserExisted check whether the user exists by name and email.
func CheckUserExisted(name, email string) bool {
	sameEmailChannel, sameNameChannel, done := make(chan bool), make(chan bool), make(chan struct{})
	defer close(done)
	defer close(sameEmailChannel)
	defer close(sameNameChannel)

	// 检查邮箱
	go func(email string) {
		u, _ := model.GetUserByEmail(email)
		select {
		case <-done:
			return
		default:
			if u != nil {
				sameEmailChannel <- false
			} else {
				sameEmailChannel <- true
			}
		}
	}(email)

	// 检查用户名
	go func(name string) {
		u, _ := model.GetUserByName(name)
		select {
		case <-done:
			return
		default:
			if u != nil {
				sameNameChannel <- false
			} else {
				sameNameChannel <- true
			}
		}
	}(name)

	var userExisted = false

	// 最多循环两次
	for round := 0; !userExisted && round < 2; round++ {
		select {
		case result := <-sameEmailChannel:
			if !result {
				userExisted = true
				break
			}
		case result := <-sameNameChannel:
			if !result {
				userExisted = true
				break
			}
		}
	}
	return userExisted
}

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
	if userExisted, err := CheckUserExisted(req.Name, req.Email); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	} else if userExisted {
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
func CheckUserExisted(name, email string) (bool, error) {
	sameEmailChannel, sameNameChannel, errChannel := make(chan bool), make(chan bool), make(chan error)

	// 检查邮箱
	go func(email string) {
		u, err := model.GetUserByEmail(email)
		if err != nil {
			errChannel <- err
		} else if u != nil {
			sameEmailChannel <- true
		} else {
			sameEmailChannel <- false
		}
	}(email)

	// 检查用户名
	go func(name string) {
		u, err := model.GetUserByName(name)
		if err != nil {
			errChannel <- err
		} else if u != nil {
			sameNameChannel <- true
		} else {
			sameNameChannel <- false
		}
	}(name)

	var userExisted = false
	var err error

	// 循环两次
	for round := 0; round < 2; round++ {
		select {
		case curErr := <-errChannel:
			err = curErr
		case result := <-sameEmailChannel:
			if result {
				userExisted = true
			}
		case result := <-sameNameChannel:
			if result {
				userExisted = true
			}
		}
	}

	close(sameEmailChannel)
	close(sameNameChannel)
	close(errChannel)

	return userExisted, err
}

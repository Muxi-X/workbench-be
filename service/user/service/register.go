package service

import (
	"context"

	"muxi-workbench-user/errno"
	"muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"

	uuid "github.com/satori/go.uuid"
)

type RegisterInfo struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func RegisterUser(info *RegisterInfo) error {
	// 本地 user 数据库创建用户
	name := info.Name
	// 用户是否存在
	if userExisted, nameExisted, err := CheckUserExisted(info.Name, info.Email); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	} else if userExisted {
		return e.ServerErr(errno.ErrUserExisted, "")
	} else if nameExisted {
		u4 := uuid.NewV4()
		name = "用户" + u4.String()
	}

	// 创建用户
	user := &model.UserModel{
		Name:   name,
		Email:  info.Email,
		Role:   constvar.AuthLevelNormal,
		TeamID: 0, // 默认无team
	}
	return user.Create()
}

// Register ... 注册
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest, res *pb.Response) error {
	info := &RegisterInfo{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := RegisterUser(info); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// CheckUserExisted check whether the user exists by name and email.
func CheckUserExisted(name, email string) (bool, bool, error) {
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
	var nameExisted = false
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
				nameExisted = true
			}
		}
	}

	close(sameEmailChannel)
	close(sameNameChannel)
	close(errChannel)

	return userExisted, nameExisted, err
}

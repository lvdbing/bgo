package service

import (
	"errors"
	"time"

	"github.com/lvdbing/bgo/internal/model"
)

func (svc *Service) Register(req *model.RegisterReq) (*model.User, error) {
	var user model.User
	svc.userDB.Where("username = ?", req.Username).First(&user)
	if user.ID > 0 {
		return nil, errors.New("用户已经存在！")
	}

	now := uint32(time.Now().Unix())

	user.CreatedAt = now
	user.Username = req.Username
	user.Password = req.Password
	user.Phone = req.Phone

	result := svc.userDB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (svc *Service) Login(req *model.LoginReq) (*model.User, error) {
	return nil, nil
}

func (svc *Service) GetUsers(req *model.UserReq) ([]model.User, int, error) {
	return nil, 0, nil
}

func (svc *Service) GetUserByID(id uint32) (*model.User, error) {
	return nil, nil
}

func (svc *Service) CreateUser(user *model.User) error {
	return nil
}

func (svc *Service) UpdateUser(user *model.User) error {
	return nil
}

func (svc *Service) DeleteUser(id uint32) error {
	return nil
}

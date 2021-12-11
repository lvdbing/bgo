package service

import (
	"time"

	"github.com/lvdbing/bgo/internal/model"
)

func (svc *Service) Register(req *model.RegisterReq) (*model.User, error) {
	now := uint32(time.Now().Unix())

	var user model.User
	user.CreatedAt = now
	user.Username = req.Username
	user.Password = req.Password
	user.Phone = req.Phone

	err := svc.dao.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

// func (dao *Dao) Register(req *model.RegisterReq) (*model.User, error) {
// 	var user model.User
// 	now := uint32(time.Now().Unix())
// 	user.CreatedAt = now
// 	user.Username = req.Username
// 	user.Password = req.Password
// 	user.Phone = req.Phone

// 	result := dao.UserDB.Create(&user)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &user, nil
// }

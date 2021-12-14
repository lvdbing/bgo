package dao

import (
	"github.com/lvdbing/bgo/internal/model"
)

func (dao *Dao) GetUserByID(id uint32) (*model.User, error) {
	return nil, nil
}

func (dao *Dao) GetUsers(req *model.UserReq) (users []model.User, err error) {
	return
}

func (dao *Dao) CreateUser(user *model.User) error {
	result := dao.UserDB.Create(user)
	return result.Error
}

func (dao *Dao) UpdateUser(user *model.User) (*model.User, error) {
	return nil, nil
}

func (dao *Dao) DeleteUser(userID uint32) error {
	return nil
}

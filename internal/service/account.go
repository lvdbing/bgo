package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/lvdbing/bgo/internal/helper/jwthelper"
	"github.com/lvdbing/bgo/internal/model"
	"github.com/lvdbing/bgo/pkg/utils"
)

func (svc *Service) Register(req *model.RegisterReq) (*model.User, error) {
	// 查询用户是否已存在。
	var user model.User
	svc.userDB.Where("username = ?", req.Username).First(&user)
	if user.ID > 0 {
		return nil, errors.New("用户已经存在")
	}

	// 保存用户信息到数据库。
	now := uint32(time.Now().Unix())
	user.CreatedAt = now
	user.Username = req.Username
	user.Password = utils.EncodeMD5(req.Password)
	user.Phone = req.Phone

	result := svc.userDB.Create(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("insert user err: %v", result.Error)
	}

	return &user, nil
}

func (svc *Service) Login(req *model.LoginReq) (*model.UserToken, error) {
	// 根据用户名和密码查询用户。
	var user model.User
	password := utils.EncodeMD5(req.Password)
	svc.userDB.Where("username=? and password=?", req.Username, password).First(&user)
	if user.ID <= 0 {
		return nil, errors.New("用户名或密码不正确")
	}

	// 生成token。
	token, err := jwthelper.GenerateToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		return nil, err
	}
	// 保存token到数据库。
	var jwtToken model.JwtToken
	jwtToken.ID = user.ID
	jwtToken.Token = token
	result := svc.userDB.Save(jwtToken)
	if result.Error != nil {
		return nil, fmt.Errorf("save jwt_token err: %v", result.Error)
	}

	// 返回用户和token信息。
	var userToken model.UserToken
	userToken.User = user
	userToken.Token = jwtToken.Token
	return &userToken, nil
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

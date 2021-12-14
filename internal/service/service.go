package service

import (
	"context"

	"github.com/lvdbing/bgo/global"
	// "github.com/lvdbing/bgo/internal/dao"
	"gorm.io/gorm"
)

type Service struct {
	ctx context.Context
	// dao *dao.Dao
	userDB *gorm.DB
}

func NewService(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	// svc.dao = dao.NewDao(global.UserDB)
	svc.userDB = global.UserDB
	return svc
}

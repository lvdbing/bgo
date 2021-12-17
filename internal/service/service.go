package service

import (
	"context"

	"github.com/lvdbing/bgo/global"
	"gorm.io/gorm"
)

type Service struct {
	ctx    context.Context
	userDB *gorm.DB
}

func NewService(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.userDB = global.UserDB
	return svc
}

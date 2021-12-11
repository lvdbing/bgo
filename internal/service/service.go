package service

import (
	"context"

	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func NewService(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.NewDao(global.UserDB)
	return svc
}

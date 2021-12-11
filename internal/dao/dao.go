package dao

import "gorm.io/gorm"

type Dao struct {
	UserDB *gorm.DB
}

func NewDao(userDB *gorm.DB) *Dao {
	return &Dao{UserDB: userDB}
}

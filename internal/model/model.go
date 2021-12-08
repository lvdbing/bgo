package model

import (
	"fmt"

	"github.com/lvdbing/bgo/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID         uint32 `json:"id"`          // ID
	CreatedAt  uint32 `json:"created_at"`  // 创建时间
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedAt uint32 `json:"modified_at"` // 修改时间
	ModifiedBy string `json:"modified_by"` // 修改人
	IsDel      uint8  `json:"is_del"`      // 是否已删除
	DeletedAt  uint32 `json:"deleted_at"`  // 删除时间
}

func NewDBEngine(dbSetting *global.DatabaseSettings, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbSetting.Username,
		dbSetting.Password,
		dbSetting.Host,
		dbName,
		dbSetting.Charset,
		dbSetting.ParseTime)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	if dbSetting.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(dbSetting.MaxIdleConns)
	}

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	if dbSetting.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(dbSetting.MaxOpenConns)
	}

	// // SetConnMaxLifetime 设置了连接可复用的最大时间。
	// sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

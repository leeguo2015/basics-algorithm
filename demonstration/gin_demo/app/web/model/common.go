package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Model 定义
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedTime time.Time
	UpdatedTime time.Time
	DeletedTime *time.Time
}

// 初始化
func init() {
	var err error
	var constr string
	constr = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "47.108.140.163", 8166, "gin_test")
	db, err = gorm.Open(mysql.Open(constr), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

	err = db.AutoMigrate(&todoModel{})
	if err != nil {
		panic(fmt.Sprintf("DB auto migrate err: %s", err))
	}
	sqlDB, err := db.DB()

	if err != nil {
		panic(fmt.Sprintf("DB auto migrate err: %s", err))
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

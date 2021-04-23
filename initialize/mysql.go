package initialize

import (
	"fmt"
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQL() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/research_sys?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("初始化mysql异常: %v", err))
	}
	global.Mysql = db
	// 表结构
	autoMigrate()
}

func autoMigrate() {
	err := global.Mysql.AutoMigrate(
		new(models.User),
	)
	if err != nil {
		panic(fmt.Sprintf("数据库迁移异常: %v", err))
	}
}

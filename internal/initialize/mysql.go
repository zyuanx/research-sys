package initialize

import (
	"fmt"
	"gin-research-sys/internal/conf"
	"gin-research-sys/internal/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func MySQL() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.database"),
		viper.GetString("mysql.charset"),
		viper.GetString("mysql.parseTime"),
		viper.GetString("mysql.loc"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("初始化mysql异常: %v", err))
	}
	conf.Mysql = db
	// 表结构
	autoMigrate()

}

func autoMigrate() {
	err := conf.Mysql.AutoMigrate(
		model.User{}, model.Role{}, model.Permission{},
		model.Research{}, model.Record{}, model.OpenRecord{},
	)
	if err != nil {
		panic(fmt.Sprintf("数据库迁移异常: %v", err))
	}
	log.Println("数据库迁移完成")
}

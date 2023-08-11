package mysql

import (
	"fmt"
	"time"

	"github.com/zyuanx/research-sys/internal/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(MySQLConfig *config.MySQLConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		MySQLConfig.Username,
		MySQLConfig.Password,
		MySQLConfig.Host,
		MySQLConfig.Port,
		MySQLConfig.Database,
		MySQLConfig.Charset,
		MySQLConfig.ParseTime,
		MySQLConfig.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetConnMaxLifetime(time.Hour)
	// 设置连接池大小
	sqlDb.SetMaxOpenConns(MySQLConfig.MaxOpenSize)
	sqlDb.SetMaxIdleConns(MySQLConfig.MaxIdleSize)
	return db
}

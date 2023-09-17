package database

import (
	"fmt"
	"path"
	"time"

	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/pkg/config"
	"github.com/zyuanx/research-sys/internal/pkg/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
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
	migrate(db)
	return db
}

func NewSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path.Join(global.RootDir, "database.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
}

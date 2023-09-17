package main

import (
	"flag"
	"fmt"

	"path"

	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/controller"
	"github.com/zyuanx/research-sys/internal/pkg/config"
	"github.com/zyuanx/research-sys/internal/pkg/database"
	"github.com/zyuanx/research-sys/internal/pkg/global"

	"github.com/zyuanx/research-sys/internal/router"
	"github.com/zyuanx/research-sys/internal/service"
	"github.com/zyuanx/research-sys/tools"
)

func init() {
	global.RootDir = tools.GetWorkingDirPath()
}

var ConfigFilePath string

func main() {
	flag.StringVar(&ConfigFilePath,
		"c", path.Join(global.RootDir, "configs", "dev.yaml"),
		"-c 选项用于指定要使用的配置文件")
	flag.Parse()

	c := config.NewViper(ConfigFilePath)

	global.MySQL = database.NewSqlite()

	gin.SetMode(c.Server.Mode)
	r := gin.Default()
	// initialize.Zap()
	// initialize.CreateAdmin()
	s := service.NewService(global.MySQL)
	controller.NewController(s)
	router.SetupRouter(r, s)

	r.Run(fmt.Sprintf("0.0.0.0:%d", c.Server.Port))
}

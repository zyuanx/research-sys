package internal

import (
	"gin-research-sys/internal/initialize"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	initialize.Zap()
	initialize.Viper()
	initialize.MySQL()
	initialize.MongoDB()
	initialize.CreateAdmin()
	//initialize.Redis()
	initialize.Casbin()
	return initialize.Router()
}

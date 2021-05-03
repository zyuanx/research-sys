package initialize

import (
	"gin-research-sys/pkg/log"
	"go.uber.org/zap"
)

func Zap() {
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE)
}

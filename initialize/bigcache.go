package initialize

import (
	"gin-research-sys/pkg/global"
	"github.com/allegro/bigcache/v3"
	"time"
)

func BigCache() {
	global.Cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
}

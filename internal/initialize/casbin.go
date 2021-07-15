package initialize

import (
	"fmt"
	"gin-research-sys/internal/conf"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Casbin() {
	adapter, err := gormadapter.NewAdapterByDB(conf.Mysql) // mysql 适配器
	if err != nil {
		panic(fmt.Sprintf("初始化casbin异常: %v", err))
	}
	// 通过mysql适配器新建一个enforcer
	conf.Enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("初始化casbin异常: %v", err))
	}
	conf.Enforcer.EnableLog(true)    // 日志记录
	err = conf.Enforcer.LoadPolicy() // 加载策略规则
	if err != nil {
		fmt.Println("loadPolicy error")
	}
}

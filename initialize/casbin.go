package initialize

import (
	"fmt"
	"gin-research-sys/pkg/global"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Casbin() {
	// mysql 适配器
	adapter, err := gormadapter.NewAdapterByDB(global.Mysql)
	if err != nil {
		panic(fmt.Sprintf("初始化casbin异常: %v", err))
	}
	// 通过mysql适配器新建一个enforcer
	global.Enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("初始化casbin异常: %v", err))
	}
	// 日志记录
	global.Enforcer.EnableLog(true)
	if ok, _ := global.Enforcer.AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
		fmt.Println("Policy已经存在")
	} else {
		fmt.Println("增加成功")
	}
}

package initialize

import (
	"fmt"
	"gin-research-sys/internal/conf"
	"github.com/casbin/casbin/v2"
)

func Casbin() {
	//adapter, err := gormadapter.NewAdapterByDB(conf.Mysql)
	//if err != nil {
	//	panic(fmt.Sprintf("initalize casbin err: %v", err))
	//}
	//conf.Enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", adapter)

	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", "config/rbac_policy.csv")
	if err != nil {
		fmt.Println("initalize casbin err", err)
	}
	enforcer.EnableLog(true)
	if err = enforcer.LoadPolicy(); err != nil {
		fmt.Println("loadPolicy error", err)
	}
	conf.Enforcer = enforcer
}

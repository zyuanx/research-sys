package main

import (
	"gin-research-sys/initialize"
)

// @title Research Sys API
// @version 1.0
// @description An example of gin
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080


// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	initialize.MySQL()
	initialize.MongoDB()
	initialize.Casbin()
	r := initialize.Routers()
	panic(r.Run())
}

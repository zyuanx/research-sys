package main

import (
	"gin-research-sys/internal"
)

// @title Research sys API
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
	r := internal.App()
	panic(r.Run())
}

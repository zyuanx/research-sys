package main

import (
	"gin-research-sys/initialize"
)

func main() {
	initialize.MySQL()
	r := initialize.Routers()

	panic(r.Run())
}

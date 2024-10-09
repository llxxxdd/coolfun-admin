package main

import (
	"fmt"

	"coolfun-admin/configs"
	"coolfun-admin/data"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// init
	data.InitDataConfig()

	// start server
	c := configs.Data()
	h := server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", c.Host, c.Port)))

	register(h)
	h.Spin()
}

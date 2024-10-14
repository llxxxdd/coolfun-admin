package main

import (
	"fmt"

	"coolfun-admin/configs"
	"coolfun-admin/data"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// 初始化
	data.InitDataConfig()
	// 初始化 Web 服务器实例。
	c := configs.Data()
	h := server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", c.Host, c.Port)))
	// 注册路由
	register(h)
	// 启动服务器
	h.Spin()
}

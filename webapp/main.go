package main

import (
	"fmt"
	"webapp/dao/mysql"
	"webapp/logger"
	"webapp/routes"
	"webapp/settings"
)

//go web 开发通用脚手架

func main() {

	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("init setting failed... err:", err)
		return
	}
	//2.初始化日志
	if err := logger.InitLogger(); err != nil {
		fmt.Println("init logger failed... err:", err)
		return
	}
	//3.初始化Mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Println("init mysql failed... err:", err)
		return
	}
	//5.注册路由
	r := routes.SetUpRouter()
	//6.启动服务
	r.Run(":8080")
}

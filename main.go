package main

import (
	"gin_mvt/app/urls"

	"github.com/gin-gonic/gin"
)

func main() {

	//创建引擎
	engine := gin.Default()

	//初始化路由
	urls.Init(engine)

	//启动服务器
	engine.Run()

}

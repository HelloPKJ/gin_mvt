package urls

import (
	"gin_mvt/app/views"

	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {

	//注册路由
	myUrl := e.Group("/" /*这里可以注册中间件*/)
	{
		BindView(myUrl, "/index", views.Index)
	}

}

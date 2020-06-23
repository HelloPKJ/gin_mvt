package views

import (
	"github.com/gin-gonic/gin"
)

func Index(e *gin.RouterGroup, url string) {
	e.GET(url, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "get hello world !",
		})
	})

	e.POST(url, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "post hello world !",
		})
	})
}

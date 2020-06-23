package urls

import (
	"github.com/gin-gonic/gin"
)

func BindView(e *gin.RouterGroup, url string, view func(*gin.RouterGroup, string)) {
	view(e, url)
}

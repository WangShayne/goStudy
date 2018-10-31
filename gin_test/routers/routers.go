package routers

import (
	. "github.com/study/gin_test/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	routers := gin.Default()
	routers.LoadHTMLGlob("views/*")
	routers.GET("/", ShowHTML)
	return routers
}

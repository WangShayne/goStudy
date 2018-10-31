package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回HTML文件
func ShowHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "this is home html",
	})
}

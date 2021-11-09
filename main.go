package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/topic/:topic_id", func(c *gin.Context) {
		c.String(http.StatusOK, "获取帖子ID为 %s", c.Param("topic_id"))
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "欢迎来到首页")
	})

	router.Run(":8080") // 配置端口号
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	路由的使用
    响应格式
*/

type Article struct {
	// 配置返回的json格式
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	route := gin.Default()

	// 响应字符串
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "欢迎来到首页")
	})

	// 响应 json
	route.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"msg":     "你好Gin",
		})
	})

	// gin.H 和 map[string]interface{} 相同
	route.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好Gin2",
		})
	})

	// 使用结构体 响应 json
	route.GET("/json3", func(c *gin.Context) {
		a := Article{
			Title:   "123",
			Desc:    "567",
			Content: "987",
		}
		c.JSON(http.StatusOK, a)
	})

	// 响应 JSONP (主要为了处理跨域问题)
	// http://localhost:8080/jsonp?callback=xxx   xxx( { jsonp: "我是一个jsonp", url: "www.baidu.com" } )
	route.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(http.StatusOK, map[string]interface{}{
			"jsonp": "我是一个jsonp",
			"url":   "www.baidu.com",
		})
	})

	// 响应 XML 数据
	route.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"return":  "xml",
		})
	})

	// 响应html
	route.GET("/html", func(c *gin.Context) {
		// 配置模板文件
		route.LoadHTMLGlob("templates/*")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "我是后台数据",
		})
	})
	route.Run(":8080")
}

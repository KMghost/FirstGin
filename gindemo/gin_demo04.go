package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	GET/POST 及动态路由传值
*/
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginForm struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// Get 请求参数
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	// POST
	r.POST("/", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// 绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}

		err := c.ShouldBind(&user) //绑定数据到结构体

		if err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.POST("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err,
			})
		}
	})

	r.Run()
}

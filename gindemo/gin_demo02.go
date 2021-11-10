package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	路由的使用
    响应格式
*/
type Acticle struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*") // 导入templates  **：代表所有文件夹，*: 代表所有文件
	r.GET("/index", func(c *gin.Context) {
		a := &Acticle{
			Title:   "主页",
			Content: "如何起飞",
		}
		c.HTML(http.StatusOK, "demo2/index.html", gin.H{
			"title": "新闻页面",
			"new":   a,
		})
	})

	r.GET("/news", func(c *gin.Context) {
		a := &Acticle{
			Title:   "主页",
			Content: "如何起飞",
		}
		c.HTML(http.StatusOK, "demo2/news.html", gin.H{
			"title":  "新闻消息",
			"news":   a,
			"choice": 1,
			"array":  []int{3, 5, 6, 3, 2, 1},
			"newsList": []interface{}{
				&Acticle{
					Title:   "主页1",
					Content: "如何起飞1",
				},
				&Acticle{
					Title:   "主页2",
					Content: "如何起飞2",
				},
			},
			"testSlice": []int{},
		})
	})
	r.Run()
}

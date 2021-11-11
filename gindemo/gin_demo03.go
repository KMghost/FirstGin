package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

/*
	路由的使用
    响应格式
*/
type Acticle struct {
	Title   string
	Content string
}

func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")

}

func main() {
	r := gin.Default()
	// 自定义模板函数，需要放在 LoadHTMLGlob 前面
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	r.LoadHTMLGlob("templates/**/*") // 导入templates  **：代表所有文件夹，*: 代表所有文件

	r.GET("/news", func(c *gin.Context) {
		a := &Acticle{
			Title:   "主页",
			Content: "如何起飞",
		}
		c.HTML(http.StatusOK, "demo3/news.html", gin.H{
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
			"time":      1636604830,
		})
	})
	r.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器发请求携带的query string参数
		name := c.Query("query") // 通过Query获取请求携带的querystring参数
		age := c.Query("age")
		//name := c.DefaultQuery("query", "somebody")//返回不到就用指定的默认值
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9090")
}

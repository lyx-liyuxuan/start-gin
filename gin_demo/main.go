package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	// 200 状态码
	c.JSON(200, gin.H{
		"messange": "Hello Golang\nHello World",
	})
}

func main() {
	// 返回的是默认的路由引擎
	r := gin.Default()

	// 指定用户使用GET请求访问/hello时，执行sayhello这个函数
	r.GET("/hello", sayHello)

	/*	r.GET("/book", ...)
		r.GET("/create_book", ...)
		r.GET("/update_book", ...)
		r.GET("/delete_book", ...)*/

	// RESTful风格
	/*	r.GET()  用来获取资源
		r.POST()  用来创建资源
		r.PUT()  用来更新资源
		r.DELETE()  用来删除资源*/

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	// http.StatusOK 表示状态码200
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	// 启动服务
	r.Run()
}

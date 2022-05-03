package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func sayHello(c *gin.Context) {
	// 200 状态码
	c.JSON(http.StatusOK, gin.H{
		"messange": "Hello Golang\nHello World",
	})
}

func main() {
	// 返回的是默认的路由引擎
	r := gin.Default()

	// 加载静态文件
	r.Static("/xxx", "./statics")

	// gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("templates/users/index.tmpl", "templates/posts/index.tmpl") // 解析模板
	r.LoadHTMLGlob("templates/**/*") // 加载多个文件

	// 指定用户使用GET请求访问/hello时，执行sayhello这个函数
	r.GET("/hello", sayHello)

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 模板渲染
			"title": "<a href=https://blog.csdn.net/weixin_43738067>Flynn的博客</a>",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ // 模板渲染
			"title": "<a href=https://leetcode-cn.com>Leetcode</a>",
		})
	})

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

	// 返回从网上下载的模板
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	// 启动服务
	r.Run(":9090")
}

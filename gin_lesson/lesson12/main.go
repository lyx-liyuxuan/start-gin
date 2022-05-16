package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// 一次请求对应一次相应
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		passward := c.PostForm("passward")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Passward": passward,
		})

	})
	r.Run(":9090")
}

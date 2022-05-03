package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		/*定义json
		方法一：数据使用map
		data := map[string]interface{}{
			"name":    "小王子",
			"message": "hello world",
			"age":     18,
		}*/
		data := gin.H{"name": "小王子", "message": "hello world", "age": 18}

		c.JSON(http.StatusOK, data)
	})

	//方法2：结构体, 灵活使用tag来对结构体字段做定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}
	r.GET("/ano_json", func(c *gin.Context) {
		data := msg{
			"小明",
			"Hello Golang",
			20,
		}
		c.JSON(http.StatusOK, data) //Json的序列化
	})

	r.Run(":9090")
}

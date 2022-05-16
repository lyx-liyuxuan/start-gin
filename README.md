# start-gin
[TOC]

**gin是学习Go语言编写的的Web框架study Gin Web Frame**

根据 [李文周的博客](https://www.liwenzhou.com/) 和 [B站教学视频](https://www.bilibili.com/video/BV1gJ411p7xC) 做的笔记

## 安装
`go get -u github.com/gin-gonic/gin`

`go mod tidy`

## RESTful架构
[阮一峰 RESTful架构理解](http://www.ruanyifeng.com/blog/2011/09/restful.html)
REST(Representational State Transfer),资源表现形式的状态转换
- GET方法获取资源
- POST方法创建资源
- PUT方法更新资源
- DELETE方法删除资源

## Gin渲染
### 模板渲染
```
r := gin.Default()
r.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{ //map[string]interface{}
        "title": "https://blog.csdn.net/weixin_43738067",
    })
})
r.Run()
```

### JSON渲染
```
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
```
### 获取query string参数
```
r.GET("/web", func(c *gin.Context) {
// 获取浏览器发请求携带的query string参数
name := c.Query("query") // 通过Query获取请求携带的querystring参数
age := c.Query("query")
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
```


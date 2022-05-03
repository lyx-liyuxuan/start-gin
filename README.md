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

## Gin渲染模板
```
r := gin.Default()
r.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{ //map[string]interface{}
        "title": "https://blog.csdn.net/weixin_43738067",
    })
})
r.Run()
```
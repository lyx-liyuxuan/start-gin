# start-gin
gin是学习Go语言编写的的Web框架study Gin Web Frame

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
模板引擎的使用分三步：定义模板、解析模板、渲染模板
- 定义模板  
创建 `.tmpl` 模板文件

- 解析模板  
常用以下方法解析模板，得到模板对象
```
func (t *Template) Parse(src string) (*Template, error)
func ParseFiles(filenames ...string) (*Template, error)
func ParseGlob(pattern string) (*Template, error)
```
- 渲染模板
```
func (t *Template) Execute(wr io.Writer, data interface{}) error
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```


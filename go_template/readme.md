## go模板渲染

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

## template语法

``` html
{{/* 这是一个模板注释。*/}}

// Execute()渲染模板的时候可以传入结构体也可以传入map
// 模板想访问结构体的变量，属性首字母大写才能被外部访问
// 使用map，属性名可以不用首字母大写，因为调用的是key
<p>m1</p> //m1是map结构
<p>Hello {{ .m1.name }}</p>
<p>年龄：{{ .m1.age }}</p>
<p>性别：{{ .m1.gender }}</p>

// 移除空格
// 以使用{{-语法去除模板内容左侧的所有空白符号， 使用-}}去除模板内容右侧的所有空白符号
{{- .Name -}}

// 变量
//在模板中声明变量，用来保存传入模板的数据或其他语句生成的结果
$obj := {{.}}

// pipeline
{{/* 管道符号 | 链接多个命令 */}}

{{/* 条件判断 */}}
{{ $age := .m1.age }}
{{$v1 := 0}}
{{ if $v1 }}
    {{ $age}}
{{ else }}
    没获取到
{{end}}

// range
{{/* range 循环pipeline，pipeline必须是数组、切片、字典或者通道 */}}
{{ range $index,$hobby := .hobby }}
    <p>{{$index}} - {{$hobby}}</p>
{{end}}

// with
{{ with .m1}} // 
    <p>Hello {{ .name }}</p>
    <p>年龄：{{ .age }}</p>
    <p>性别：{{ .gender }}</p>
{{end}}

// index
// 索引数组、切片或者字典
{{index .hobby 2}}
```

## 嵌套模板

```html

```

## 修改模板的标识符

```html

```

## text/template与html/tempalte

html/tempalte渲染JS代码，针对的是需要返回HTML内容的场景，在模板渲染过程中会对一些有风险的内容进行转义，以此来防范跨站脚本攻击。

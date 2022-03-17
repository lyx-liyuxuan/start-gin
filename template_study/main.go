package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 属性首字母大写才能被外部访问
type User struct {
	Name   string
	Gender string
	Age    int
}

// 1. f.tmpl 就是第一步定义模板
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./f.tmpl") // 不能直接右键Run go build
	if err != nil {
		fmt.Printf("Parse template failed, err:%v", err)
		return
	}

	// 3.渲染模板
	u1 := User{
		Name:   "小明",
		Gender: "男",
		Age:    18,
	}

	// 使用map，属性名可以不用首字母大写，因为调用的是key
	m1 := map[string]interface{}{
		"name":   "小红",
		"gender": "女",
		"age":    16,
	}

	hobbyList := []string{
		"篮球",
		"羽毛球",
		"乒乓球",
	}

	// 传入结构体
	//err = t.Execute(w, u1)
	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}

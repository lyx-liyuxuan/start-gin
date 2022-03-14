package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 1. hello.tmpl 就是第一步定义模板
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./hello.tmpl") // 不能直接右键Run go build
	if err != nil {
		fmt.Printf("Parse template failed, err:%v", err)
		return
	}
	// 3.渲染模板
	name := "小巨人"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 解析模板，先写根模板，在写继承的模板
	t, err := template.ParseFiles("base.tmpl", "index.tmpl")
	if err != nil {
		fmt.Printf("parse template faliles, err:%v\n", err)
		return
	}
	// 渲染模板
	name := "小明"
	t.ExecuteTemplate(w, "index.tmpl", name)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 解析模板，先写根模板，在写继承的模板
	t, err := template.ParseFiles("base.tmpl", "home.tmpl")
	if err != nil {
		fmt.Printf("parse template faliles, err:%v\n", err)
		return
	}
	// 渲染模板
	name := "小明"
	t.ExecuteTemplate(w, "home.tmpl", name)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server failed, err:%v\n", err)
		return
	}
}

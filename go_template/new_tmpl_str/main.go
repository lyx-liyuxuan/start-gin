package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	name := "小明"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("execute template failed,err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v\n", err)
		return
	}
}

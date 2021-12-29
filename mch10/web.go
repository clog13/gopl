package mch10

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Web() {
	http.HandleFunc("/clob", sayhelloName)   // 设置访问的路由
	err := http.ListenAndServe(":8089", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	defer func() { fmt.Println("--------END--------") }()
	err := r.ParseForm() // 解析参数，默认是不会解析的
	if err != nil {
		return
	}
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("-------------------")
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_parameter"])
	for k, v := range r.Form {
		fmt.Printf("key:%s, val:%s\n", k, strings.Join(v, "-"))
		//fmt.Println("val:", strings.Join(v, ""))
	}
	_, err = fmt.Fprintf(w, "Hello astaxie!") // 这个写入到w的是输出到客户端的
	if err != nil {
		return
	}
}

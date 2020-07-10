package main

import (
	"net/http"
	"ts_golang/LearnGoWithTests/DI"
)

func main() {
	// 编写HTTP处理程序时，将获得一个http.ResponseWriter和用于发出请求的http.Request。
	// 当您实现服务器时，可以使用writer来编写响应。
	http.ListenAndServe(":5000", http.HandlerFunc(HelloWorld.MyGreeterHandler))
}

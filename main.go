package main

import (
	"os"
	"time"
	HelloWorld "ts_golang/LearnGoWithTests/Mocking"
)

func main() {
	// DI: 编写HTTP处理程序时，将获得一个http.ResponseWriter和用于发出请求的http.Request。
	// DI: 当您实现服务器时，可以使用writer来编写响应。
	//http.ListenAndServe(":5000", http.HandlerFunc(HelloWorld.MyGreeterHandler))
	//HelloWorld.CountDown(os.Stdout)
	//sleeper := &HelloWorld.DefaultSleeper{}
	//HelloWorld.CountDown(os.Stdout, sleeper)
	sleeper := &HelloWorld.ConfigurableSleeper{1 * time.Second, time.Sleep}
	HelloWorld.CountDown(os.Stdout, sleeper)
}

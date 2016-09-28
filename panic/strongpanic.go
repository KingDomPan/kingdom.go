package main

import (
	"log"
	"net/http"
)

// 如何让一个web程序健壮的运行
// 服务器异常的时候记录异常, 程序继续运行
func main() {
	http.HandleFunc("/", WrapperRecover(HomePage))
}

// 针对每个响应函数都设置一个延迟recover处理函数, (重复, 添加包装方法)
func HomePage(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			log.Printf("[%v] caught panic %v", r.RemoteAddr, x)
		}
	}()
}

func WrapperRecover(function func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] cught panic: %v", request.RemoteAddr, x)
			}
		}()
		function(writer, request)
	}
}

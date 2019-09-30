package main

import (
	"go-web-test/pkg/setting"
	"go-web-test/routers"
	"net/http"
)


func main() {

	r := routers.InitRouter()

	s := & http.Server{
		Addr:              setting.HttpPort,  //监听的TCP地址
		Handler:           r,  //http句柄，实质为ServeHTTP，用于处理程序响应HTTP请求
		TLSConfig:         nil,  //安全传输层协议（TLS）的配置
		ReadTimeout:       0,  //允许读取的最大时间
		ReadHeaderTimeout: 0,  //允许读取请求头的最大时间
		WriteTimeout:      0,  //允许写入的最大时间
		IdleTimeout:       0,  //等待的最大时间
		MaxHeaderBytes:    1 << 20,  //请求头的最大字节数
	}

	s.ListenAndServe()

	//port := setting.HttpPort
	//_ = r.Run(":"+ port)
}
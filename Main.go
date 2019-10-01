package main

import (
	"go-web-test/pkg/setting"
	"go-web-test/routers"
	"net/http"
	"time"
)


func main() {

	r := routers.InitRouter()

	s := & http.Server{
		Addr:              setting.HttpPort,  //监听的TCP地址
		Handler:           r,  //http句柄，实质为ServeHTTP，用于处理程序响应HTTP请求
		ReadTimeout:       setting.ReadTimeout * time.Second,  //允许读取的最大时间
		WriteTimeout:      setting.WriteTimeout * time.Second,  //允许写入的最大时间
		MaxHeaderBytes:    1 << 20,  //请求头的最大字节数
	}

	s.ListenAndServe()
}
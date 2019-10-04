package main

import (
	"github.com/unknwon/com"
	"go-web-test/models"
	"go-web-test/pkg/gredis"
	"go-web-test/pkg/setting"
	"go-web-test/routers"
	"net/http"
)

/**
初始化项目
 */
func init()  {
	setting.Setup()
	models.SetUp()
	gredis.Setup()
}

func main() {

	r := routers.InitRouter()

	s := & http.Server{
		Addr:              ":" + com.ToStr(setting.ServerSetting.HttpPort),  //监听的TCP地址
		Handler:           r,  //http句柄，实质为ServeHTTP，用于处理程序响应HTTP请求
		ReadTimeout:       setting.ServerSetting.ReadTimeout,  //允许读取的最大时间
		WriteTimeout:      setting.ServerSetting.WriteTimeout,  //允许写入的最大时间
		MaxHeaderBytes:    1 << 20,  //请求头的最大字节数
	}

	s.ListenAndServe()
}
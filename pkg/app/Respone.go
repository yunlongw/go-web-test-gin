package app

import (
	"github.com/gin-gonic/gin"
	"go-web-test-gin/pkg/e"
)

type Gin struct {
	C *gin.Context
}
/**
统一定义 api 的接口返回类型
 */
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}
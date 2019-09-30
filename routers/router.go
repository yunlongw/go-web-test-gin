package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine  {
	r := gin.New()

	r.LoadHTMLGlob("templates/*")

	Log(r)

	r.Use(gin.Recovery())

	return r
}
package routers

import (
	"github.com/gin-gonic/gin"
	"go-web-test/controller"
	"net/http"
)

func Web(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code" : 1,
		})
	})

	r.GET("/index", controller.IndexController{}.Index())

}

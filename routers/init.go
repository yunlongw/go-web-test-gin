package routers

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")
	//
	Log(r)
	//
	Api(r)
	//
	Web(r)
}

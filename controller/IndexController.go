package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {

}

func (i IndexController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
		//c.JSON(http.StatusOK, gin.H{
		//	"code" : 1,
		//})
	}
}
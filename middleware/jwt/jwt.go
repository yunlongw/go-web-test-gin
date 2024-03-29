package jwt

import (
	"github.com/gin-gonic/gin"
	"go-web-test-gin/pkg/e"
	"go-web-test-gin/pkg/util"
	"net/http"
	"time"
)

var Claims *util.Claims
var err error

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.GetHeader("token")

		if token == "" {
			code = e.INVALID_PARAMS
		}else {
			Claims, err = util.ParseToken(token)  // 验证token
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}else if time.Now().Unix() > Claims.ExpiresAt{   // 超时检测
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS{
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()  //继续执行
	}

}

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-test/pkg/setting"
	v1 "go-web-test/routers/api/v1"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func InitRouter() *gin.Engine  {
	r := gin.New()

	//r.LoadHTMLGlob("templates/*")

	gin.DisableConsoleColor()
	//f, _ := os.Create("logs/gin.log")
	f, err := os.OpenFile("logs/gin.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)

	//r.Use(gin.LoggerWithWriter(handle, ""))
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code" : 1,
		})
	})


	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)

		apiv1.POST("/tags", v1.AddTag)
	}




	return r
}
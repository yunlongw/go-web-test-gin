package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-test-gin/middleware/jwt"
	"go-web-test-gin/pkg/setting"
	"go-web-test-gin/routers/api"
	v1 "go-web-test-gin/routers/api/v1"
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

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code" : 1,
		})
	})


	r.POST("/auth", api.GetAuth)// 登录
	r.POST("/register", v1.Register) //注册


	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)


		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}






	return r
}
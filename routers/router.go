package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-web-test-gin/middleware/jwt"
	"go-web-test-gin/pkg/logging"
	"go-web-test-gin/pkg/setting"
	"go-web-test-gin/routers/api"
	v1 "go-web-test-gin/routers/api/v1"
	"io"
	"net/http"
	"time"
)

func RunServer() error {
	r := InitRouter()

	s := &http.Server{
		Addr:           ":" + com.ToStr(setting.ServerSetting.HttpPort), //监听的TCP地址
		Handler:        r,                                               //http句柄，实质为ServeHTTP，用于处理程序响应HTTP请求
		ReadTimeout:    setting.ServerSetting.ReadTimeout,               //允许读取的最大时间
		WriteTimeout:   setting.ServerSetting.WriteTimeout,              //允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                                         //请求头的最大字节数
	}

	err := s.ListenAndServe()
	if err != nil {
		logging.Error(err)
		return err
	}
	return err
}

func InitRouter() *gin.Engine  {
	r := gin.New()

	//r.LoadHTMLGlob("templates/*")

	gin.DisableConsoleColor()

	f := logging.OpenLogFile("runtime/gin.log")
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
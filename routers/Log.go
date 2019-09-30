package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

func Log(r *gin.Engine)  {

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

}

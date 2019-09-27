package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" bdinding:"required"`
	Age      int    `form:"age" json:"age"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hellow world")
	})

	router.Use(MiddleWare())
	{
		router.GET("/middleware", func(c *gin.Context) {
			request := c.MustGet("request").(string)
			req, _ := c.Get("request")
			c.JSON(http.StatusOK, gin.H{
				"middile_request": request,
				"request":         req,
			})
		})
	}

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	/**
	图片上传
	*/
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)

		file, header, err := c.Request.FormFile("upload")

		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
		}

		filename := header.Filename
		fmt.Println(file, err, filename)

		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		defer out.Close()

		//for k,v :=  range file.Read()  {
		//	fmt.Println(k, v)
		//}

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusCreated, "upload successfull")
	})

	/**
	多图上传
	*/
	router.POST("/multi/upload", func(c *gin.Context) {
		err := c.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
		}

		formdata := c.Request.MultipartForm
		files := formdata.File["upload"]

		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}

			out, err := os.Create(files[i].Filename)
			defer out.Close()

			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(out, file)

			if err != nil {
				log.Fatal(err)
			}

			c.String(http.StatusCreated, "upload success \n")

		}
	})

	router.LoadHTMLGlob("templates/*")
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	router.POST("/login", func(c *gin.Context) {
		var user User
		var err error

		contentType := c.Request.Header.Get("Content-type")

		//switch contentType {
		//case "application/json":
		//	err = c.BindJSON(&user)
		//case "application/x-www-form-urlencoded":
		//	err = c.BindWith(&user, binding.Form)
		//}

		err = c.Bind(&user)

		if err != nil {
			fmt.Println("contentType:" + contentType)
			fmt.Println(user)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"user":   user.Username,
			"passwd": user.Passwd,
			"age":    user.Age,
		})

	})

	/**
	重定向
	*/
	router.GET("/redict/google", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://google.com")
	})

	v1 := router.Group("/v1")
	v1.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "\nv1 login\n")
	})

	v2 := router.Group("/v2")
	v2.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "\nv2 login\n")
	})

	router.Run(":8000")
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "clinet_request")
		c.Next()
		fmt.Println("before middleware")
	}
}

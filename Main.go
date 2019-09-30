package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-web-test/routers"
	"log"
	"os"
)


func main() {
	r := gin.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routers.Init(r)

	port := os.Getenv("CONFIG_PORT")
	_ = r.Run(":"+ port)
}
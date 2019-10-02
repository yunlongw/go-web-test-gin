package setting

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Debug        bool
	RunMode      string

	PageSize int
	JwtSecret string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase()  {
	Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))

}

func LoadServer()  {
	RunMode = os.Getenv("RUN_MODE")
	HttpPort = ":" + os.Getenv("HTTP_PORT")
	ReadTimeout = 60
	WriteTimeout = 60
}

func LoadApp()  {
	JwtSecret = os.Getenv("JWT_SECRET")
	PageSize = 15
}
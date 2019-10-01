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
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PageSize = 15
	ReadTimeout = 60
	WriteTimeout = 60

	RunMode = os.Getenv("RUN_MODE")
	HttpPort = ":" + os.Getenv("HTTP_PORT")
	Debug, err = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		log.Fatal(err)
	}

}

func ParseBool(str string) (bool) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true
	case "0", "f", "F", "false", "FALSE", "False":
		return false
	}
	return false
}

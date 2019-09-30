package setting

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	Debug bool
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HttpPort = os.Getenv("HttpPort")
	r, _ := strconv.ParseInt(os.Getenv("READ_TIMEOUT"), 10, 64)
	ReadTimeout = time.Duration(r)  * time.Second

	w, _ := strconv.ParseInt(os.Getenv("WRITE_TIMEOUT"), 10, 64)
	WriteTimeout = time.Duration(w)  * time.Second

	Debug = os.Getenv("CONFIG_DEBUG")
}


func ParseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False":
		return false, nil
	}
	return false, nil
}

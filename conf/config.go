package conf

import (
	"fmt"
	"os"
)

const (
	MachineryServer = "conf/config.yml"
)

func GetConfigYml() string {
	dir, _ := os.Getwd()
	return fmt.Sprintf("%s/%s", dir, MachineryServer)
}

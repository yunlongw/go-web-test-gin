package main

import (
	"github.com/urfave/cli"
	"go-web-test-gin/models"
	"go-web-test-gin/pkg/gredis"
	"go-web-test-gin/pkg/logging"
	"go-web-test-gin/pkg/queue"
	"go-web-test-gin/pkg/setting"
	"go-web-test-gin/routers"
	"log"
	"os"
)

func init() {
	setting.Setup()
	models.SetUp()
	err := gredis.Setup()
	if err != nil {
		logging.Error(err)
	}
	err = queue.Setup()
	if err != nil {
		logging.Error(err)
	}
}

func main() {
	app := *cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				if err := queue.RunWorker(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "server",
			Usage: "server is running",
			Action: func() error {
				if err := routers.RunServer(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	    logging.Error(err)
	}
}





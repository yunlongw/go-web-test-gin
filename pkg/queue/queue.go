package queue

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"go-web-test-gin/conf"
	"go-web-test-gin/pkg/logging"
	"go-web-test-gin/service/task_service"
)

var (
	MServer *machinery.Server
	cnf     *config.Config
	err     error
	tasks   map[string]interface{}
)

func Setup() error {

	tasks = map[string]interface{}{
		"login":               task_service.LoginTask,
		"long_running_task": task_service.LongRunningTask,
	}

	configPath := conf.GetConfigYml()
	if cnf, err = loadConfig(configPath); err != nil {
		logging.Error(err)
		return err
	}
	if MServer, err = machinery.NewServer(cnf); err != nil {
		logging.Error(err)
		return err
	}

	return nil
}

func loadConfig(configPath string) (*config.Config, error) {
	if configPath != "" {
		return config.NewFromYaml(configPath, true)
	}
	return config.NewFromEnvironment(true)
}

func RunWorker() error {
	err = MServer.RegisterTasks(tasks)
	if err != nil {
		logging.Error(err)
		return err
	}
	workers := MServer.NewWorker("worker_test", 10)
	err = workers.Launch()
	if err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

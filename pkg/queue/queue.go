package queue

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"go-web-test-gin/conf"
	"go-web-test-gin/pkg/logging"
	"go-web-test-gin/service/task_service"
)

var (
	cnf     *config.Config
	tasks   map[string]interface{}
)

const consumerTag = "machinery_worker"
const concurrency = 10

func init() {
	tasks = map[string]interface{}{
		"login":             task_service.LoginTask,
		"long_running_task": task_service.LongRunningTask,
	}
}

/**
返回 queue 的 server
 */
func StartServer() (servers *machinery.Server, err error) {

	configPath := conf.GetConfigYml()

	if cnf, err = loadConfig(configPath); err != nil {
		logging.Error(err)
		return nil, err
	}

	if servers, err = machinery.NewServer(cnf); err != nil {
		logging.Error(err)
		return nil, err
	}

	return servers, servers.RegisterTasks(tasks)
}

func loadConfig(configPath string) (*config.Config, error) {
	if configPath != "" {
		return config.NewFromYaml(configPath, true)
	}
	return config.NewFromEnvironment(true)
}


/**
创建 worker 负责处理队列中的内容
 */
func RunWorker() error {
	servers, err := StartServer()
	if err != nil {
		logging.Error(err)
		return err
	}

	workers := servers.NewWorker(consumerTag, concurrency)
	err = workers.Launch()
	if err != nil {
		logging.Error(err)
		return err
	}

	return nil
}

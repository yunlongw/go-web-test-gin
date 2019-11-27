package task_service

import (
	"github.com/RichardKnop/machinery/v1/log"
	"time"
)

func LongRunningTask() error {
	log.INFO.Print("Long running task started")
	for i := 0; i < 100; i++ {
		log.INFO.Print(10 - i)
		time.Sleep(1 * time.Second)
		//这里面可以通过redis更新任务进度
	}
	//这里可以用redis取存储任务结果，sender可以异步轮询或者指定多久后取redis取结果
	log.INFO.Print("Long running task finished")
	return nil
}

func LoginTask(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}
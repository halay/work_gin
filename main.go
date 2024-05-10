package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"work_gin/mqueue/asynq/job"
	"work_gin/mqueue/asynq/scheduler"
	"work_gin/routes"
	"work_gin/utils/log"
)

var (
	g errgroup.Group
)

func main() {
	log.InitLog()
	//model.InitDb()
	//model.InitRedis()
	g.Go(func() error {
		return job.InitJob()
	})
	
	g.Go(func() error {
		return scheduler.InitScheduler()
	})

	g.Go(func() error {
		return routes.InitRouter()
	})

	if err := g.Wait(); err != nil {
		log.Log.Info("start service err," + err.Error())
		fmt.Println("err", err)
	}

}

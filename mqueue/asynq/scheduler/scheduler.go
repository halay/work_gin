package scheduler

import (
	"context"
	"github.com/hibiken/asynq"
	"time"
	"work_gin/mqueue/asynq/scheduler/logic"
	"work_gin/utils"
	"work_gin/utils/log"
)

var (
	Scheduler *asynq.Scheduler
)

func InitScheduler() error {
	location, _ := time.LoadLocation("Asia/Shanghai")
	Scheduler = asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     utils.RedisArr,
			Password: utils.RedisPassword,
			DB:       1,
		},
		&asynq.SchedulerOpts{
			// Schedule
			Location: location,
			Logger:   log.Log,
			LogLevel: 1,
			PostEnqueueFunc: func(taskInfo *asynq.TaskInfo, err error) {
				log.Log.Infof("postEnqueueFunc task info,%+v \n", taskInfo)
				if err != nil {
					log.Log.Errorf("postEnqueueFunc error,%+v \n", err)
				}
			},
		},
	)
	cronScheduler := logic.NewCronScheduler(context.Background(), Scheduler)
	cronScheduler.Register()
	log.Log.Infof("scheduler init success \n")
	if err := Scheduler.Run(); err != nil {
		log.Log.Errorf("scheduler run error,%+v \n", err)
		return err
	}
	return nil
}

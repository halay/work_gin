package job

import (
	"context"
	"github.com/hibiken/asynq"
	"time"
	"work_gin/mqueue/asynq/job/logic"
	"work_gin/utils"
	"work_gin/utils/log"
)

func InitJob() error {
	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     utils.RedisArr,
			Password: utils.RedisPassword,
			DB:       1,
		},
		asynq.Config{
			Concurrency:         10,
			Logger:              log.Log,
			LogLevel:            1,
			HealthCheckInterval: 10 * time.Second,
			HealthCheckFunc: func(err error) {
				if err != nil {
					log.Log.Infof("asynq server exec task HealthCheckFunc ======== >>>>>>>>>>>  err : %+v \n", err)
				}
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				if err != nil {
					log.Log.Infof("asynq server exec task info :%v, ErrorHandler ======== >>>>>>>>>>>  err : %+v \n", task, err)
				}
			}),
			IsFailure: func(err error) bool {
				if err != nil {
					log.Log.Infof("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				}
				return true
			},
		},
	)
	Cronjob := logic.NewCronjob(context.Background())
	mux := Cronjob.Register()
	log.Log.Info("asynq启动成功")
	if err := server.Run(mux); err != nil {
		log.Log.Infof("启动asynq失败，请检查参数 err : %+v \n", err)
		return err
	}
	return nil
}

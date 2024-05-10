package logic

import (
	"github.com/hibiken/asynq"
	"work_gin/mqueue/asynq/common"
	"work_gin/utils/log"
)

func (c *CronScheduler) pushDataScheduler() {
	task := asynq.NewTask(common.JobPushData, nil)
	//entryID, err := c.Scheduler.Register("* * * * *", task)
	entryID, err := c.Scheduler.Register("@every 30s", task)
	if err != nil {
		log.Log.Errorf("pushDataScheduler error,task %v,err:%v ", task, err)
	}
	log.Log.Infof("pushDataScheduler register success,task %v,entryID:%v ", task, entryID)
}

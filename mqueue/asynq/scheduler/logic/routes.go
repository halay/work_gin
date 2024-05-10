package logic

import (
	"context"
	"github.com/hibiken/asynq"
)

type CronScheduler struct {
	ctx       context.Context
	Scheduler *asynq.Scheduler
}

func NewCronScheduler(ctx context.Context, scheduler *asynq.Scheduler) *CronScheduler {
	return &CronScheduler{
		ctx:       ctx,
		Scheduler: scheduler,
	}
}
func (c *CronScheduler) Register() {
	c.pushDataScheduler()
}

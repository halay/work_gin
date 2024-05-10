package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"work_gin/mqueue/asynq/common"
)

type Cronjob struct {
	ctx context.Context
}

func NewCronjob(ctx context.Context) *Cronjob {
	return &Cronjob{ctx: ctx}
}
func (c *Cronjob) Register() *asynq.ServeMux {
	// 1. 定义一个任务队列
	mux := asynq.NewServeMux()
	mux.Handle(common.JobSendEmail, NewSendEmail())
	mux.Handle(common.JobPushData, NewPushData())
	return mux
}

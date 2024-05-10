package logic

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"work_gin/utils/log"
)

type PushData struct {
}

func NewPushData() *PushData {
	return &PushData{}
}
func (s *PushData) ProcessTask(ctx context.Context, t *asynq.Task) error {
	fmt.Println("this push data processTask success,every minute run once!")
	log.Log.Infof("this push data processTask success,every minute run once!\n")
	return nil
}
